package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("i am defer func")
	}()

	defer fmt.Println("i am defer print")

	a := 1
	defer fmt.Println("a at defer:", a)
	a = 2

	for i := 0; i < 5; i++ {
		defer fmt.Println("in loop", i)
	}
}
