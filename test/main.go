package main

import "fmt"

func main() {
	rn := []rune("123")

	fmt.Println(rn, '6' < '7')

	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl[1:])

	s := "abc"
	subStr := ""
	for _, char := range s {
		subStr += string(char)
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

	fmt.Printf("%T\n", s[0]-'A')
	var num int = 123
	fmt.Println(num / 10)
}
