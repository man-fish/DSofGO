package MATRIX

import (
	"errors"
	"fmt"
	"reflect"
)

type Matrix struct{
	elem	[][]int	/** 储存数组 */
	m,n		int		/** m行 n列 */
}

/** 初始化矩阵空间 */
func NewEmprtyMatrix(m, n int) *Matrix {
	elem := make([][]int,m)
	for idx := range elem {
		elem[idx] = make([]int,n)
	}
	//创建一个m行n列全为0的矩阵。
	return &Matrix{elem, m, n}
}
/** 初始化矩阵 */
func NewMatrix(m, n int, triples []Triple) *Matrix {
	elem := make([][]int,m)
	for idx := range elem {
		elem[idx] = make([]int,n)
	}
	//创建一个m行n列全为0的矩阵。
	mx := &Matrix{elem, m, n}
	for _, triple := range triples {
		mx.Set(triple)
	}
	//将值填入矩阵，
	return mx
}
/** 初始化临界矩阵图 */
func (mx *Matrix) InitMatrixForGraph (){
	//对于邻接矩阵图，我们需要没有边的点都为无穷，对角线上的点都是0。
	for y := 0; y < mx.m; y++ {
		for x := 0; x < mx.n; x++ {
			if x != y {
				mx.elem[y][x]= 0x0000ffff
			}
		}
	}
}
/** 设置矩阵点的值 */
func (mx *Matrix) Set(triple Triple) {
	y := triple.Row
	x := triple.Col
	if y >= 0 && x >= 0 && y < mx.m && x < mx.n{
		mx.elem[y][x] = triple.Data
	}else{
		panic(errors.New("下标越界"))
	}
}
/** 获取矩阵点的值 */
func (mx *Matrix) Get(i,j int) int {
	y := i
	x := j
	if y >= 0 && x >= 0 && y < mx.m && x < mx.n{
		return mx.elem[y][x]
	}else{
		panic(errors.New("下标越界"))
	}
}
/** String */
func (mx *Matrix) String() string {
	str := ""
	str += fmt.Sprint(reflect.TypeOf(mx).Elem().Name(),":\n")
	for _,row := range mx.elem {
		str += fmt.Sprint(row,"\n")
	}
	return str
}
/** 获取矩阵行数 */
func (mx *Matrix) GetRow() int {
	return mx.n
}
/** 获取矩阵列数 */
func (mx *Matrix) GetCol() int {
	return mx.n
}
/** 重新设置矩阵的行数和列数 */
func (mx *Matrix) SetRowCol(m,n int) {
	elem := make([][]int,m)
	for idx := range elem {
		elem[idx] = make([]int,n)
	}
	//创建一个新的矩阵
	/** 构造一个新的数组 */
	if mx.m > m || mx.n > n {
		for y := 0; y < m; y ++ {
			for x := 0; x < n; x ++ {
				elem[y][x] = mx.elem[y][x]
			}
		}
	}else{
		for y,row := range mx.elem {
			for x,item := range row {
				elem[y][x] = item
			}
		}
	}

	mx.elem = elem
	mx.m = m
	mx.n = n
}
