package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	size := 10
	numberCh := numberGenerator(size)
	numberCh2 := numberGenerator(size)

	for i := 0; i < size; i++ {
		//both will print whenever one is available. i.e numberCh will print and it might wait for some time before printing numberCh2
		log.Print(<-numberCh)
		log.Print(<-numberCh2)

	}

	numberChP := numberGeneratorWithPause("me", size)
	numberCh2P := numberGeneratorWithPause("you", size)

	for i := 0; i < size; i++ {
		//both will be printed exactly at same time because both are waiting other to complete. even if they have a random time
		msg1 := <-numberChP
		log.Print(msg1.str)
		msg2 := <-numberCh2P
		log.Print(msg2.str)
		//mark current message as done so next can be read
		msg1.pause <- true
		msg2.pause <- true

	}

}

func numberGenerator(size int) <-chan int {
	numberCh := make(chan int)

	go func() {
		for i := 0; i < size; i++ {
			time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
			numberCh <- i
		}
		close(numberCh)
	}()

	return numberCh
}

////////////////////////////////////////////////////////
type Message struct {
	str   string
	pause chan bool
}

func numberGeneratorWithPause(id string, size int) <-chan Message {
	numberCh := make(chan Message)
	waitCh := make(chan bool)

	go func() {
		for i := 0; i < size; i++ {
			numberCh <- Message{fmt.Sprintf("%d", i) + id, waitCh}
			time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
			<-waitCh
		}
		close(numberCh)
	}()

	return numberCh
}
