// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

type item struct {
	key  int
	freq int
}

type priorityQ []*item

func (p priorityQ) Len() int {
	return len(p)
}

func (p priorityQ) Less(i, j int) bool {
	return p[i].freq > p[j].freq
}

func (p priorityQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQ) Push(x interface{}) {
	item := x.(*item)
	*p = append(*p, item)
}

func (p *priorityQ) Pop() interface{} {
	cur := *p
	toBePopped := cur[len(cur)-1]
	*p = cur[:len(cur)-1]
	return toBePopped
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[int]int{
		1: 3, 4: 2, 3: 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(priorityQ, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &item{
			key:  value,
			freq: priority,
		}
		i++
	}
	heap.Init(&pq)

	fmt.Println(pq[1])

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*item)
		fmt.Printf("%d:%v ", item.key, item.freq)
	}
}
