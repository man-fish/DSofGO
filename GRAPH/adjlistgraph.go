package GRAPH

import (
	"go_DataStruct/LIST"
	"go_DataStruct/MATRIX"
)

type AdjListGraph struct{
	vertexs 	*LIST.SeqList
	matrix 		*MATRIX.LinkedMatrix
	maxlength	int
}

func NewEmptyAdjListGraph(length int) *AdjListGraph {
	return &AdjListGraph{vertexs:LIST.NewEmptySeqList(length),matrix:MATRIX.NewEmptyLinkedMatrix(length,length),maxlength:0x0000ffff}
}

func NewAdjListGraph(vertexs []interface{},triples []MATRIX.Triple) *AdjListGraph {
	alg := NewEmptyAdjListGraph(len(vertexs))
	for idx, vertex := range vertexs {
		alg.vertexs.Insert(idx,vertex)
	}

	for _, triple := range triples {
		alg.InsertEdge(triple)
	}

	return alg
}

func (alg *AdjListGraph) InsertEdge(triple MATRIX.Triple) {
	if triple.Col != triple.Row {
		if triple.Data < 0 || triple.Data > alg.maxlength {
			triple.Data = 9
		}
		alg.matrix.Set(triple)
	}else{
		panic("can not making up a self circle")
	}
}

func (alg *AdjListGraph) InsertVertex(x interface{}) int {
	i := alg.vertexs.Insert(alg.vertexs.Size(),x)
	if i >= alg.matrix.GetCols() {
		alg.matrix.SetRowCol(i+1,i+1)
	}
	return i
}

func (alg *AdjListGraph) RemoveEdge(i,j int) {
	alg.matrix.Set(*MATRIX.NewTriple(i,j,0))
}

func (alg *AdjListGraph) RemoveVertex(i int) {
	n := alg.VertexCount()
	if i >= 0 && i < n {
		link := alg.matrix.RowList.Get(i)
		for p := link.(LIST.SingleList).Head.Next; p != nil; p = p.Next {
			alg.RemoveEdge(p.Data.(MATRIX.Triple).ToSymmetry().Col,p.Data.(MATRIX.Triple).ToSymmetry().Row)
		}
		n--
		alg.matrix.RowList.Del(i)
		alg.matrix.SetRowCol(n,n)
		for j := 0;j < n;j++ {
			list := alg.matrix.RowList.Get(i)
			for p := list.(LIST.SingleList).Head.Next; p != nil; p = p.Next {
				if p.Data.(MATRIX.Triple).Row > i {
					p.Data.(MATRIX.Triple).Row--
				}
				if p.Data.(MATRIX.Triple).Col > i {
					p.Data.(MATRIX.Triple).Col--
				}
			}
		}
		alg.vertexs.Del(i)
	}else{
		panic("out of range")
	}
}

func (alg *AdjListGraph) VertexCount() int {
	return alg.vertexs.Size()
}