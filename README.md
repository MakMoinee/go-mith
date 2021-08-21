# go-mith
## Features
- Consolidates useful formulas for starters of golang programming

## Packages
- Palindrome - checks if the value given is palindrome or not
- Power Formula - One of the science formula. It's used to calculate the power from a given work and time values
- Stair Case (Hacker Rank Solution) - prints a staircase of size n.

## Installation
- `go get github.com/MakMoinee/go-mith`

## Sample Code
```
import (
	"fmt"

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
}
```

## Stair Case
```
import (
	"fmt"

	"github.com/MakMoinee/go-mith/pkg/manipulate"
)

func main() {
	fmt.Println("Starting main.go")

	num2 := 2
	manipulate.GetStairCase(int32(num2))
}
```
- Result:
```
      #
     ##
    ###
   ####
  #####
 ######
#######
```