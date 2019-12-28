package MATRIX

import (
	"go_DataStruct/LIST"
)

type LinkedMatrix struct {
	rows,cols 	int
	RowList		*LIST.SeqList

}

func NewEmptyLinkedMatrix(m,n int) *LinkedMatrix {
	if m > 0 && n > 0 {
		lm := &LinkedMatrix{rows:m,cols:n,RowList:LIST.NewEmptySeqList(m)}
		for i := 0; i < m; i++ {
			lm.RowList.Insert(i,LIST.NewEmptySingleList())
		}
		return lm
	}else{
		panic("not valid matrix form")
	}
}

func NewLinkedMatrix(m,n int, triples []Triple) *LinkedMatrix {
	lm := NewEmptyLinkedMatrix(m,n)
	for _,triple := range triples {
		lm.Set(triple)
	}
	return lm
}

func (lm *LinkedMatrix) Set(triple Triple) {
	row := triple.Row
	col := triple.Col
	if col < lm.cols && row < lm.rows && row >= 0 && col >= 0 {
		link := lm.RowList.Get(row)
		if triple.Data == 0 {
			link.(*LIST.SingleList).Remove(triple)
		}else{
			find := link.(*LIST.SingleList).Search(triple)
			if find != nil {
				find.Data = triple
			}else{
				link.(*LIST.SingleList).InsertAtLast(triple)
			}
		}
	}else{
		panic("row or col out of range")
	}
}

func (lm *LinkedMatrix) Get(i,j int) int {
	if i < lm.rows && j < lm.cols && i >= 0 && j >= 0 {
		link := lm.RowList.Get(i)
		if link == nil {
			return 0
		}
		node := link.(*LIST.SingleList).Search(*NewTriple(i,j,0))
		if node == nil || node.Data == nil {
			return 0
		}
		return node.Data.(Triple).Data
	}
	return 0
}

func (lm *LinkedMatrix) GetRows() int {
	return lm.rows
}

func (lm *LinkedMatrix) GetCols() int {
	return lm.cols
}

func (lm *LinkedMatrix) SetRowCol(m,n int) {
	if m > 0 && n > 0 {
		if m > lm.rows {
			for i := lm.rows; i < m; i++  {
				lm.RowList.Insert(i,LIST.NewEmptySingleList())
			}
		}
		lm.rows = m
		lm.cols = n
	}else{
		panic("row and col neither can be below 0")
	}
}

func (lm *LinkedMatrix) String() string {
	str := ""
	str += "LinkedMatrix:\n"
	for i := 0;i < lm.RowList.Size() ;i++  {
		str += lm.RowList.Get(i).(*LIST.SingleList).String()+"\n"
	}
	return str
}