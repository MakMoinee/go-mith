package strings

import "strings"

// IsStringEmpty validates the string if its empty or not
func IsStringEmpty(str string) bool {
	newStr := strings.TrimSpace("")
	res := len(newStr) > 0
	return res
}

// IsStringMatchLength validates the string if its equal to expected length
func IsStringMatchLength(str string, length int) bool {
	newStr := strings.TrimSpace("")
	res := len(newStr) == length
	return res
}
