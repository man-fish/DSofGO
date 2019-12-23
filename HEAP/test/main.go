package main

import (
	"fmt"
	"go_DataStruct/HEAP"
)

func main() {
	heap := HEAP.NewSeqHrap([]int{2,3,19,4,11,5})
	fmt.Println(heap)
	fmt.Println(heap.HeadSort())
}
