package main

import (
	"fmt"
	"go_DataStruct/TREE"
)


func main() {
	haff := TREE.NewHaffTree([]int{7,5,1,2})
	fmt.Println(haff.GetCode(0))
	fmt.Println(haff)
	fmt.Println(haff.Decode(haff.Encode("ACA")))
	fmt.Println('我')
}
