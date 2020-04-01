package skiplist

import (
	"go_dataStruct/util"
	"math/rand"
)

const (
	maxLevel = 32
	p        = 0.25
)

type (
	SkipList struct {
		size  int   // 容量
		level int   // 实际层数
		first *Node // 首节点，层数等于最高层数
	}

	Node struct {
		Key   interface{}
		Value interface{}
		Nexts []*Node
	}
)

func NewNode(k interface{}, v interface{}, level int) *Node {
	return &Node{
		Key:   k,
		Value: v,
		Nexts: make([]*Node, level),
	}
}

func New() *SkipList {
	return &SkipList{
		first: &Node{
			Nexts: make([]*Node, maxLevel),
		},
	}
}

func (sl *SkipList) IsEmpty() bool {
	return sl.size == 0
}

func (sl *SkipList) Size() int {
	return sl.size
}

func (sl *SkipList) Get(key util.Comparable) interface{} {
	keyChecking(key)
	node := sl.first
	for i := sl.level - 1; i >= 0; i-- {
		cmp := -1
		for node.Nexts[i] != nil {
			cmp = key.CompareTo(key, node.Nexts[i].Key)
			if cmp <= 0 {
				break
			}
			node = node.Nexts[i]
		}
		if cmp == 0 {
			return node.Nexts[i].Value
		}
	}
	return nil
}

func (sl *SkipList) Put(key util.Comparable, val interface{}) interface{} {
	keyChecking(key)
	node := sl.first
	prevs := make([]*Node, sl.level)

	for i := sl.level - 1; i >= 0; i-- {
		for node.Nexts[i] != nil && key.CompareTo(key, node.Nexts[i].Key) > 0 {
			node = node.Nexts[i]
		}
		if key.CompareTo(key, node.Nexts[i].Key) == 0 {
			oldV := node.Nexts[i].Value
			node.Nexts[i].Value = val
			return oldV
		}
		prevs[i] = node
	}

	newLevel := randomLevel()
	newNode := NewNode(key, val, newLevel)

	for i := 0; i < newLevel; i++ {
		if i >= sl.level {
			sl.first.Nexts[i] = newNode
		}
		newNode.Nexts[i] = prevs[i].Nexts[i]
		prevs[i].Nexts[i] = newNode
	}
	sl.size++
	if sl.level < newLevel {
		sl.level = newLevel
	}
	return nil
}

func (sl *SkipList) Remove(key util.Comparable) interface{} {
	keyChecking(key)
	node := sl.first
	prevs := make([]*Node, sl.level)
	var mth *Node
	for i := sl.level; i >= 0; i-- {
		for node.Nexts[i] != nil && key.CompareTo(key, node.Nexts[i].Key) > 0 {
			node = node.Nexts[i]
		}
		if key.CompareTo(key, node.Nexts[i].Key) == 0 {
			mth = node.Nexts[i]
		}

		prevs[i] = node
	}

	if mth == nil {
		return nil
	}
	sl.size--
	for i := 0; i < len(mth.Nexts); i++ {
		prevs[i].Nexts[i] = prevs[i].Nexts[i].Nexts[i]
	}

	for sl.level > 0 && sl.first.Nexts[sl.level-1] == nil {
		sl.level = sl.level - 1
	}

	return mth.Value
}

func keyChecking(key util.Comparable) {
	if key == nil {
		panic("key should not be a nil.")
	}
}

func randomLevel() int {
	level := 1
	for rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}
