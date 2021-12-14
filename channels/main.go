package main

import (
	"fmt"
	"time"
)

func main() {
	// basicChannels()
	bufferedChannels()
}

func bufferedChannels() {
	chn := make(chan int, 3)
	chn <- 1
	fmt.Println(<-chn)
}

func basicChannels() {
	chn := make(chan int)

	go send(chn)
	go receive(chn)
	time.Sleep(5 * time.Second)
}

func send(c chan int) {
	//this will block in below line until channel reader is not ready on line fmt.Println(<-c)
	//observe the output of this code
	c <- 1
}

func receive(c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("timed out")
	fmt.Println(<-c)
}
