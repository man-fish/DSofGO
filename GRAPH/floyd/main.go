package main

import "fmt"

const maxWeight int = 0x0000ffff

type graph struct {
	vertexs     []interface{}
	vertexCount int

	dis [][]int // 存储最短距离
	pre [][]int // 存储所有路径
}

func new(v int, triples [][]int) *graph {
	start := 'a'
	vertexs := make([]interface{}, 0)
	for i := 0; i < v; i++ {
		cur := int(start) + i
		vertexs = append(vertexs, rune(cur))
	}

	pre := make([][]int, v)
	for i := 0; i < v; i++ {
		pre[i] = make([]int, v)
		for j := 0; j < v; j++ {
			pre[i][j] = i
		}
	}

	return &graph{
		vertexCount: v,
		vertexs:     vertexs,
		dis:         triples,
		pre:         pre,
	}
}

func (g *graph) printGraph() {
	for i := 0; i < g.vertexCount; i++ {
		for j := 0; j < g.vertexCount; j++ {
			if g.dis[i][j] == maxWeight {
				fmt.Print("∆\t")
				continue
			}
			fmt.Print(g.dis[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < g.vertexCount; i++ {
		for j := 0; j < g.vertexCount; j++ {
			s := 'a'
			fmt.Printf("%c\t", rune(int(s)+g.pre[i][j]))
		}
		fmt.Println()
	}

}

func (g *graph) floyd() {
	minDis := 0

	for k := 0; k < g.vertexCount; k++ {

		for i := 0; i < g.vertexCount; i++ {
			for j := 0; j < g.vertexCount; j++ {
				minDis = g.dis[i][k] + g.dis[k][j]
				if i != j && minDis < g.dis[i][j] {
					g.pre[i][j] = k
					g.dis[i][j] = minDis
				}
			}
		}

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
	g.floyd()
	fmt.Println()
	g.printGraph()
}
