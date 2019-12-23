package HEAP

import "fmt"

/** https://www.bilibili.com/video/av47196993?from=search&seid=18158741109556862051 */

type SeqHeap struct {
	tree []int
	size int
}

func NewSeqHrap(arr []int) *SeqHeap{
	heap := &SeqHeap{}
	heap.tree = arr
	heap.size = len(arr)
	heap.buildHeap(heap.tree,heap.size)
	return heap
}

func swap(arr []int,i int,j int) {
	arr[i] = arr[i]^arr[j]
	arr[j] = arr[j]^arr[i]
	arr[i] = arr[i]^arr[j]
}

func (s *SeqHeap) heapify (tree []int,n int,i int) {
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n &&tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	if (max != i) {
		swap(tree, max, i)
		s.heapify(tree,n,max)
	}
}

func (s *SeqHeap) buildHeap (tree []int,n int) {
	lastnode := n - 1
	parent := (lastnode-1)/2
	for i := parent; i >= 0; i-- {
		s.heapify(tree,n,i)
	}
}

func (s *SeqHeap) HeadSort() []int {
	var sortedArr []int
	for i := s.size-1; i >= 0; i--  {
		s.buildHeap(s.tree,i)
		swap(s.tree,i-1,0)
		sortedArr = append(sortedArr,s.tree[i-1])
	}
	return sortedArr
}

func (s *SeqHeap) String() string {
	return fmt.Sprintf("heap ->%v",s.tree)
}