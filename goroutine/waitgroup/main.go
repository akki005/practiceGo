package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go task(&wg, i)
	}

	wg.Wait()
}

func task(wg *sync.WaitGroup, taskId int) {
	time.Sleep(2 * time.Second)
	log.Printf("task %d is done", taskId)
	wg.Done()
}
