package main

import (
	"fmt"
	"go_DataStruct/QUEUE"
	"reflect"
)

type Queue interface {
	Poll() interface{}
	Peek() interface{}
	Add(interface{}) bool
}

func testQueue(q Queue) {
	switch v := q.(type) {
	case *QUEUE.SeqQueue:
		fmt.Printf("===============测试%v队列===============\n",reflect.TypeOf(v).Elem().Name())
	default:
		fmt.Printf("===============测试%v队列===============\n",reflect.TypeOf(v).Elem().Name())
	}
	fmt.Println(q.Peek())
	for i := 0; i < 5 ; i++  {
		q.Add(i)
	}
	fmt.Println(q.Peek())
	fmt.Println(q.Peek())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
}

type chick struct {
	name string
}

type dog struct {
	name string
	father chick
}

type stu struct {
	name string
	pet  dog
}

func (s *stu)setName() {
	s.pet.father.name = "kk"
}
func main() {
	stu1 := &stu{name:"wang",pet:dog{"papi",chick{}}}
	//testQueue(QUEUE.NewSequeue(3))
	//testQueue(QUEUE.NewLinkedQueue())
	//testQueue(QUEUE.NewDoubleLinkedQueue())

	//fmt.Println(reflect.TypeOf(stu1).Kind())
	stu1.setName()
	fmt.Println(stu1.pet.father.name)
}
/**
	首先我们可以明白第一层，不管是结构体指针还是结构体都可以打点调用其属性，
	指针接收者方法会自动转换值，值接受者方法不会自动转换指针。
	向函数内传入结构体的时候是值类型传递，在函数内部修改其属性不会对实际传入的结构体产生影响
	但是如果传入的结构体的某一个属性是一个指针，那么修改这个属性就会对结构体造成影响
	反之修改结构体指针上的所有属性都会产生实际效果
	联合数据结构去想，结构体的私有属性是值类型和指针类型其实都可以，因为我们只能通过指针接收者函数去访问他们
	但是对于公共属性来说，最好是指针类型，因为我们可能直接通过结构体（不在指针结构者方法内）打点修改其属性，比如说node的Next
 */


