package main

import "fmt"

//number
func main() {

	numberP := numberProducer(2, 3)
	sqNumberP := squareProducer(numberP)

	for num := range sqNumberP {
		fmt.Println(num)
	}

	for num := range squareProducer(squareProducer(numberProducer(4, 5))) {
		fmt.Println(num)
	}
}

//The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
func numberProducer(nums ...int) <-chan int {
	numbers := make(chan int)

	go func() {
		defer close(numbers)
		for _, num := range nums {
			numbers <- num
		}
	}()

	return numbers
}

//The second stage, sq, receives integers from a channel and returns a channel that emits the square of each received integer.
func squareProducer(numCh <-chan int) <-chan int {
	sqNumbers := make(chan int)

	go func() {
		defer close(sqNumbers)
		for num := range numCh {
			sqNumbers <- num * num
		}
	}()

	return sqNumbers
}
