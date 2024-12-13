package main

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type Bucket struct {
	mu   sync.RWMutex
	node *Node
}

func NewBucket() *Bucket {
	return &Bucket{}
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func NewBucketNode(key, value int) *Bucket {
	return &Bucket{
		node: NewNode(key, value),
	}
}

type ConcurrentHashMap struct {
	buckets []*Bucket
}

func Constructor() ConcurrentHashMap {
	// Ensure that all buckets are initialized during the ConcurrentHashMap creation. This way, we avoid the need to check for nil and initialize buckets dynamically during Put or Get. otherwise we have to add mutex on ConcurrentHashMap also.
	buckets := make([]*Bucket, 100000)

	for i := 0; i < 100000; i++ {
		buckets[i] = NewBucket()
	}

	return ConcurrentHashMap{
		buckets: buckets,
	}
}

func (this *ConcurrentHashMap) Put(key int, value int) {
	index := key % len(this.buckets)

	bucket := this.buckets[index]

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	currentNode := bucket.node

	if currentNode == nil {
		bucket.node = NewNode(key, value)
		return
	}

	prevNode := currentNode

	for currentNode != nil {

		if currentNode.Key == key {
			currentNode.Value = value
			return
		}

		prevNode = currentNode
		currentNode = currentNode.Next
	}

	prevNode.Next = NewNode(key, value)
}

func (this *ConcurrentHashMap) Get(key int) int {
	index := key % len(this.buckets)

	bucket := this.buckets[index]

	if bucket == nil {
		return -1
	}

	// Lock the bucket for reading
	bucket.mu.RLock()
	defer bucket.mu.RUnlock()

	currentNode := bucket.node

	for currentNode != nil {
		if currentNode.Key == key {
			return currentNode.Value
		}
		currentNode = currentNode.Next
	}

	return -1
}

func (this *ConcurrentHashMap) Remove(key int) {
	index := key % len(this.buckets)

	currentNode := this.buckets[index].node

	if currentNode == nil {
		return
	}

	var prevNode *Node = nil

	for currentNode != nil && currentNode.Key != key {
		prevNode = currentNode
		currentNode = currentNode.Next
	}

	if prevNode == nil {
		// remove head;
		this.buckets[index].node = currentNode.Next
		currentNode = nil
		return
	}

	prevNode.Next = currentNode.Next
	currentNode = nil
}

func main() {
	hashMap := Constructor()

	// hashMap.Put(1, 1)
	// fmt.Println(hashMap.Get(1))
	// hashMap.Put(2, 2)
	// fmt.Println(hashMap.Get(2))
	// hashMap.Put(1, 10)
	// fmt.Println(hashMap.Get(1))
	// hashMap.Put(4, 4)
	// fmt.Println(hashMap.Get(4))
	// hashMap.Put(5, 5)
	// fmt.Println(hashMap.Get(5))
	// hashMap.Put(2, 11)
	// fmt.Println(hashMap.Get(2))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sleeping 1")
		time.Sleep(time.Millisecond * 5)
		hashMap.Put(1, 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sleeping 2")
		time.Sleep(time.Millisecond * 5)
		hashMap.Put(1, 10)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sleeping 3")
		time.Sleep(time.Millisecond * 5)
		fmt.Println(hashMap.Get(1))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sleeping 4")
		time.Sleep(time.Millisecond * 5)
		fmt.Println(hashMap.Get(1))
	}()

	wg.Wait()

	fmt.Println(hashMap.Get(1))
}
