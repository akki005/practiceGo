package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	defer close(done)

	go func() {
		//sequential
		prod := producer()
		consumer(prod)
		done <- true
	}()

	go func() {
		size := 5
		//parallel
		prod2 := producerParallel(size)
		consumerParallel(prod2, size)
		done <- true
	}()

	for i := 0; i < 2; i++ {
		fmt.Println(<-done)
	}
}

func producer() <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		for i := 0; i <= 5; i++ {
			task()
			result <- i
		}
	}()

	return result

}

func consumer(intCh <-chan int) {
	for number := range intCh {
		fmt.Println("seq", number)
	}
}

func producerParallel(size int) chan int {
	result := make(chan int)

	for i := 0; i <= size; i++ {
		go func(number int) {
			task()
			result <- number
		}(i)
	}

	return result

}

func consumerParallel(intCh chan int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println("parallel", <-intCh)
	}
}

func task() {
	time.Sleep(2 * time.Second)
}
