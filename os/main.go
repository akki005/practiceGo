package main

import (
	"fmt"
	"runtime"
)

func main() {
	//this will give us total number of logical cpus
	fmt.Println(runtime.NumCPU())

}
