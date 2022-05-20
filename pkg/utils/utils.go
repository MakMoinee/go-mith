package utils

import "fmt"

func ValidField[T ~int | ~string | ~[]int | ~[]string](t T) {
	fmt.Println(fmt.Sprintf("%T", t))
}
