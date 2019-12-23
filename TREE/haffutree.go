package TREE

import (
	"fmt"
	"strconv"
)

type HaffeeTree struct{
	charset		string
	triELes		[]*TriNode
}

func NewEmptyHaffTree() *HaffeeTree {
	return &HaffeeTree{}
}

func NewHaffTree(weights []int) *HaffeeTree {
	ht := NewEmptyHaffTree()
	charset := ""
	for idx := range weights {
		charset += string('A'+idx)
	}
	//构造编码字符集
	ht.charset = charset
	ht.triELes = make([]*TriNode,2*len(weights)-1)
	//哈夫曼结点数组
	for i := 0; i < len(weights); i++ {
		ht.triELes[i] = NewEmptyTriNode(weights[i])
	}
	//初始化所有子叶结点
	for j := len(weights); j < 2*len(weights)-1; j++ {
		min1 := 1000
		min2 := min1
		x1 := -1
		x2 := -1
		for l := 0; l < j ; l++  {
			if ht.triELes[l].Parent == -1 {
				if ht.triELes[l].Data < min1 {
					min2 = min1
					x2 = x1
					min1 = ht.triELes[l].Data
					x1 = l
				}else if ht.triELes[l].Data < min2 {
					min2 = ht.triELes[l].Data
					x2 = l
				}
			}
		}
		ht.triELes[x1].Parent = j
		ht.triELes[x2].Parent = j
		ht.triELes[j] = NewTriNode(min1+min2,-1,x1,x2)
	}
	return ht
}

func (ht *HaffeeTree) GetCode(i int) string {
	parent := ht.triELes[i].Parent
	if parent == -1 {
		return ""
	}
	if ht.triELes[parent].Left == i {
		return ht.GetCode(parent) + strconv.Itoa(0)
	}else {
		return ht.GetCode(parent) + strconv.Itoa(1)
	}
}
//if ht.triELes[parent].Left == i {
//	return 0
//}else {
//	return 1
//}
func (ht *HaffeeTree) String() string {
	str := ""
	for i := 0; i < len(ht.charset); i++ {
		str += string([]rune(ht.charset)[i]) + ":" + ht.GetCode(i) + "\n"
	}
	return str
}

func (ht *HaffeeTree) Encode(text string) string {
	compressed := ""
	for _,item := range []rune(text) {
		compressed += ht.GetCode(int(item-'A'))
	}
	return compressed
}

func (ht *HaffeeTree) Decode(compressed string) string {
	text := ""
	p := len(ht.triELes)-1
	fmt.Println(p)
	for idx, item := range compressed {
		if item == '0' {
			p = ht.triELes[idx].Left
		}else{
			p = ht.triELes[idx].Right
		}
		if ht.triELes[idx].IsLeaf() {
			text += string([]rune(ht.charset)[idx])
			p = len(ht.triELes)-1
		}
	}
	return text
}

