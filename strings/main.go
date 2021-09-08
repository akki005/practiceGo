package main

import "fmt"

//ref:
//https://www.dotnetperls.com/substring-go

func main() {
	//string in go are represented in memory using UTF-8 encoding.
	s := "Hello World"

	fmt.Printf("value=%s , type=%T\n", s, s)
	//len(s) doesn't give us length of string it gives number of bytes required to represent a string
	fmt.Printf("size=%v\n", len(s))
	fmt.Printf("s[1]=%v\n", s[1])
	fmt.Printf("char s[1]=%v\n", string(s[1]))

	b := []byte(s)
	fmt.Printf("value=%v, Type=%T\n", b, b)

	//rune will contain only one character which are valid 'UTF-8'
	//rune is big enough to represent any printable character
	//rune is just an alias for int32
	var r rune = 't'
	fmt.Printf("value=%v, Type=%T\n", r, r)

	//following will print UTF-8 code instead of 'A'
	fmt.Println('A')
	fmt.Println('A' + 20)
	fmt.Printf("65=%c\n", 65)
	fmt.Printf("65=%d\n", 65)
	fmt.Printf("65=%X\n", 65) // exadecimal format

	//string to rune
	// For non-ASCII characters, more than 1 byte is required for a character. To take substrings correctly on non-ASCII strings, we must use runes.
	rn := []rune(s)
	fmt.Printf("rn[0]= %v\n", rn[0])             // this gives us utf-8 code
	fmt.Printf("rn[0]= %v\n", string(rn[0]))     // this gives us actual string
	fmt.Printf("rn[0:5]= %v\n", string(rn[0:5])) // to get substring

}
