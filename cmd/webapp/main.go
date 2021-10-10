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

	str2 := "abab"
	fmt.Println(palindrome.IsStringPalindrome(str2)) // it must print false

	num2 := 7
	manipulate.GetStairCase(int32(num2))
	d1, err1 := manipulate.CompareData(1.10, 1.10)
	d2, err2 := manipulate.CompareData(2, 2)
	fmt.Println("CompareData (1,1.0) == " + fmt.Sprintf("%v,%v", d1, err1))
	fmt.Println("CompareData (2,2) == " + fmt.Sprintf("%v,%v", d2, err2))
}
