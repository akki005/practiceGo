package main

import (
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	startWorkers(5, &wg)
	wg.Wait()

}

func startWorkers(n int, wg *sync.WaitGroup) {

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			log.Println("i am worker ", i)
			wg.Done()
		}(i)
	}

}
