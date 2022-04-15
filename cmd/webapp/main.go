package main

import (
	"fmt"
	"sync"

	"github.com/MakMoinee/go-mith/pkg/concurrency"
	"github.com/MakMoinee/go-mith/pkg/encrypt"
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
	fmt.Println()
	fmt.Println()

	// hashPass, _ := encrypt.HashPassword("admin123")
	// fmt.Println("Hash Pass", hashPass)
	fmt.Println("Hash Check:", encrypt.CheckPasswordHash("admin123", "$2a$14$aqNaRmfnkcoM6wD5SfUAlOOJUKGffU2QTKimFWgfLNBG7b0fiXHdq"))

	// default concurrent sample
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		concurrentService := concurrency.NewService()
		serviceStruct := concurrentService.GetServiceStruct()

		data1 := concurrency.ProcessItemSliceString(1, []string{"1", "2"}, &serviceStruct)
		data2 := concurrency.ProcessItemSliceString(2, []string{"1", "2"}, &serviceStruct)
		// data3 := concurrency.ProcessItemSliceString(3, []string{"1", "2"}, &serviceStruct) //error
		// fmt.Println(data3)
		fmt.Println(data1)
		fmt.Println(data2)

		data4, err := concurrency.ProcessItem(1, []string{"1", "2"}, concurrentService)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println("[]string >>", data4)

		data5, err := concurrency.ProcessItem(2, []int{1, 2, 3, 4, 5}, concurrentService)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println("[]int >>> ", data5)

	}()
	wg.Wait()

}

func processSomething(data interface{}) interface{} {
	return data
}
