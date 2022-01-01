package main

import "fmt"

var x, y = 1, 2

func dummy() {
	var x = 2
	var y = 3

	fmt.Println(x, y)
}

func main() {
	fmt.Println(x, y)
	dummy()
}
