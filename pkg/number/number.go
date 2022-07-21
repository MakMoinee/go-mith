package number

import (
	"fmt"
)

func IsNumberMatchLength(n int, length int) bool {
	str := fmt.Sprintf("%v", n)
	if len(str) == length {
		return true
	}
	return false
}

// IsConsecutiveNumber() - checks if the []int have consecutive value numbers in its elements
func IsConsecutiveNumber(n []int) bool {
	IsConsecutive := true
	for i := 1; i < len(n); i++ {
		if n[i] != n[i-1]+1 {
			IsConsecutive = false
		}
	}
	return IsConsecutive
}
