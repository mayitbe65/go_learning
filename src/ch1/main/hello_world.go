package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(len(os.Args))
	fmt.Println("hello world!")
	os.Exit(-1)
}
