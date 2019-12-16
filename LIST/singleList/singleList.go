package singleList

import (
	"errors"
	"fmt"
	"reflect"
)

type singleList struct{
	head *node
}

func NewEmptySingleList() *singleList {
	return &singleList{
		head:new(node),
	}
}

func NewSingleList(values []interface{}) *singleList {
	sls := NewEmptySingleList()
	rear := sls.head
	for i := 0;i < len(values) ;i++  {
		rear.next = &node{data:values[i],next:nil}
		rear = rear.next
	}
	return sls
}


func (sls *singleList) String() (str string) {
	str = fmt.Sprintf("List (%v) :{",reflect.TypeOf(sls).Elem().Name())
	rear := sls.head.next
	for rear != nil {
		if rear.next == nil  {
			str += rear.String()
		}else{
			str += rear.String()
			str += ","
		}
		rear = rear.next
	}
	str += fmt.Sprint("}")
	return
}

func (sls *singleList) Get(i int) interface{} {
	p := sls.head.next
	for j := 0;p != nil&&j <= i ;j++  {
		p = p.next
	}
	if i >= 0 && p != nil {
		return p.data
	}
	return nil
}

func (sls *singleList) Set(i int, data interface{}) {
	p := sls.head.next
	for j := 0;p != nil&&j < i ;j++  {
		p = p.next
	}
	if p != nil {
		p.data = data
	}
}


//插入和删除操作不同于设置和获取每次操作的都是要插入位置的前一个元素。
func (sls *singleList) Insert(i int, data interface{}) *node {
	if data == nil {
		panic(errors.New("params data can not be nil"))
	}
	p := sls.head
	for j := 0; p != nil&&j < i; j++ {
		p = p.next
	}
	p.next = &node{data,p.next}
	return p.next
}

func (sls *singleList) Delete(i int) *node {
	p := sls.head
	for j := 0; p != nil && j < i; j++ {
		p = p.next
	}
	defer func() {
		p.next = p.next.next
	}()
	return p.next
}

func (sls *singleList) Search(x interface{}) *node {
	p := sls.head.next
	for p != nil {
		if p.data == x {
			return p
		}
		p = p.next
	}
	return nil
}

//https://www.javazhiyin.com/32787.html
func (sls *singleList) Reverse() *node {
	if sls.head.next == nil {
		return sls.head
	}
	head := sls.head

	preNode := new(node)
	nextNode := new(node)

	for head != nil {
		nextNode = head.next
		head.next = preNode
		preNode = head
		head = nextNode
	}

	return preNode
}