package main

import (
	"fmt"
	"unicode"
)

func main() {
	c := 'A'
	fmt.Println("isLeeter=", unicode.IsLetter(c))
}
