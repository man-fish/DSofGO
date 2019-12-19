package LIST

import (
	"errors"
	"fmt"
	"reflect"
)

type SeqList struct {
	element       []interface{}
	currentLength int
	maxLength     int
}

func NewSeqList(values []interface{}) *SeqList {
	sl := SeqList{
		element:       make([]interface{}, len(values)),
		maxLength:     len(values),
		currentLength: len(values),
	}

	for idx, item := range values {
		sl.element[idx] = item
	}

	return &sl
}

func (sl *SeqList) IsEmpty() bool {
	return sl.currentLength == 0
}

func (sl *SeqList) Size() int {
	return sl.currentLength
}

func (sl *SeqList) MaxSize() int {
	return sl.maxLength
}

func (sl *SeqList) String() (str string) {
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


func (sl *SeqList) Get(i int) interface{} {
	if i > sl.currentLength || i < 0 {
		return nil
	}
	return sl.element[i]
}

func (sl *SeqList) Set(i int, x interface{}) {
	if x == nil {
		panic("params x can not be nil")
	}

	if i > sl.currentLength || i < 0 {
		panic(errors.New("params i out of bounds"))
	}

	sl.element[i] = x
}

func (sl *SeqList) Insert(i int, x interface{}) int {
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

func (sl *SeqList) InsertDef(x interface{}) {
	sl.Insert(sl.currentLength,x)
}

func (sl *SeqList) Del(i int) interface{} {
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