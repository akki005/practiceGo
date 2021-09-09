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

	//solution. copy slices
	c := make([]int, len(a))
	copy(c, a)
	fmt.Printf("before a=%v, c=%v\n", a, c)
	a[0] = 10
	fmt.Printf("after a=%v, c=%v\n", a, c)

}
