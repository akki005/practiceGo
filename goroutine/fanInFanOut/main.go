package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func randIntGenerator(done <-chan bool, limit int) <-chan int {

	numCh := make(chan int)

	go func() {
		defer close(numCh)
		for {
			select {
			case <-done:
				return
			case numCh <- rand.Intn(limit):
			}
		}
	}()

	return numCh
}

func primeProducer(done <-chan bool, numbers <-chan int) <-chan int {

	numCh := make(chan int)

	go func() {
		defer close(numCh)
		for {
			select {
			case <-done:
				return
			case num := <-numbers:
				if checkPrime(num) {
					numCh <- num
				}
			}
		}
	}()

	return numCh

}

func checkPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func taker(done <-chan bool, numbers <-chan int) <-chan int {

	numCh := make(chan int)

	go func() {
		defer close(numCh)
		for i := 0; i < 20; i++ {
			numCh <- <-numbers
		}
	}()

	return numCh

}

func fanInFinders(done <-chan bool, numChs ...<-chan int) <-chan int {

	var wg sync.WaitGroup

	finalCh := make(chan int)

	reader := func(numCh <-chan int) {
		defer wg.Done()
		for i := range numCh {
			select {
			case <-done:
				return
			case finalCh <- i:
			}
		}
	}

	wg.Add(len(numChs))
	for _, ch := range numChs {
		go reader(ch)
	}

	go func() {
		defer close(finalCh)
		wg.Wait()
	}()

	return finalCh
}

func main() {
	// runtime.GOMAXPROCS(2)
	//this is fan out at we are fanning out prime finding process as it it time consuming
	fanInFanOutPrimeFinder()
	// pipelinePrimeFinder()

}

func fanInFanOutPrimeFinder() {
	closeCh := make(chan bool)
	defer close(closeCh)

	start := time.Now()
	numOfPrimeFinders := runtime.NumCPU()

	finders := make([]<-chan int, numOfPrimeFinders)

	for i := 0; i < len(finders); i++ {
		finders[i] = primeProducer(closeCh, randIntGenerator(closeCh, 200))
	}

	for num := range taker(closeCh, fanInFinders(closeCh, finders...)) {
		log.Println(num)
	}
	log.Println("took", time.Since(start).Seconds())
}

func pipelinePrimeFinder() {
	start := time.Now()

	closeCh := make(chan bool)
	defer close(closeCh)
	log.Println("started pipeline")

	for num := range taker(closeCh, primeProducer(closeCh, randIntGenerator(closeCh, 5000000000))) {
		log.Println(num)
	}
	log.Println("took", time.Since(start).Seconds())

}
