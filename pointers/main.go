package main

import "fmt"

func main() {
	var x = 1
	var y *int = &x

	fmt.Printf("address of x:  %p\n", y)
	fmt.Printf("value of x through y: %d\n", *y)

	*y = 2
	fmt.Printf("setting value of x through y: %d\n", x)
}
