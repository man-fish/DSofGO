package GRAPH

import (
	"fmt"
	"go_DataStruct/QUEUE"
)

type Graph interface {
	Next(i, j int) int
	GetVertex(i int) interface{}
	VertexCount() int
}

func DFSTraverse(graph Graph, i int) {
	n := graph.VertexCount()
	visited := make([]bool, n)
	j := (i + 1) % n
	for j != i {
		if !visited[j] {
			fmt.Print(j, ":{ ")
			depthfs(graph, j, visited)
			fmt.Print("}")
		}
		j = (j + 1) % n
	}
	if !visited[j] {
		fmt.Print(j, ":{ ")
		depthfs(graph, j, visited)
		fmt.Print("}")
	}
	fmt.Println()
}

func depthfs(graph Graph, i int, visited []bool) {
	fmt.Print(graph.GetVertex(i), " ")
	visited[i] = true
	j := graph.Next(i, -1)
	for j != -1 {
		if !visited[j] {
			depthfs(graph, j, visited)
		}
		j = graph.Next(i, j)
	}
}

func BFSTraverse(graph Graph, i int) {
	n := graph.VertexCount()
	visited := make([]bool, n)
	j := (i + 1) % n
	for j != i {
		if !visited[j] {
			fmt.Print(j, ":{ ")
			breadfs(graph, j, visited)
			fmt.Print("}")
		}
		j = (j + 1) % n
	}
	if !visited[j] {
		fmt.Print(j, ":{ ")
		breadfs(graph, j, visited)
		fmt.Print("}")
	}
	fmt.Println()
}

func breadfs(graph Graph, i int, visited []bool) {
	fmt.Print(graph.GetVertex(i), " ")
	visited[i] = true
	queue := QUEUE.NewLinkedQueue()
	queue.Add(i)
	for !queue.IsEmpty() {
		i = queue.Poll().(int)
		for j := graph.Next(i, -1); j != -1; j = graph.Next(i, j) {
			if !visited[j] {
				fmt.Print(graph.GetVertex(j), " ")
				visited[j] = true
				queue.Add(j)
			}
		}
	}
}
