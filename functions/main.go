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

	adder := adder()
	fmt.Println(adder(5))
	fmt.Println(adder(2))
	fmt.Println(adder(3))

	fmt.Println("fib generator")
	fibgen := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibgen())
	}
}

//Function closures
func adder() func(number int) int {
	sum := 0
	return func(number int) int {
		sum += number
		return sum
	}
}

//Fibonacci closure
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		new := a
		a, b = b, a+b
		return new
	}
}
