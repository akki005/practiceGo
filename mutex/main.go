package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncMutext struct {
	mx      sync.Mutex
	counter int
}

func (s *SyncMutext) inc() {
	// Lock so only one goroutine at a time can access the value counter.
	s.mx.Lock()
	s.counter++
	s.mx.Unlock()
}

func (s *SyncMutext) read() int {
	// Lock so only one goroutine at a time can access the count.
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.counter
}

func main() {

	mutextNew := new(SyncMutext)

	//without mutext value of counter in SyncMutext would be randomly incremented as all go routines will share the same variable and they will try to read and write same variable concurrently so final value of counter is not predictable
	for i := 1; i <= 100; i++ {
		go mutextNew.inc()
	}

	time.Sleep(time.Second * 2)
	fmt.Println(mutextNew.read())

}
