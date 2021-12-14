package main

import (
	"fmt"
	"time"
)

func main() {
	//continuously ticking timer ever 2 seconds
	ticker := time.NewTicker(2 * time.Second)

	for tick := range ticker.C {
		fmt.Println("Tick at", tick)
	}
}
