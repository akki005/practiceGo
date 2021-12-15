package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

func main() {
	var x interface{}

	x = 100

	fmt.Printf("x: type=%v, value=%v\n", reflect.TypeOf(x), reflect.ValueOf(x))

	u := user{
		name: "Ak",
		age:  29,
	}

	uType := reflect.TypeOf(u)
	uVal := reflect.ValueOf(u)

	totalFields := uType.NumField()

	for i := 0; i < totalFields; i++ {
		fmt.Printf(" %v : %v : %v\n", uType.Field(i).Name, uType.Field(i).Type, uVal.Field(i))
	}

	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}

	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	fmt.Printf("a: kind %v\n", aType.Kind())
	fmt.Printf("a: kind %v\n", bType.Kind())

}
