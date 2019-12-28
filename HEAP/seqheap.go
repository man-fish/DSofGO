package HEAP

import "fmt"

/** https://www.bilibili.com/video/av47196993?from=search&seid=18158741109556862051 */

type SeqHeap struct {
	tree []int
	size int
}

func NewEmprySeqHrap() *SeqHeap {
	return &SeqHeap{}
}

func NewSeqHrap(arr []int) *SeqHeap {
	heap := NewEmprySeqHrap()
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


func (s *SeqHeap) heapifydown (tree []int,n int,i int) {//n是堆化范围，i是当前堆化元素
	c1 := 2*i + 1		//获取左子树
	c2 := 2*i + 2		//获取右子树
	max := i			//最大值
	if c1 < n &&tree[c1] > tree[max] {
		max = c1		//根据堆化范围交换最大元素
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	if (max != i) {
		swap(tree, max, i)
		s.heapifydown(tree,n,max)	//向下堆化
	}
}

func (s *SeqHeap) heapifyup(i int) {
	parent := (i-1)/2
	if parent >= 0 && s.tree[i] > s.tree[parent] {
		swap(s.tree,i,parent)
		s.heapifyup(parent)
	}
}

func (s *SeqHeap) buildHeap (tree []int,n int) {
	lastnode := n - 1
	parent := (lastnode-1)/2
	for i := parent; i >= 0; i-- {
		s.heapifydown(tree,n,i)
	}
}

func (s *SeqHeap) Insert(x int) {
	s.tree = append(s.tree,x)
	s.size = s.size + 1
	s.heapifyup(s.size -1)
}

func (s *SeqHeap) HeapSort() []int {
	for i := s.size-1; i > 0; i--  {
		s.heapifydown(s.tree, i,0)
		swap(s.tree,i,0)
	}
	defer s.buildHeap(s.tree,s.size)
	arr := make([]int,s.size)
	copy(arr,s.tree)
	return arr
}

func (s *SeqHeap) String() string {
	return fmt.Sprintf("heap ->%v",s.tree)
}