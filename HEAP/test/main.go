package main

import (
	"fmt"
	heap "go_DataStruct/HEAP"
)

type comInt int

func (i comInt) IsEqual(a, b interface{}) bool {
	return true
}
func (i comInt) CompareTo(a, b interface{}) int {
	if a.(comInt) > b.(comInt) {
		return 1
	} else if a.(comInt) < b.(comInt) {
		return -1
	} else {
		return 0
	}
}

func main() {
	// heap := heap.NewSeqHrap([]int{2, 3, 19, 4, 11, 5})
	// fmt.Println(heap)
	// heap.Insert(100)
	// fmt.Println(heap)
	// arr := heap.HeapSort()
	// fmt.Println(arr)
	// fmt.Println(heap)

	pq := heap.NewPqHeap(10)
	var a comInt = 1
	var b comInt = 3
	var c comInt = 5

	pq.Insert(a)
	pq.Insert(b)
	pq.Insert(c)

	fmt.Println(pq.DelMax())
	fmt.Println(pq.DelMax())
}
