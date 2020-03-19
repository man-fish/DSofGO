/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
示例 1:
	输入:
	[
	  [1,1,1],
	  [1,0,1],
	  [1,1,1]
	]
	输出:
	[
	  [1,0,1],
	  [0,0,0],
	  [1,0,1]
	]
进阶:
	一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
	一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
	你能想出一个常数空间的解决方案吗？
思路：
	用第一行做记录
*/

package main

import "fmt"

func setZeroes(matrix [][]int) {
	var (
		storeZone = -1
		hasZore   = false
	)

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 && storeZone == -1 {
				storeZone = i
				break
			}
			if storeZone != -1 {
				if matrix[i][j] == 0 {
					matrix[storeZone][j] = 0
					for k := storeZone; k < i; k++ {
						matrix[k][j] = 0
					}
					hasZore = true
					continue
				}

				if matrix[storeZone][j] == 0 {
					matrix[i][j] = 0
					continue
				}
			}
		}
		if hasZore {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
			hasZore = false
		}

	}
	for i := 0; i <= storeZone; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == storeZone {
				matrix[i][j] = 0
				continue
			}
			if matrix[storeZone][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroesPro(matrix [][]int) {
	var (
		hasZore    = false
		firHasZore = false
	)

	for i := 0; i < len(matrix); i++ {
		hasZore = false
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firHasZore = true //因为第一行用来记录，所以说不能判断第一行是否应该清0，所以添加一个标志位。
					break
				} else {
					for k := 0; k < i; k++ { //将之前的值改为0
						matrix[k][j] = 0
					}
					hasZore = true
				}
				continue
			}

			if matrix[0][j] == 0 {
				matrix[i][j] = 0
				continue
			}
		}
		if hasZore {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	if firHasZore {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}

/*
	[
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	]
	[
		[]int{0, 0, 0, 5},
		[]int{4, 3, 1, 4},
		[]int{0, 1, 1, 4},
		[]int{1, 2, 1, 3},
		[]int{0, 0, 1, 1},
	]
*/

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

func main() {
	matrix := [][]int{
		[]int{0, 0, 0, 5},
		[]int{4, 3, 1, 4},
		[]int{0, 1, 1, 4},
		[]int{1, 2, 1, 3},
		[]int{0, 0, 1, 1},
	}
	setZeroesPro(matrix)
	printMatrix(matrix)

}
