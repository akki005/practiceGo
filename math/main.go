package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("MaxInt16 value=%v, type=%T\n", math.MaxInt64, math.Max)
	fmt.Printf("MaxInt16 value=%v, type=%T\n", math.MinInt64, math.MinInt64)
	fmt.Println(4/2, 5/2)
	//parse to int
	fmt.Println(int(11 / 10)) //discards fractional value
	x := 235
	y := int(1e9 + 7) //10^9+7
	fmt.Println(x % y)
}
