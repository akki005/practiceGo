package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int64 = 10
	fmt.Println("x", strconv.FormatInt(x, 2))
	x = x << 1
	fmt.Println("x << 1", strconv.FormatInt(x, 2))
	x = x << 1
	fmt.Println("x << 1", strconv.FormatInt(x, 2))
}
