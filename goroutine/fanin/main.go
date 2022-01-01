package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//number
func main() {

	size := 10

	ch1 := numberGenerator("me", size)
	ch2 := numberGenerator("you", size)

	finalCh := fanInWithForSelect(ch1, ch2)

	// for i := 0; i < size; i++ {
	// 	fmt.Println(<-finalCh)
	// }

	//here each message will have timeout
	for {
		select {
		case f := <-finalCh:
			log.Print(f)
		case <-time.After(1 * time.Second):
			log.Printf("to Slow!!!")
			return
		}
	}

}

/////

func fanInAsync(ch1, ch2 <-chan string) <-chan string {

	finalCh := make(chan string)

	go func() {
		for {
			finalCh <- <-ch1
		}
	}()

	go func() {
		for {
			finalCh <- <-ch2
		}
	}()

	return finalCh
}

//this is same as fanInAsync we are just using for select
func fanInWithForSelect(ch1, ch2 <-chan string) <-chan string {

	finalCh := make(chan string)

	go func() {
		for {
			select {
			case data1 := <-ch1:
				finalCh <- data1
			case data2 := <-ch2:
				finalCh <- data2
			}
		}
	}()

	return finalCh
}

func numberGenerator(id string, size int) <-chan string {
	numberCh := make(chan string)

	go func() {
		for i := 0; i < size; i++ {
			numberCh <- fmt.Sprintf("%d", i) + id
			time.Sleep(time.Duration(rand.Intn(2000) * int(time.Millisecond)))
		}
		close(numberCh)
	}()

	return numberCh
}
