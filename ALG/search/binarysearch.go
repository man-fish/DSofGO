package search

import (
	"go_DataStruct/LIST"
	"go_DataStruct/util"
)

type SearchAbleSeqList struct {
	LIST.SeqList
}

func (bsas *SearchAbleSeqList) Search(x util.Comparable) int {
	for i := 0; i < bsas.Size() ; i++  {
		if x.IsEqual(bsas.Get(i),x) {
			return i
		}
	}
	return -1
}

func (bsas *SearchAbleSeqList) BinarySearch(x util.Comparable) int {
	start := 0
	end := bsas.Size()-1
	for start <= end {
		middle := start + end
		if x.CompareTo(x,bsas.Get(middle)) == 0 {
			return middle
		}else if x.CompareTo(x,bsas.Get(middle)) == 1 {
			start = middle+1
		}else {
			end = middle - 1
		}
	}
	return -1
}