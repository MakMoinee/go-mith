# Go-Mith [![Go Report Card](https://goreportcard.com/badge/github.com/MakMoinee/go-mith)](https://goreportcard.com/report/github.com/MakMoinee/go-mith)

## Features
- Consolidates useful formulas for starters of golang programming

## Packages
- Palindrome - checks if the value given is palindrome or not
- Power Formula - One of the science formula. It's used to calculate the power from a given work and time values
- Stair Case (Hacker Rank Solution) - prints a staircase of size n.
- Concurrency Package - useable interface for any concurrent calls.
- Goserve Package - build http service to start your API with the support of injecting certs and reading config from settings.yaml

## Installation
- `go get github.com/MakMoinee/go-mith`

## Sample Code
```go
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
```go
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

## Concurrent Package

```go
package main

import (
	"fmt"
	"sync"

	"github.com/MakMoinee/go-mith/pkg/concurrency"
)

func main() {
	// default concurrent sample
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		// initialize concurrent service
		concurrentService := concurrency.NewService()

		// ProcessItem - dynamically process item passed on the function.
		// Current supported data types are: []string, []int
		// TODO: int, string
		data, err := concurrency.ProcessItem(1, []string{"1", "2"}, concurrentService)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println("[]string >>", data)
	}()
	wg.Wait()
}

```

## goserve package

```go
package main

import (
	"log"

	"github.com/MakMoinee/go-mith/pkg/goserve"
)

func main() {
	httpService := goserve.NewService(SERVER_PORT)
	httpService.EnableProfiling(SERVER_ENABLE_PROFILING)
	log.Println("Server Starting in Port ", SERVER_PORT)
	if err := httpService.Start(); err != nil {
		panic(err)
	}
}
```
