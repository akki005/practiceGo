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

func SumOfSquares(c, quit chan int) {
	count := 1
	// your code here
	for {
		select {
		case c <- count * count:
			count++
		case <-quit:
			return
		}
	}

}

func main() {

	mychannel := make(chan int)

	quitchannel := make(chan int)

	sum := 0

	go func() {
		for i := 0; i < 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 1
	}()
	SumOfSquares(mychannel, quitchannel)

	fmt.Println("done sum of squares")

	treeCh := make(chan int)

	go Walk(tree.New(1), treeCh)

	for t := range treeCh {
		fmt.Println(t)
	}

	fmt.Println(Same(tree.New(1), tree.New(2)))

}
