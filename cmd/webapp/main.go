package main

import (
	"fmt"

	"github.com/MakMoinee/go-mith/pkg/manipulate"
	"github.com/MakMoinee/go-mith/pkg/palindrome"
)

func main() {
	fmt.Println("Starting main.go")

	// Testing palindrome

	// Pass Palindrome Number
	num1 := 121
	fmt.Println(palindrome.IsNumberPalindrome(num1)) // It must print true

	str1 := "aabbaa"
	fmt.Println(palindrome.IsStringPalindrome(str1)) // it must print true

	num2 := 7
	manipulate.GetStairCase(int32(num2))
}
