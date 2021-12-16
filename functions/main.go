package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("i will run on exit")
	}()

	fn := func() {
		fmt.Println("this is fn")
	}

	fn()

	func() {
		fmt.Println("Anonymous fn")
	}()

}
