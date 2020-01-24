package main

import (
	"fmt"
	"go_DataStruct/MATRIX"
)

func editMx(mx *[][]int) {
	*mx = append(*mx, []int{3,1})
}

func main() {
	a := [][]int{{1,2},{2,3}}
	fmt.Println(a)
	editMx(&a)
	fmt.Println(a)

	tris := []MATRIX.Triple{*MATRIX.NewTriple(1,1,90),*MATRIX.NewTriple(2,1,7),*MATRIX.NewTriple(1,2,78)}
	mx := MATRIX.NewMatrix(5,5,tris)
	fmt.Println(mx)
	mx.SetRowCol(10,10)
	fmt.Println(mx)

	lmx := MATRIX.NewLinkedMatrix(5,5,tris)
	fmt.Println(lmx)
	lmx.Set(*MATRIX.NewTriple(1, 1,0))
	fmt.Println(lmx)
	fmt.Println(lmx.Get(1,2))
}
