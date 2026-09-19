package idemstore

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	statusProcessing = "processing"
	statusCompleted  = "completed"
)

type RedisStore struct {
	client *redis.Client
	config Config
}

type redisValue struct {
	Status      string `json:"status"`
	RequestHash string `json:"request_hash,omitempty"`
	Result      []byte `json:"result,omitempty"`
}

func NewRedis(client *redis.Client, config Config) *RedisStore {
	if config.TTL <= 0 {
		config.TTL = 24 * time.Hour
	}

	if config.ProcessingTTL <= 0 {
		config.ProcessingTTL = 5 * time.Minute
	}

	if config.PollInterval <= 0 {
		config.PollInterval = 100 * time.Millisecond
	}

	if config.WaitTimeout <= 0 {
		config.WaitTimeout = 10 * time.Second
	}

	return &RedisStore{
		client: client,
		config: config,
	}
}

func (s *RedisStore) Execute(
	ctx context.Context,
	key string,
	requestHash string,
	fn func(context.Context) ([]byte, error),
) ([]byte, error) {

	if key == "" {
		return nil, fmt.Errorf("idempotency key cannot be empty")
	}

	if requestHash == "" {
		return nil, fmt.Errorf("request hash cannot be empty")
	}

	redisKey := "idempotency:" + key

	// Try to acquire the idempotency key.
	acquired, err := s.acquire(ctx, redisKey, requestHash)
	if err != nil {
		return nil, err
	}

	if acquired {
		// We own the request.
		result, err := fn(ctx)

		if err != nil {
			// Remove the key so a retry can execute again.
			_ = s.client.Del(ctx, redisKey).Err()

			return nil, err
		}

		if err := s.complete(ctx, redisKey, requestHash, result); err != nil {
			return nil, err
		}

		return result, nil
	}

	// Another request already owns or completed this key.
	return s.waitForResult(ctx, redisKey, requestHash)
}

func (s *RedisStore) acquire(
	ctx context.Context,
	key string,
	requestHash string,
) (bool, error) {

	value := redisValue{
		Status:      statusProcessing,
		RequestHash: requestHash,
	}

	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	ok, err := s.client.SetNX(
		ctx,
		key,
		data,
		s.config.ProcessingTTL,
	).Result()

	if err != nil {
		return false, fmt.Errorf("acquire idempotency key: %w", err)
	}

	return ok, nil
}

func (s *RedisStore) complete(
	ctx context.Context,
	key string,
	requestHash string,
	result []byte,
) error {

	value := redisValue{
		Status:      statusCompleted,
		RequestHash: requestHash,
		Result:      result,
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal idempotency result: %w", err)
	}

	if err := s.client.Set(
		ctx,
		key,
		data,
		s.config.TTL,
	).Err(); err != nil {
		return fmt.Errorf("store idempotency result: %w", err)
	}

	return nil
}

func (s *RedisStore) waitForResult(
	ctx context.Context,
	key string,
	requestHash string,
) ([]byte, error) {

	waitCtx, cancel := context.WithTimeout(
		ctx,
		s.config.WaitTimeout,
	)
	defer cancel()

	ticker := time.NewTicker(s.config.PollInterval)
	defer ticker.Stop()

	for {
		value, err := s.get(waitCtx, key)
		if err != nil {
			return nil, err
		}

		if value == nil {
			// The processing key expired or was removed.
			// Try executing the request again.
			acquired, err := s.acquire(
				waitCtx,
				key,
				requestHash,
			)

			if err != nil {
				return nil, err
			}

			if acquired {
				return nil, ErrInProgress
			}

			continue
		}

		if value.RequestHash != requestHash {
			return nil, ErrConflict
		}

		if value.Status == statusCompleted {
			return value.Result, nil
		}

		select {
		case <-waitCtx.Done():
			return nil, ErrInProgress

		case <-ticker.C:
		}
	}
}

func (s *RedisStore) get(
	ctx context.Context,
	key string,
) (*redisValue, error) {

	data, err := s.client.Get(ctx, key).Bytes()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get idempotency key: %w", err)
	}

	var value redisValue

	if err := json.Unmarshal(data, &value); err != nil {
		return nil, fmt.Errorf("unmarshal idempotency value: %w", err)
	}

	return &value, nil
}

func HashRequest(body []byte) string {
	hash := sha256.Sum256(body)

	return hex.EncodeToString(hash[:])
}
