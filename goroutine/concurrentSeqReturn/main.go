package main

import (
	"fmt"
)

func main() {

	chang1 := make(chan string)
	chang2 := make(chan string)
	chang3 := make(chan string)

	channels := []chan string{chang1, chang2, chang3}

	for i := 1; i < 10; i += 3 {

		go g1(i, channels[i%3])
		go g2(i+1, channels[(i+1)%3])
		go g3(i+2, channels[(i+2)%3])

		fmt.Print(<-channels[i%3])
		fmt.Print(<-channels[(i+1)%3])
		fmt.Print(<-channels[(i+2)%3])
	}

}

func g1(i int, chanel chan string) {

	chanel <- fmt.Sprintln("g1", i)

}

func g2(i int, chanel chan string) {

	chanel <- fmt.Sprintln("g2", i)

}

func g3(i int, chanel chan string) {

	chanel <- fmt.Sprintln("g3", i)

}
