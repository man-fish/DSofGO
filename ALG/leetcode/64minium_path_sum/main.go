package main

import (
	"fmt"
	"go_dataStruct/util"
	"math"
)

// 说明：每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {
	r := len(grid[0])
	c := len(grid)

	if r == 0 || c == 0 {
		return 0
	}

	dp := make([][]int, c)
	for idx := range dp {
		dp[idx] = make([]int, r)
	}

	res := dfs(grid, dp, 0, 0)
	util.Print2DArray(dp)
	return res
}

func dfs(grid, dp [][]int, x, y int) int {
	m := len(grid[0]) - 1
	n := len(grid) - 1
	if y == n && x == m {
		return grid[y][x]
	}
	if dp[y][x] != 0 {
		return dp[y][x]
	}
	curBestChoice := math.MaxInt64
	if x+1 <= m {
		curBestChoice = min(curBestChoice, dfs(grid, dp, x+1, y))
	}

	if y+1 <= n {
		curBestChoice = min(curBestChoice, dfs(grid, dp, x, y+1))
	}
	dp[y][x] = curBestChoice + grid[y][x]
	return dp[y][x]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func mldp(grid [][]int) int {
	x, y := len(grid[0]), len(grid)
	if x == 0 || y == 0 {
		return 0
	}

	dp := make([][]int, y)
	for idx := range dp {
		dp[idx] = make([]int, x)
	}

	for i := y - 1; i >= 0; i-- {
		for j := x - 1; j >= 0; j-- {
			if i == y-1 {
				if j == x-1 {
					dp[i][j] = grid[i][j]
					continue
				}
				dp[i][j] = dp[i][j+1] + grid[i][j]
			} else if j == x-1 {
				dp[i][j] = dp[i+1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
			}
		}
	}

	return dp[0][0]
}

func mldp2(grid [][]int) int {
	x, y := len(grid[0]), len(grid)
	if x == 0 || y == 0 {
		return 0
	}

	// dp := make([][]int, y)
	// for idx := range dp {
	// 	dp[idx] = make([]int, x)
	// }

	for i := y - 1; i >= 0; i-- {
		for j := x - 1; j >= 0; j-- {
			if i == y-1 {
				if j == x-1 {
					continue
				}
				grid[i][j] = grid[i][j+1] + grid[i][j]
			} else if j == x-1 {
				grid[i][j] = grid[i+1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i+1][j], grid[i][j+1]) + grid[i][j]
			}
		}
	}

	return grid[0][0]
}

// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
func main() {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	util.Print2DArray(grid)
	fmt.Println(mldp2(grid))
	fmt.Println(minPathSum(grid))
}
