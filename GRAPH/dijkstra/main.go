package main

import "fmt"

const maxWeight int = 0x0000ffff

type (
	visitedVertex struct {
		alreadyVisited []int // 该顶点是否已经被访问过
		preVisited     []int // 该顶点现最短路径的上一个顶点。
		dis            []int // 该顶点到出发点的现最短距离
	}

	graph struct {
		vertexs     []interface{}
		triples     [][]int
		vertexCount int

		vv *visitedVertex // 储存状态量
	}
)

func newVisitedVertex(vertexCount int, topIndex int) *visitedVertex {
	vv := &visitedVertex{
		alreadyVisited: make([]int, vertexCount),
		preVisited:     make([]int, vertexCount),
		dis:            make([]int, vertexCount),
	}

	for i := range vv.dis {
		if i == topIndex {
			vv.dis[i] = 0
		} else {
			vv.dis[i] = maxWeight
		}
	}
	vv.alreadyVisited[topIndex] = 1 // 标记顶点已访问

	return vv
}

func (vv *visitedVertex) in(idx int) bool {
	return vv.alreadyVisited[idx] == 1
}

func (vv *visitedVertex) updateDis(idx int, dis int) {
	vv.dis[idx] = dis
}

func (vv *visitedVertex) updatePro(pIdx int, idx int) {
	vv.preVisited[pIdx] = idx
}

func (vv *visitedVertex) getDis(idx int) int {
	return vv.dis[idx]
}

func (vv *visitedVertex) update() int {
	idx := 0
	min := maxWeight
	for i := 0; i < len(vv.dis); i++ {
		if !vv.in(i) && vv.getDis(i) < min { // 推选出下一个最小距离顶点
			min = vv.getDis(i)
			idx = i
		}
	}
	vv.alreadyVisited[idx] = 1
	return idx
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

// 实际上djs算法其实是广度优先遍历，遍历从出发点可以到达所有顶点的所有路径。
func (g *graph) djs(idx int) {
	g.vv = newVisitedVertex(g.vertexCount, idx)
	g.updateVV(idx)
	for i := 1; i < len(g.vertexs); i++ {
		idx = g.vv.update()
		g.updateVV(idx)
	}

	{
		fmt.Println(g.vv.dis)
		fmt.Println(g.vv.alreadyVisited)
		fmt.Println(g.vv.preVisited)
	}
}

func (g *graph) updateVV(index int) { // 主要代码
	min := 0
	// 当前顶点拓扑下，所有相关未访问节点到顶点的距离。
	for i := 0; i < len(g.triples[index]); i++ {
		min = g.vv.getDis(index) + g.triples[index][i]

		// 顶点 j 到出发点的新最短距离为顶点到index的距离 + index 和 j 的边。
		if min < g.vv.getDis(i) && !g.vv.in(i) {
			g.vv.updateDis(i, min)
			g.vv.updatePro(i, index)
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
	g.djs(0)
}
