package main

import (
	"fmt"
	"time"
)

func Logger(number int) {
	time.Sleep(1 * time.Second)
	fmt.Println(number)
}

func main() {

	var sem = make(chan int, 5)
	//we are rate limiting logger method here. there shouldn't be more than 5 concurrent logger instances
	for i := 0; i < 10; i++ {
		sem <- 1
		go func(number int) {
			Logger(number)
			<-sem
		}(i)
	}
	time.Sleep(5 * time.Second)
}
