package main

import "fmt"

/*
一条包含字母 A-Z 的消息通过以下方式进行了编码：
	'A' -> 1
	'B' -> 2
	...
	'Z' -> 26
*/

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))

	if s[len(s)-1] != '0' {
		dp[len(s)-1] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			continue
		}
		if s[i] == 49 || (s[i] == 50 && s[i+1] <= 54) {

			if i == len(s)-2 {
				dp[i] = 1 + dp[i+1]
				continue
			}
			dp[i] = dp[i+1] + dp[i+2]
			continue
		}

		dp[i] = dp[i+1]
	}
	return dp[0]
}

func numDecodingsV2(s string) int {
	if s[0] == '0' {
		return 0
	}
	start := len(s) - 1
	if s[start] == '0' {
		return foo(s, start)
	} else {
		return foo(s, start) + 1
	}
}

func foo(s string, idx int) int {
	idx--
	if idx == -1 {
		return 0
	}
	if s[idx] == 49 && s[idx+1] != '0' || (s[idx] == 50 && s[idx+1] <= 54) {
		return foo(s, idx) + 1
	} else {
		return foo(s, idx)
	}
}

func main() {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodingsV2("100"))
}
