package manipulate

import (
	"fmt"
	"strings"
)

// GetStairCase prints a staircase of size .
func GetStairCase(n int32) {
	count := int(n)
	for i := 1; i <= int(n); i++ {
		count--
		fmt.Print(strings.Repeat(" ", count))
		fmt.Println(strings.Repeat("#", i))

	}
}
