package main

import (
	"go_dataStruct/util"
)

func rotate(matrix [][]int) {
	size := len(matrix)
	subsize := size - 1
	for i := 0; i < size/2; i++ {
		for j := i; j < subsize-i; j++ {
			matrix[i][j], matrix[j][subsize-i] = matrix[j][subsize-i], matrix[i][j]
			matrix[i][j], matrix[subsize-i][subsize-j] = matrix[subsize-i][subsize-j], matrix[i][j]
			matrix[i][j], matrix[subsize-j][i] = matrix[subsize-j][i], matrix[i][j]
		}
	}
}

func rotate2(matrix [][]int) {
	idx, n := 0, len(matrix)
	for n > 0 && idx <= n/2 {
		fsEle, lsEle := idx, n-idx-1
		tmp := matrix[fsEle][fsEle]
		for i := fsEle + 1; i <= lsEle; i++ {
			tmp, matrix[fsEle][i] = matrix[fsEle][i], tmp
		}

		for i := fsEle + 1; i <= lsEle; i++ {
			tmp, matrix[i][lsEle] = matrix[i][lsEle], tmp
		}

		for i := lsEle - 1; i >= fsEle; i-- {
			tmp, matrix[lsEle][i] = matrix[lsEle][i], tmp
		}

		for i := lsEle - 1; i >= fsEle; i-- {
			tmp, matrix[i][fsEle] = matrix[i][fsEle], tmp
		}
		idx++
	}
}

func main() {
	matrix := [][]int{
		// []int{1, 2},
		// []int{4, 5},
	}
	rotate(matrix)
	util.Print2DArray(matrix)
}
