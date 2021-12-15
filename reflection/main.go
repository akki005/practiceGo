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

	fmt.Printf("x = 100, type=%v, value=%v\n", reflect.TypeOf(x), reflect.ValueOf(x))

	u := user{
		name: "Ak",
		age:  29,
	}

	uType := reflect.TypeOf(u)
	uVal := reflect.ValueOf(u)

	totalFields := uType.NumField()

	fmt.Println(`
user{
	name: "Ak",
	age:  29,
}
	`)

	fmt.Println()

	for i := 0; i < totalFields; i++ {
		fmt.Printf("%v : %v : %v\n", uType.Field(i).Name, uType.Field(i).Type, uVal.Field(i))
	}

	fmt.Println()

	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}

	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	fmt.Printf("a: kind %v\n", aType.Kind())
	fmt.Printf("a: kind %v\n", bType.Kind())

	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}

	u1 := &user{
		name: "Ak",
		age:  29,
	}

	u2 := &user{
		name: "Ak",
		age:  29,
	}

	fmt.Printf("deep equal a1==b1? %v\n", reflect.DeepEqual(a1, b1))
	fmt.Printf("deep equal u1==u2? %v\n", reflect.DeepEqual(u1, u2))
	fmt.Printf("simple equal u1==u2? %v\n", u1 == u2)

}
