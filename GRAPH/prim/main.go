package main

import "fmt"

const maxWeight int = 0x0000ffff

type graph struct {
	vertexs     []interface{}
	triples     [][]int
	vertexCount int
}

func new(v int, triples [][]int) *graph {
	start := 'a'
	vertexs := make([]interface{}, 0)
	for i := 0; i < v; i++ {
		cur := int(start) + i
		vertexs = append(vertexs, rune(cur))
	}
	return &graph{
		vertexs:     vertexs,
		triples:     triples,
		vertexCount: v,
	}
}

func (g *graph) printGraph() {
	for i := 0; i < g.vertexCount; i++ {
		for j := 0; j < g.vertexCount; j++ {
			if g.triples[i][j] == maxWeight {
				fmt.Print("∆\t")
				continue
			}
			fmt.Print(g.triples[i][j], "\t")
		}
		fmt.Println()
	}
}

func (graph *graph) prim(top int) {
	visited := make([]bool, graph.vertexCount)

	h1 := -1
	h2 := -1
	minWeight := maxWeight

	visited[top] = true

	for i := 0; i < graph.vertexCount-1; i++ {

		for j := 0; j < graph.vertexCount; j++ {
			for k := 0; k < graph.vertexCount; k++ {
				if visited[j] && !visited[k] && graph.triples[j][k] < minWeight {
					h1 = j
					h2 = k
					minWeight = graph.triples[j][k]
				}
			}
		}
		visited[h2] = true
		fmt.Printf("edge：<%s,%s> with weight %d \n", string(graph.vertexs[h1].(rune)), string(graph.vertexs[h2].(rune)), minWeight)
		minWeight = maxWeight
	}
}

func main() {
	triples := make([][]int, 7)
	for i := range triples {
		triples[i] = make([]int, 7)
		for j := range triples[i] {
			triples[i][j] = maxWeight
		}
	}
	{
		triples[0][1] = 5
		triples[0][2] = 7
		triples[0][6] = 2

		triples[1][0] = 5
		triples[1][3] = 9
		triples[1][6] = 3

		triples[2][0] = 7
		triples[2][4] = 8

		triples[3][1] = 9
		triples[3][5] = 4

		triples[4][2] = 8
		triples[4][5] = 5
		triples[4][6] = 4

		triples[5][3] = 4
		triples[5][4] = 5
		triples[5][6] = 6

		triples[6][0] = 2
		triples[6][1] = 3
		triples[6][4] = 4
		triples[6][5] = 6
	}
	g := new(7, triples)
	g.printGraph()
	g.prim(0)

	var a interface{} = 'a'
	fmt.Println(a)
	fmt.Printf("%c", a.(rune))
}
