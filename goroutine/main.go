package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkTree(t, ch)
	close(ch)
}

func WalkTree(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkTree(t.Left, ch)
	ch <- t.Value
	WalkTree(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	treeCh1 := make(chan int)
	treeCh2 := make(chan int)

	go Walk(t1, treeCh1)
	go Walk(t2, treeCh2)
	for {
		node1, ok1 := <-treeCh1
		node2, ok2 := <-treeCh2
		if ok1 != ok2 || node1 != node2 {
			return false
		}
		if !ok1 {
			break
		}
	}

	return true
}

func SquaresProducer(done <-chan bool) chan int {

	outCh := make(chan int)

	go func() {
		defer close(outCh)
		count := 1
		// your code here
		for {
			select {
			case outCh <- count * count:
				count++
			case <-done:
				return
			}
		}
	}()

	return outCh
}

func SumCalculator(doneCh chan<- bool, squaresChannel <-chan int) chan int {

	outCh := make(chan int)

	go func() {
		defer close(outCh)
		sum := 0
		// your code here
		for i := 1; i <= 10; i++ {
			sum += <-squaresChannel
		}
		doneCh <- true
		outCh <- sum
	}()

	return outCh
}

func SumOfSquaresNew() {
	doneCh := make(chan bool)
	sum := SumCalculator(doneCh, SquaresProducer(doneCh))
	for val := range sum {
		fmt.Println(val)
	}
}

func main() {
	SumOfSquaresNew()
	SameTreeWalker()
}

func SameTreeWalker() {
	fmt.Println("done sum of squares")

	treeCh := make(chan int)

	go Walk(tree.New(1), treeCh)

	for t := range treeCh {
		fmt.Println(t)
	}

	fmt.Println(Same(tree.New(1), tree.New(2)))
}
