package main

import (
	"fmt"
	"go_DataStruct/LIST/cirDoubleList"
	"go_DataStruct/LIST/seqList"
	"go_DataStruct/LIST/singleList"
)

func testSeqList(){
	fmt.Println("=================测试线性表=================")
	testArr := []interface{}{"1",2,"coconut",nil}
	list := seqList.CreateSeqList(testArr)
	fmt.Println(list.String())
	fmt.Println(list.Size())
	fmt.Println(list.Del(3))
	fmt.Println(list.Size())
	fmt.Println(list.String())

	fmt.Println(list.Insert(1,"wang"))
	fmt.Println(list.MaxSize())
	fmt.Println(list.Insert(1,"wang"))
	fmt.Println(list.MaxSize())
	fmt.Println(list.Insert(1,"wang"))
	fmt.Println(list.MaxSize())
	fmt.Println(list.Insert(1,"wang"))
	fmt.Println(list.MaxSize())
	fmt.Println(list.Insert(1,"wang"))
	fmt.Println(list.MaxSize())
	fmt.Println(list.Insert(1,"wang"))
	fmt.Println(list.MaxSize())
	fmt.Println(list.String())
}

func testSingleList() {
	fmt.Println("=================测试单链表=================")
	testArr := []interface{}{"1",2,"coconut"}
	list := singleList.NewSingleList(testArr)
	fmt.Println(list.String())
	fmt.Println(list.Insert(0,888))
	fmt.Println(list.String())
	fmt.Println(list.Delete(0))
	fmt.Println(list.String())
	fmt.Println(list.Get(0))
	list.Set(1,10)
	fmt.Println(list.String())
	n := list.Search("coconut")
	fmt.Println(n)
	list.Reverse()
	fmt.Println(list.String())
}

func testDoubleSingleList() {
	fmt.Println("=================测试双链表=================")
	testArr := []interface{}{"1",2,"coconut"}
	list := cirDoubleList.NewCirDoubleList(testArr)
	fmt.Println(list.IsEmpty())
	fmt.Println(list.String())
	fmt.Println(list.Insert(1,"goDS").String())
	fmt.Println(list.String())
	fmt.Println(list.Remove(1).String())
	fmt.Println(list.String())
}
func main() {
	testSeqList()
	testSingleList()
	testDoubleSingleList()
}


