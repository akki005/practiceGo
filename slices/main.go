package main

import "fmt"

func main() {

	// 	Slices when its passed it’s passed with the pointer to underlying array, a slice value only contains a pointer to the array where the elements are actually stored. The slice value does not include its elements (unlike arrays). slice value is a header , describing a contiguous section of a backing array, when you pass a slice to a function, a copy will be made from this header, including the pointer, which will point to the same backing array. Modifying the elements of the slice implies modifying the elements of the backing array, and so all slices which share the same backing array will “observe” the change.
	// A slice 3 is a small structure that points to an underlying array. The small structure is copied, but it still points to the same underlying array. the memory block containing the slice elements is passed by “reference”. The slice information triplet holding the capacity, the number of element and the pointer to the elements is passed by value.
	a := []int{1, 2, 3}
	b := [][]int{}
	fmt.Printf("before a=%v, b=%v\n", a, b)
	b = append(b, a)
	a[0] = 5
	fmt.Printf("after a=%v, b=%v\n", a, b)
	//swap elements
	a[0], a[1] = a[1], a[0]
	fmt.Printf("swap 0&1 =%v\n", a)

	//solution. copy slices
	c := make([]int, len(a))
	copy(c, a)
	fmt.Printf("before a=%v, c=%v\n", a, c)
	a[0] = 10
	fmt.Printf("after a=%v, c=%v\n", a, c)

	//basic operations
	fmt.Printf("before a=%v\n", a)
	a = append(a, 4)
	fmt.Printf("push 4 %v\n", a)
	a = a[:len(a)-1]
	fmt.Printf("pop %v\n", a)
	a = a[1:]
	fmt.Printf("dequeue %v\n", a)

	//insert at index
	x := []int{1, 2, 4, 5}
	pos := 2
	value := 3
	x = insertElementAtIndex(x, pos, value)
	fmt.Printf("x=%v", x)

}

func insertElementAtIndex(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	//a=[1,2,4,4,5]
	a[index] = value
	//a=[1,2,3,4,5]
	return a
}
