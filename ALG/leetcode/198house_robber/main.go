package main

import (
	"fmt"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, 2)
	}
	// bestChoice := 0
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[0][0] = 0
			dp[0][1] = nums[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]

	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 1}
	res := rob(nums)
	fmt.Println(res)
}
