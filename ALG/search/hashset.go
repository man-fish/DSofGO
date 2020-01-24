package search

import (
	"fmt"
	"go_DataStruct/LIST"
	"go_DataStruct/util"
	"hash/crc32"
)

const LOAD_FATOR float32 = 0.75

type HashSet struct {
	table 		[]*LIST.SingleList
	Count		int
}

func NewEmptyHashSet(length int) *HashSet {
	if length < 10 {
		length = 10
	}
	hs := &HashSet{make([]*LIST.SingleList,length),0}
	for i := 0;i < length ;i++  {
		hs.table[i] = LIST.NewEmptySingleList()
	}
	return hs
}

func NewHashSet(values []interface{}) *HashSet {
	length := (int)(float32(len(values))/LOAD_FATOR)
	hs := NewEmptyHashSet(length)
	hs.Count = len(values)
	for _,item := range values {
		hs.table[hs.hash(item)].InsertDifferent(item)
	}
	return hs
}

func (hs *HashSet) hash(x interface{}) int {
	hashkey := int(crc32.ChecksumIEEE([]byte(fmt.Sprint(x))))
	key := hashkey % len(hs.table)
	return key
}

func (hs *HashSet) Search(x util.Comparable) interface{} {
	find := hs.table[hs.hash(x)].Search(x)
	if find != nil && find.Data != nil{
		return find.Data
	}else{
		return nil
	}
}

func (hs *HashSet) Remove(x util.Comparable) interface{} {
	find := hs.table[hs.hash(x)].Remove(x)
	if find != nil {
		hs.Count--
	}
	return find
}

func (hs *HashSet) Insert(x interface{}) {

}