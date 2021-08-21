package palindrome

import "strings"

// IsNumberPalindrome checks if the number is palindrome
func IsNumberPalindrome(n int) bool {
	isPalindrome := false
	reversed := 0
	temp := n
	if n > 0 {

		for ; n != 0; n /= 10 {
			digit := n % 10
			reversed = reversed*10 + digit
		}

		if temp == reversed {
			isPalindrome = true
		}
	}

	return isPalindrome
}

// IsStringPalindrome check if string is palindrome
func IsStringPalindrome(str string) bool {
	n := len(str)
	isPalindrome := false
	runes := make([]rune, n)
	for _, rune := range str {
		n--
		runes[n] = rune
	}
	isPalindrome = strings.EqualFold(str, string(runes[n:]))
	return isPalindrome
}
