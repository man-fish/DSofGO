package main

const (
	null = iota
	hand
	freeze
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, 3)
	}
	dp[0][hand] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][hand] = max(dp[i-1][hand], dp[i-1][null]-prices[i])
		dp[i][freeze] = dp[i-1][hand] + prices[i]
		dp[i][null] = max(dp[i-1][freeze], dp[i-1][null])
	}

	return max(dp[n-1][null], dp[n-1][freeze])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxProfit([]int{1, 2, 3, 0, 2})
}
