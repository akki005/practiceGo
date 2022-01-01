package main

import "fmt"

type I interface {
	Method()
}

type Obj struct {
}

func (o *Obj) Method() {
	if o == nil {
		fmt.Println("calling on nil")
	}
	fmt.Println("method")
}

func main() {
	EmptyInterface()
	InterfaceAssignments()
}

func EmptyInterface() {
	var x interface{}

	x = 100
	fmt.Printf("x type=%T,value=%v\n", x, x)

	x = 100.1
	fmt.Printf("x type=%T,value=%v\n", x, x)

}

func InterfaceAssignments() {
	//works
	var i I
	//below line wont work as i doesn't have any solid struct
	// i.Method()

	obj := new(Obj)
	i = obj
	i.Method()

	//works even if obj2 is nil
	var obj2 *Obj
	i = obj2
	i.Method()
}
