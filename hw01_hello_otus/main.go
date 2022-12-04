package main

import "golang.org/x/example/stringutil"

func main() {
	originalString := "Hello, Otus!"
	reverseString := stringutil.Reverse(originalString)
	print(reverseString)
}
