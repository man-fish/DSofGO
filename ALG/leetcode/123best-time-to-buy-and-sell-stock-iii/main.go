package main

func maxProfit(prices []int) int {
	n := len(prices)
	max_k := 2
	dp := make([][][]int, n)
	for idx := range dp {
		dp[idx] = make([][]int, max_k+1)
		for iidx := range dp[idx] {
			dp[idx][iidx] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := max_k; k >= 1; k++ {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][max_k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
