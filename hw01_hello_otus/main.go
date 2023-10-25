package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	originalString := "Hello, OTUS!"
	reverseString := reverse.String(originalString)
	fmt.Println(reverseString)
}
