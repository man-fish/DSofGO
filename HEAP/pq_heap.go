package heap

import (
	"fmt"
	"go_DataStruct/util"
)

// PqHeap priority queue heap
// 优先队列是逻辑存储结构，首先要满足常规队列的队头出队，队尾入队。
// 其次还要满足每一次出队的是最大元素，入队之后保证最大元素优先结构。
// 我们只在队头和队尾进行swim或者sink保证对头元素符合最大堆。
type PqHeap struct {
	tree []interface{} // 第一位不存数，数组的实际储存容量是size+1，效果是可以让序号从一开始。
	size int           // size映射数组的实际存储容量
}

// NewPqHeap struct builder
func NewPqHeap(n int) PriorityQueue {
	return &PqHeap{
		tree: make([]interface{}, n+1),
	}
}

// Tree set for tree
func (pq *PqHeap) Tree() {
	fmt.Println(pq.tree...)
}

// Insert insert a new ele and keep the queue order
func (pq *PqHeap) Insert(value interface{}) {
	pq.size++
	pq.tree[pq.size] = value
	pq.swim(pq.size)
}

// Max return Max ele
func (pq *PqHeap) Max() interface{} {
	return pq.tree[1]
}

// DelMax Del the Max value and keep the queue order
func (pq *PqHeap) DelMax() interface{} {
	max := pq.tree[1]
	pq.swap(1, pq.size)
	pq.size--
	pq.sink(1)
	pq.tree[pq.size+1] = nil
	return max
}

// IsEmpty is queue empty
func (pq *PqHeap) IsEmpty() bool {
	return pq.size == 0
}

// Size setter of pq.size
func (pq *PqHeap) Size() int {
	return pq.size
}

// swim func to build heap
func (pq *PqHeap) swim(index int) {
	for index > 1 && pq.less(index/2, index) {
		pq.swap(index/2, index)
		index = index / 2
	}
}

// sink func to build heap，
func (pq *PqHeap) sink(index int) {
	for index*2 <= pq.size {
		k := index * 2
		if k < pq.size && pq.less(k, k+1) {
			// 选出将较大的元素上浮
			k++
		}
		if pq.less(k, index) {
			break
		}
		pq.swap(index, k)
		// 当前元素下沉
		index = k
	}
}

func (pq *PqHeap) less(i, j int) bool {
	return pq.tree[i].(util.Comparable).CompareTo(pq.tree[i], pq.tree[j]) == -1
}

func (pq *PqHeap) swap(i, j int) {
	pq.tree[i], pq.tree[j] = pq.tree[j], pq.tree[i]
}
