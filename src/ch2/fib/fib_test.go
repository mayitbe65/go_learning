package fib

import (
	"fmt"
	"testing"
)

func TestFibFunc(t *testing.T) {
	var a int = 1
	var b int = 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}
