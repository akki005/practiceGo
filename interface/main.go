package main

import "fmt"

func main() {

	var x interface{}

	x = 100
	fmt.Printf("x type=%T,value=%v\n", x, x)

	x = 100.1
	fmt.Printf("x type=%T,value=%v\n", x, x)

}
