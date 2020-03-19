package main

import (
	"fmt"
	bst "go_dataStruct/ALG/search/binary_search_tree"
)

func main() {
	bstTree, err := bst.New([]int{67, 87, 43, 23, 45, 37, 90})
	if err != nil {
		fmt.Println(err)
	}

	bstTree.Traverse()
	bstTree.Delete(43)
	bstTree.Insert(44)
	bstTree.Traverse()
}