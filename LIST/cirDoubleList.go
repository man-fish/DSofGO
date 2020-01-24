package LIST

import (
	"fmt"
	"reflect"
)

type cirDoubleList struct {
	head *DoubleNode
}

func NewEmptyCirDoubleList() *cirDoubleList {
	cls := &cirDoubleList{
		new(DoubleNode),
	}
	cls.head.Next = cls.head
	cls.head.Prev = cls.head
	return cls
}

func NewCirDoubleList(values []interface{}) *cirDoubleList {
	cls := NewEmptyCirDoubleList()
	p := cls.head
	for i := 0; i < len(values); i++ {
		p.Next = &DoubleNode{values[i],p,nil}
		p = p.Next
	}
	p.Next = cls.head
	cls.head.Prev = p
	return cls
}

func (cls *cirDoubleList) String() (str string) {
	str = fmt.Sprintf("List (%v) :{",reflect.TypeOf(cls).Elem().Name())
	rear := cls.head
	for rear.Next != cls.head {
		rear = rear.Next
		if rear.Next == cls.head  {
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
	return cls.head.Next == cls.head
}

func (cls *cirDoubleList) Insert(i int, data interface{}) *DoubleNode {
	if data == nil {
		panic("params data can not be nil")
	}
	p := cls.head
	for j := 0;j <= i&&p.Next != cls.head ;j++ {
		p = p.Next
	}
	ins := &DoubleNode{
		Data:data,
		Next:p,
		Prev:p.Prev,
	}
	p.Prev.Next = ins
	p.Prev = ins
	return ins
}

func (cls *cirDoubleList) Remove(i int) *DoubleNode {
	p := cls.head.Next
	for j := 0;j < i;j++ {
		if p == nil {
			return nil
		}
		p = p.Next
	}
	defer func() {
		p.Prev.Next = p.Next
		p.Next.Prev = p.Prev
	}()
	return p
}