package GRAPH

import (
	"go_DataStruct/LIST"
	"go_DataStruct/MATRIX"
	"strconv"
)

type MatrixGraph struct {
	vertexs 	*LIST.SeqList	//顶点
	matrix 		*MATRIX.Matrix	//边
	maxWeight	int				//最大权值
}

func NewEmptyMatrixGraph(length int) *MatrixGraph {
	mxgp := &MatrixGraph{vertexs:LIST.NewEmptySeqList(length),matrix:MATRIX.NewEmprtyMatrix(length,length),maxWeight:0x0000ffff}
	//为邻接矩阵和图申请储存空间，设置最大权值
	mxgp.matrix.InitMatrixForGraph()
	//初始化矩阵
	return mxgp
}
/** 初始化空图 */
func NewMatrixGraph(vertexs []interface{}, tris []MATRIX.Triple) *MatrixGraph {
	mxgp := NewEmptyMatrixGraph(len(vertexs))
	//创建空图
	for _,vertex := range vertexs {
		mxgp.vertexs.Insert(mxgp.vertexs.Size(),vertex)
	}
	//插入顶点
	for _,triple := range tris {
		mxgp.InsertEdge(triple)
	}
	//插入边
	return mxgp
}
/** 初始化图 */
func (g *MatrixGraph) String() string {
	str := ""
	str += g.vertexs.String()+"\nmatrix:\n"
	for y := 0; y < g.matrix.GetRow(); y++ {
		for x := 0; x < g.matrix.GetCol(); x++ {
			if g.matrix.Get(y,x) == 0x0000ffff {
				str += "∞\t"
			}else{
				str += strconv.Itoa(g.matrix.Get(y,x)) +"\t"
			}
		}
		str += "\n"
	}
	return str
}
/** string */
func (g *MatrixGraph) InsertEdge(triple MATRIX.Triple) {
	if triple.Row != triple.Col {
		if triple.Data < 0 || triple.Data > g.maxWeight {
			triple.Data = g.maxWeight
		}
		g.matrix.Set(triple)
	}else{
		panic("can not making up a self circle")
	}
}
/** 插入边，其实就是把边的值设置为权值 */
func (g *MatrixGraph) InsertVertex(vertex interface{}) int {
	i := g.vertexs.Insert(g.vertexs.Size(),vertex)
	//插入顶点的位置
	if i >= g.matrix.GetRow() {
		g.matrix.SetRowCol(i + 1,i + 1)
		//i+1就是现在的长度加一
	}
	for j := 0; j < i; j++ {
		g.matrix.Set(MATRIX.Triple{i,j,g.maxWeight,})
		g.matrix.Set(MATRIX.Triple{j,i,g.maxWeight,})
	}
	/** 注意这里我们的j故意比当前行数小一，保证不会将最后以为设成自身环。*/
	return i
}
/** 插入顶点 */
func (g *MatrixGraph) RemoveEdge(i,j int) {
	if i != j {
		g.matrix.Set(*MATRIX.NewTriple(i,j,g.maxWeight))
	}
}
/** 移除边，其实就是把边的权值设置为最大全权值 */
func (g *MatrixGraph) RemoveVertex(i int) {
	n := g.vertexs.Size()
	if i < n && i >= 0 {
		/** 线性表删除*/
		g.vertexs.Del(i)
		/** 矩阵行列前移 */
		for j := 0; j < n; j++ {
			for k := i+1; k < n;k++ {
				g.matrix.Set(*MATRIX.NewTriple(j ,k-1,g.matrix.Get(j,k)))
			}
		}
		for j := 0; j < n; j++ {
			for k := i+1; k < n;k++ {
				g.matrix.Set(*MATRIX.NewTriple(k-1 ,j,g.matrix.Get(k,j)))
			}
		}
		/** 矩阵缩水 */
		g.matrix.SetRowCol(n-1,n-1)
	}
}
/** 移除顶点 */
func (g *MatrixGraph) VertexCount() int {
	return g.vertexs.Size()
}
/**
 * @param i	行数
 * @param j	列数
 * @return	比如我们有（0，1，2）（0，4，6），传入（0，1）返回 4，传入（0，-1）返回 1。
 */
func (g *MatrixGraph) Mext(i,j int) int {
	n := g.VertexCount()
	if n > i && n > j && i >= 0 && j >= -1 && i != j {
		for k:=j+1; k<n ;k++ {
			if g.matrix.Get(i, k) > 0 && g.matrix.Get(i,k) < g.maxWeight {
				return k
			}
		}
	}
	return -1
}