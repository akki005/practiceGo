package main

import "fmt"

type Vertext struct {
	X int
	Y int
}

func main() {
	v1 := new(Vertext)
	fmt.Printf("vertex %v\n", v1)
	v2 := Vertext{X: 10, Y: 20}
	fmt.Printf("vertex2 %v\n", v2)

	v3 := &v2
	v3.X = 2e5
	fmt.Printf("vertex2 %v\n", v2)
	fmt.Printf("vertex3 %v\n", v3)

}
