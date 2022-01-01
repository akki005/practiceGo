package main

import (
	"fmt"

	"github.com/akki005/practiceGo/stringer/types"
)

type Currency float64

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", c)
}

func main() {
	pk1 := types.Aspirin
	fmt.Println(pk1)

	var balance Currency = 10.20

	fmt.Println(balance)
}
