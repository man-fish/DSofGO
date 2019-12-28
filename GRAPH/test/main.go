package main

import (
	"fmt"
	"go_DataStruct/GRAPH"
	"go_DataStruct/MATRIX"
)

func xy(arr [][]int) {
	arr[0][0] = 999
}

type a struct {
	name string
}

func main() {
	tris := []MATRIX.Triple{*MATRIX.NewTriple(1,2,2)}
	graph := GRAPH.NewMatrixGraph([]interface{}{"A","B","C"},tris)
	fmt.Println(graph)
	graph.InsertVertex("D")
	fmt.Println(graph)
	graph.RemoveVertex(1)
	fmt.Println(graph)

}