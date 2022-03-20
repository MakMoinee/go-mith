package number

import (
	"fmt"
)

func IsNumberMatchLength(n int, length int) bool {
	str := fmt.Sprintf("%v", n)
	if len(str) == length {
		return true
	}
	return false
}
