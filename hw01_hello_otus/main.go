package main

import "golang.org/x/example/hello/reverse"

func main() {
	originalString := "Hello, OTUS!"
	reverseString := reverse.String(originalString)
	print(reverseString)
}
