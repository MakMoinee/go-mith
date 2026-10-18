package gomithauth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MithAuthClaims struct {
	jwt.RegisteredClaims
}

type mithAuthService struct {
	contextKey   string
	username     string
	ServerIssuer string
	ServerSecret string
	TokenMapping map[string]*MithAuthClaims
}

type MithAuthService interface {
	GetToken(string) (string, error)
	ValidateToken(string) bool
	AuthMiddleware(next http.Handler) http.Handler
}

func NewMithAuthService(contextKey, serverIssuer, serverSecret string) MithAuthService {
	return &mithAuthService{
		contextKey:   contextKey,
		ServerIssuer: serverIssuer,
		ServerSecret: serverSecret,
		TokenMapping: make(map[string]*MithAuthClaims),
	}
}

func (svc *mithAuthService) GetToken(username string) (string, error) {
	svc.username = username
	claims := svc.getClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(svc.ServerSecret))
	if err == nil {
		svc.TokenMapping[tokenString] = &claims
	}
	return tokenString, err
}

func (svc *mithAuthService) getClaims() MithAuthClaims {
	now := time.Now()
	return MithAuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   svc.username,
			Issuer:    svc.ServerIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
		},
	}
}

func (svc *mithAuthService) IsClaimsSubjectEmpty(tokenString string) bool {
	result := false //default to false

	claims, exist := svc.TokenMapping[tokenString]
	if !exist {
		return result
	}
	if len(claims.Subject) == 0 {
		result = true
	}
	return result
}

func (svc *mithAuthService) getServerSecret(token *jwt.Token) (any, error) {
	return []byte(svc.ServerSecret), nil
}

func (svc *mithAuthService) parseWithClaims(tokenString string) (*jwt.Token, error) {
	claims, exist := svc.TokenMapping[tokenString]
	if !exist {
		return nil, errors.New("Token Not Exist")
	}

	return jwt.ParseWithClaims(
		tokenString,
		claims,
		svc.getServerSecret,
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
		jwt.WithIssuer(svc.ServerIssuer),
		jwt.WithExpirationRequired(),
	)
}

func (svc *mithAuthService) ValidateToken(tokenString string) bool {
	isValid := false

	token, err := svc.parseWithClaims(tokenString)
	if err != nil {
		return isValid
	}

	isSubjectClaimValid := svc.IsClaimsSubjectEmpty(tokenString)

	if token.Valid && isSubjectClaimValid {
		isValid = true
	}

	return isValid
}

func (svc *mithAuthService) getClaimsWithToken(tokenString string) *MithAuthClaims {
	return svc.TokenMapping[tokenString]
}
