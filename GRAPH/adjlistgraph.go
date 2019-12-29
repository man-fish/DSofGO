package GRAPH

import (
	"fmt"
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
		ptr := link.(*LIST.SingleList)
		rp := *ptr
		for p := rp.Head.Next; p != nil; p = p.Next {
			alg.RemoveEdge(p.Data.(MATRIX.Triple).ToSymmetry().Col,p.Data.(MATRIX.Triple).ToSymmetry().Row)
		}
		n--
		alg.matrix.RowList.Del(i)
		alg.matrix.SetRowCol(n,n)
		for j := 0;j < n;j++ {
			list := alg.matrix.RowList.Get(i)
			ptr2,ok := list.(*LIST.SingleList)
			if ok {
				rp2 := *ptr2
				for p := rp2.Head.Next; p != nil; p = p.Next {
					tp := p.Data.(MATRIX.Triple)
					tpptr := &tp
					if tp.Row > i {
						tpptr.SubRow()
						/* 血的教训，一定要传递指针 */
					}
					if tp.Col > i {
						tpptr.SubCol()
					}
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

func (alg *AdjListGraph) String() string {
	str := ""
	str += alg.vertexs.String()+"\nadjMatrix:\n"
	for y := 0; y < alg.matrix.RowList.Size(); y++ {
		str += fmt.Sprint(alg.matrix.RowList.Get(y))+"\n"
	}
	return str
}

