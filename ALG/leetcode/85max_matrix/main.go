package main

import (
	"fmt"
)

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n <= 0 {
		return 0
	}
	m := len(matrix[0])
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, m)
	}
	maxS := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 1 {
				dp[i][j] = 0
				continue
			}
			if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + 1
			}
			minWidth := dp[i][j]
			for k := i; k >= 0; k-- {
				minWidth = min(dp[k][j], minWidth)
				maxS = max(maxS, minWidth*(i-k+1))
			}
		}
	}
	return maxS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
	matrix := [][]byte{
		[]byte{1, 0, 1, 0, 0},
		[]byte{1, 0, 1, 1, 1},
		[]byte{1, 1, 1, 1, 1},
		[]byte{1, 0, 0, 1, 0}}
	res := maximalRectangle(matrix)
	fmt.Println(res)
}
