package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int64 = 10
	fmt.Println("x", strconv.FormatInt(x, 2), x)
	x = x << 1
	fmt.Println("x << 1", strconv.FormatInt(x, 2), x)
	x = x << 1
	fmt.Println("x << 1", strconv.FormatInt(x, 2), x)
	fmt.Println("x << 1", 0^4^1^2^1^2)
	fmt.Println("x << 1", 0^2^2^1)
}
