package cirDoubleList

import (
	"fmt"
	"reflect"
)

type cirDoubleList struct {
	head *doubleNode
}

func NewEmptyCirDoubleList() *cirDoubleList {
	cls := &cirDoubleList{
		new(doubleNode),
	}
	cls.head.next = cls.head
	cls.head.prev = cls.head
	return cls
}

func NewCirDoubleList(values []interface{}) *cirDoubleList {
	cls := NewEmptyCirDoubleList()
	p := cls.head
	for i := 0; i < len(values); i++ {
		p.next = &doubleNode{values[i],p,nil}
		p = p.next
	}
	p.next = cls.head
	cls.head.prev = p
	return cls
}

func (cls *cirDoubleList) String() (str string) {
	str = fmt.Sprintf("List (%v) :{",reflect.TypeOf(cls).Elem().Name())
	rear := cls.head
	for rear.next != cls.head {
		rear = rear.next
		if rear.next == cls.head  {
			str += rear.String()
		}else{
			str += rear.String()
			str += ","
		}
	}
	str += fmt.Sprint("}")
	return
}

func (cls *cirDoubleList) IsEmpty() bool {
	return cls.head.next == cls.head
}

func (cls *cirDoubleList) Insert(i int, data interface{}) *doubleNode {
	if data == nil {
		panic("params data can not be nil")
	}
	p := cls.head
	for j := 0;j <= i&&p.next != cls.head ;j++ {
		p = p.next
	}
	ins := &doubleNode{
		data:data,
		next:p,
		prev:p.prev,
	}
	p.prev.next = ins
	p.prev = ins
	return ins
}

func (cls *cirDoubleList) Remove(i int) *doubleNode {
	p := cls.head.next
	for j := 0;j < i;j++ {
		if p == nil {
			return nil
		}
		p = p.next
	}
	defer func() {
		p.prev.next = p.next
		p.next.prev = p.prev
	}()
	return p
}