package main

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	maxLen := 0
	dp := make([]int, n)

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			} else if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			}
		}

		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := longestValidParentheses("())")
	fmt.Println(res)
}
