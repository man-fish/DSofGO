package seqList

import (
	"errors"
	"fmt"
	"reflect"
)

type seqList struct {
	element       []interface{}
	currentLength int
	maxLength     int
}

func CreateSeqList(values []interface{}) *seqList {
	sl := seqList{
		element:       make([]interface{}, len(values)),
		maxLength:     len(values),
		currentLength: len(values),
	}

	for idx, item := range values {
		sl.element[idx] = item
	}

	return &sl
}

func (sl *seqList) IsEmpty() bool {
	return sl.currentLength == 0
}

func (sl *seqList) Size() int {
	return sl.currentLength
}

func (sl *seqList) MaxSize() int {
	return sl.maxLength
}

func (sl *seqList) String() (str string) {
	str = fmt.Sprintf("List (%v) :{",reflect.TypeOf(sl).Elem().Name())
	for i:=0;i<sl.currentLength;i++ {
		if i == sl.currentLength-1 {
			str += fmt.Sprintf("%v",sl.element[i])
		}else{
			str += fmt.Sprintf("%v,",sl.element[i])
		}
	}
	str += fmt.Sprint("}")
	return
}


func (sl *seqList) Get(i int) interface{} {
	if i > sl.currentLength || i < 0 {
		return nil
	}
	return sl.element[i]
}

func (sl *seqList) Set(i int, x interface{}) {
	if x == nil {
		panic("params x can not be nil")
	}

	if i > sl.currentLength || i < 0 {
		panic(errors.New("params i out of bounds"))
	}

	sl.element[i] = x
}

func (sl *seqList) Insert(i int, x interface{}) int {
	if x == nil {
		panic("params x can not be nil")
	}

	if i < 0 {
		i = 0
	}

	if i > sl.currentLength {
		i = sl.currentLength
	}

	temp := sl.element

	if sl.maxLength == sl.currentLength {
		sl.element = make([]interface{}, sl.maxLength*2)
		sl.maxLength = sl.maxLength*2
	}

	for j := 0; j < i; j++ {
		sl.element[j] = temp[j]
	}

	sl.element[i] = x

	for j := sl.currentLength-1; j >= i; j-- {
		sl.element[j+1] = temp[j]
	}

	sl.currentLength ++
	return i
}

func (sl *seqList) InsertDef(x interface{}) {
	sl.Insert(sl.currentLength,x)
}

func (sl *seqList) Del(i int) interface{} {
	if i > sl.currentLength || i < 0 {
		return nil
	}

	temp := sl.element[i]

	for j := i;j < sl.currentLength-1 ;j++ {
		sl.element[j] = sl.element[j+1]
	}

	sl.element[sl.currentLength-1] = nil
	sl.currentLength--
	return temp
}