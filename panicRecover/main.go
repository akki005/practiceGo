package main

import (
	"log"
	"runtime/debug"
)

func main() {

	defer log.Print("i will print no matter what")

	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
			debug.PrintStack()
		}
	}()

	panic("here panincing")
}
