package main

import (
	"fmt"
	"go_DataStruct/HEAP"
)

func main() {
	heap := HEAP.NewSeqHrap([]int{2,3,19,4,11,5})
	fmt.Println(heap)
	heap.Insert(100)
	fmt.Println(heap)
	arr := heap.HeapSort()
	fmt.Println(arr)
	fmt.Println(heap)
}
