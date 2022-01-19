package main

import (
	"fmt"
	"time"
)

func main() {
	batches := passInBatches(acceptAll(100))

	for batch := range batches {
		time.Sleep(2 * time.Second)
		fmt.Println(batch)
	}

}

func acceptAll(number int) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)
		for i := 1; i <= number; i++ {
			valueStream <- i
		}
	}()

	return valueStream
}

func passInBatches(valueStream <-chan interface{}) <-chan []int {
	batchCh := make(chan []int)

	go func() {
		defer close(batchCh)
		batch := []int{}
		for val := range valueStream {
			batch = append(batch, val.(int))
			if len(batch) == 10 {
				batchCh <- batch
				batch = []int{}
			}
		}
	}()

	return batchCh
}
