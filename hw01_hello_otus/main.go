package main

import "golang.org/x/example/hello/reverse"

func main() {
	originalString := "Hello, Otus!"
	reverseString := reverse.String(originalString)
	print(reverseString)
}
