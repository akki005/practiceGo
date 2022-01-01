package main

import (
	"log"
	"time"
)

func main() {

	log.Println("started")

	//main routine should exists if either of ch2 or ch2 exit
	ch1 := afterTimeGenerator(10)
	ch2 := afterTimeGenerator(5)

	<-closer(ch1, ch2)
	log.Println("done")
}

func closer(ch1, ch2 <-chan bool) <-chan bool {

	doneCh := make(chan bool)

	go func() {
		defer close(doneCh)
		select {
		case <-ch1:
		case <-ch2:
		}
	}()

	return doneCh
}

func afterTimeGenerator(duration int) <-chan bool {

	expireCh := make(chan bool)
	go func() {
		defer close(expireCh)
		time.Sleep(time.Duration(duration) * time.Second)
	}()
	return expireCh
}
