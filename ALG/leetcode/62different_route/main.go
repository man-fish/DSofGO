package main

import (
	"fmt"
	"go_dataStruct/util"
)

func uniquePaths(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, m)
	}
	res := dfs(m-1, n-1, dp)
	util.Print2DArray(dp)
	return res
}

func dfs(m, n int, dp [][]int) int {
	if m == 0 && n == 0 {
		return 1
	}
	count := 0
	if dp[n][m] != 0 {
		return dp[n][m]
	}
	if n-1 >= 0 {
		count += dfs(m, n-1, dp)
	}
	if m-1 >= 0 {
		count += dfs(m-1, n, dp)
	}
	dp[n][m] = count
	return count
}

func dp(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 {
				dp[0][j] = 1
			} else if j == 0 {
				dp[i][0] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	util.Print2DArray(dp)
	return dp[n-1][m-1]
}
func main() {
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(dp(7, 3))
}
