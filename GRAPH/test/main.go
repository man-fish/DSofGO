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
	var graph GRAPH.Graph
	graph = GRAPH.NewMatrixGraph([]interface{}{"A","B","C"},tris)
	fmt.Println(graph)
	graph.(*GRAPH.MatrixGraph).InsertVertex("D")
	fmt.Println(graph)
	GRAPH.DFSTraverse(graph,0)
	GRAPH.BFSTraverse(graph,0)
	graph.(*GRAPH.MatrixGraph).RemoveVertex(1)
	fmt.Println(graph)
	gp := GRAPH.NewAdjListGraph([]interface{}{"A","B","C"},tris)
	fmt.Println(gp)
	gp.InsertVertex("D")
	fmt.Println(gp)
	gp.InsertEdge(*MATRIX.NewTriple(3,1,78))
	fmt.Println(gp)
	gp.RemoveEdge(3,1)
	fmt.Println(gp)
	gp.RemoveVertex(3)
	fmt.Println(gp)
}