package LIST

import (
	"errors"
	"fmt"
	"reflect"
)

type SingleList struct{
	Head *Node
}

func NewEmptySingleList() *SingleList {
	return &SingleList{
		Head:new(Node),
	}
}

func NewSingleList(values []interface{}) *SingleList {
	sls := NewEmptySingleList()
	rear := sls.Head
	for i := 0;i < len(values) ;i++  {
		rear.Next = &Node{Data: values[i],Next:nil}
		rear = rear.Next
	}
	return sls
}

func (sls *SingleList) IsEmpty() bool {
	return sls.Head.Next == nil
}

func (sls *SingleList) String() (str string) {
	str = fmt.Sprintf(" (%v) :{",reflect.TypeOf(sls).Elem().Name())
	rear := sls.Head.Next
	for rear != nil {
		if rear.Next == nil  {
			str += rear.String()
		}else{
			str += rear.String()
			str += ","
		}
		rear = rear.Next
	}
	str += fmt.Sprint("}")
	return
}

func (sls *SingleList) Get(i int) interface{} {
	p := sls.Head.Next
	for j := 0;p != nil&&j < i ;j++  {
		p = p.Next
	}
	if i >= 0 && p != nil {
		return p.Data
	}
	return nil
}

func (sls *SingleList) Set(i int, data interface{}) {
	p := sls.Head.Next
	for j := 0;p != nil&&j < i ;j++  {
		p = p.Next
	}
	if p != nil {
		p.Data = data
	}
}


//插入和删除操作不同于设置和获取每次操作的都是要插入位置的前一个元素。
func (sls *SingleList) Insert(i int, data interface{}) *Node {
	if data == nil {
		panic(errors.New("params data can not be nil"))
	}
	p := sls.Head
	for j := 0; p != nil&&j < i; j++ {
		p = p.Next
	}
	p.Next = &Node{data,p.Next}
	return p.Next
}

func (sls *SingleList) Delete(i int) *Node {
	p := sls.Head
	for j := 0; p != nil && j < i; j++ {
		p = p.Next
	}
	defer func() {
		p.Next = p.Next.Next
	}()
	return p.Next
}

func (sls *SingleList) Search(x interface{}) *Node {
	p := sls.Head.Next
	for p != nil {
		if p.Data == x {
			return p
		}
		p = p.Next
	}
	return nil
}

//https://blog.csdn.net/qq_37628075/article/details/82285793
/**
	单链表反转：
		首先我们要明确有两个以上元素的单链表才有反转的意义。
		头节点在这里并不参与反转，实际上的反转起始点时头节点之后的两个节点。
 */

func (sls *SingleList) Reverse() {
	if sls.Head.Next == nil {
		return
	}
	current := sls.Head.Next.Next
	//拿到有头节点链表的第二个元素
	Head := sls.Head.Next
	//拿到有头节点链表的第一个元素，这个元素就是以后的尾元素
	Head.Next = nil
	//清除其指向
	for current != nil {
		currentNext := current.Next
		//储存当前节点的下一个节点
		current.Next = Head
		//更改当前节点指向为前一个节点
		Head = current
		//将Head指针指向当前节点
		current = currentNext
		//将当前指针指向下一个节点
	}
	sls.Head.Next = Head
}

/**
单链表反转递归算法：
	首先我们要明确有两个以上元素的单链表才有反转的意义。
	头节点在这里并不参与反转，实际上的反转起始点时头节点之后的两个节点。
*/

func (sls *SingleList) CurReverse(node *Node) *Node {
	if node.Next == nil {
		sls.Head.Next = node
		return node
	}
	c := sls.CurReverse(node.Next)
	fmt.Println(c.Data)
	if node == sls.Head {
		c.Next = nil
	}else{
		c.Next = node
	}
	return node
}