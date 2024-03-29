package main

import (
	"fmt"
	"time"
)

func main() {
	// basicChannels()
	bufferedChannels()
	ch := make(chan string, 2)

	//send only channel
	var sendOnly chan<- string = ch
	sendOnly <- "akash"
	sendOnly <- "patel"

	//receive only channel
	var receiveOnly <-chan string = ch
	fmt.Println(<-receiveOnly)
	fmt.Println(<-receiveOnly)
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
