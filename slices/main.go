package main

import (
	"fmt"
	"sort"
)

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
	fmt.Printf("add 100 to front %v\n", append([]int{100}, a...))

	//insert at index
	x := []int{1, 2, 4, 5}
	pos := 2
	value := 3
	x = insertElementAtIndex(x, pos, value)
	fmt.Printf("x=%v\n", x)

	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	nums1 = nums2
	nums1[0] = 100
	fmt.Println(nums1, nums2)

	/**
	Important
	*/
	// 	Pop from queue

	// x, a = a[0], a[1:]

	// Pop from stack

	// x, a = a[len(a)-1], a[:len(a)-1]
	// Push

	// a = append(a, x)

	sort.Slice(nums1, func(i, j int) bool {
		fmt.Println(nums1[i], nums1[j], nums1[i] < nums1[j], nums1)
		//during iteration i points to next value and j points to current value
		//swap i and j if nums[i]<nums[j]
		//final state if nums[i] current number will be less than next number nums[j]
		return nums1[i] < nums1[j]
	})

	fmt.Println("are same? [3]int{1, 2, 3} == [3]int{1, 2, 3}", [3]int{1, 2, 3} == [3]int{1, 2, 3})
	fmt.Println("are same? [3]int{1, 2, 3} == [3]int{1, 2, 4}", [3]int{1, 2, 3} == [3]int{1, 2, 4})

	//slice capacity
	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	//https://go.dev/tour/moretypes/11

	sl := make([]int, 0, 5)
	printSlice("sl", sl)
	sl = sl[:2] //changes length only as we are droping item from end
	printSlice("sl", sl)
	sl = sl[1:] //changes length and capacity as we are dropping item from start
	printSlice("sl", sl)
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

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
