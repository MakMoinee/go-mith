package utils

import (
	"errors"
	"strings"
)

type Number interface {
	int | float32 | float64 | int32 | int16 | int8
}

// SumNumber - adds all the numbers passed
func SumNumber[T Number](t []T) T {
	var total T

	for _, data := range t {
		total += data
	}

	return total
}

func GetTokenString(header string) (string, error) {
	parts := strings.Fields(header)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("Missing or invalid authorization header")
	}
	return parts[1], nil
}
