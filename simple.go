package simple

import (
	"fmt"
)

func A(val int) int {
	fmt.Println("Checking value")
	if val > 0 {
		return 1
	}
	return 0
}

func b(val int) int {
	return -1
}
