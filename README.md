<div align="center">

# Go-Mith

Consolidates useful formulas for starters of golang programming

[![Documentation](https://img.shields.io/badge/go.dev-documentation-007d9c?&style=for-the-badge)](https://pkg.go.dev/github.com/MakMoinee/go-mith)
[![Latest Version](https://img.shields.io/github/tag/MakMoinee/go-mith.svg?&style=for-the-badge&label=semver)](https://github.com/MakMoinee/go-mith/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/MakMoinee/go-mith/ci.yml?style=for-the-badge&branch=main)](https://github.com/MakMoinee/go-mith/actions/workflows/ci.yml)
[![Code Coverage](https://img.shields.io/codecov/c/github/MakMoinee/go-mith/main.svg?style=for-the-badge)](https://codecov.io/github/MakMoinee/go-mith)

</div>

## Packages
- Palindrome - checks if the value given is palindrome or not
- Power Formula - One of the science formula. It's used to calculate the power from a given work and time values
- Stair Case (Hacker Rank Solution) - prints a staircase of size n.
- Concurrency Package - useable interface for any concurrent calls.
- Goserve Package - build http service to start your API with the support of injecting certs and reading config from settings.yaml
- Email - send email by using our prebuilt function, no need to code manually for email just instantiate the package and pass the required paramaters and it should work
- MithClientRest - a small, payload-driven REST client. You map each request payload type to an HTTP method and path through an OperationResolver, then just call Do/Exec - marshalling, headers, timeouts and error parsing are handled for you

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


## email package

```go
package main

import (
	"log"

	"github.com/MakMoinee/go-mith/pkg/email"
)

func main() {
	emailService := email.NewEmailService(587, "emailHost", "emailAddress", "emailAppPass")

	isEmailSent, err := emailService.SendEmail("receiverEmail", "emailSubject", "emailMessage")
	if err != nil {
		log.Fatalf("Error sending email: %s", err)
	}

	if isEmailSent {
		log.Println("Email Send Successfully")
	} else {
		log.Println("Failed to send email")
	}
}
```


## mithclientrest package

A lightweight REST client where the request payload itself decides which
operation is called. Implement an `OperationResolver` once, then send any
payload with `Do` (returns the raw response body) or `Exec` (discards it).

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MakMoinee/go-mith/pkg/mithclientrest"
)

// Payloads
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

// Resolver maps a payload type to an HTTP method and path.
type Resolver struct{}

func (Resolver) OperationFor(payload any) (mithclientrest.Operation, error) {
	switch payload.(type) {
	case CreateUserRequest:
		return mithclientrest.Operation{
			Method: http.MethodPost,
			Path:   "/users",
		}, nil
	case DeleteUserRequest:
		return mithclientrest.Operation{
			Method:  http.MethodDelete,
			Path:    "/users",
			Headers: map[string]string{"X-Confirm": "true"}, // per-operation headers
		}, nil
	default:
		return mithclientrest.Operation{}, fmt.Errorf("unsupported payload %T", payload)
	}
}

func main() {
	client := mithclientrest.New(&mithclientrest.Config{
		BaseURL: "https://api.example.com",
		Headers: map[string]string{
			"Authorization": "Bearer <token>",
		},
		Timeout:           10 * time.Second, // defaults to 30s when omitted
		OperationResolver: Resolver{},
		// HTTPClient: &http.Client{}, // optional, supply your own client
	})

	ctx := context.Background()

	// Do - returns the raw response body
	body, err := client.Do(ctx, CreateUserRequest{
		Name:  "Juan Dela Cruz",
		Email: "juan@example.com",
	})
	if err != nil {
		// Non 2xx responses come back as *mithclientrest.Fault
		var fault *mithclientrest.Fault
		if errors.As(err, &fault) {
			log.Fatalf("status=%d message=%s body=%s",
				fault.StatusCode, fault.Message, string(fault.Body))
		}
		log.Fatalf("request failed: %v", err)
	}

	var created struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(body, &created); err != nil {
		log.Fatalf("decode response: %v", err)
	}
	log.Println("created user:", created.ID)

	// Exec - when the response body is not needed
	if err := client.Exec(ctx, DeleteUserRequest{ID: created.ID}); err != nil {
		log.Fatalf("delete failed: %v", err)
	}
	log.Println("user deleted")
}
```

### Notes
- `Content-Type: application/json` is always set; `Config.Headers` are applied
  next and `Operation.Headers` last, so an operation can override a global header.
- Every payload is JSON-encoded and sent as the request body.
- Any response outside the 2xx range returns a `*mithclientrest.Fault` carrying
  the status code, the raw body, and a message extracted from a `message` or
  `error` JSON field when present.
