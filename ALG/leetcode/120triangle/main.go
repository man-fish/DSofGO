package main

import (
	"fmt"
	"go_dataStruct/util"
	"math"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle) - 1
	bestChoice := math.MaxInt16
	// dp := make([]int, n)
	for idx := range triangle[n] {
		bestChoice = min(dfs(triangle, n, idx), bestChoice)
	}
	return bestChoice
}

func dfs(triangle [][]int, n, i int) int {
	if n == 0 {
		return triangle[0][0]
	}
	bestChoice := math.MaxInt16
	if i != len(triangle[n])-1 {
		bestChoice = min(bestChoice, dfs(triangle, n-1, i)+triangle[n][i])
	}

	if i != 0 {
		bestChoice = min(bestChoice, dfs(triangle, n-1, i-1)+triangle[n][i])
	}
	return bestChoice
}

func minimumTotalTopDown(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for idx := range triangle {
		dp[idx] = make([]int, len(triangle[idx]))
	}

	res := dfs2(triangle, dp, 0, 0)
	util.Print2DArray(triangle)
	util.Print2DArray(dp)
	return res
}

func dfs2(triangle [][]int, dp [][]int, n, i int) int {
	if n == len(triangle)-1 {
		return triangle[n][i]
	}
	if dp[n][i] != 0 {
		return dp[n][i]
	}
	bestChoice := math.MaxInt64
	bestChoice = min(bestChoice, dfs2(triangle, dp, n+1, i)+triangle[n][i])
	bestChoice = min(bestChoice, dfs2(triangle, dp, n+1, i+1)+triangle[n][i])
	dp[n][i] = bestChoice
	return bestChoice
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotalDy(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	for i := n - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

// [2],
// [3,4],
// [6,5,7],
// [4,1,8,3]

func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	res := minimumTotal(triangle)
	res = minimumTotalTopDown(triangle)
	res = minimumTotalDy(triangle)
	fmt.Println(res)
}
