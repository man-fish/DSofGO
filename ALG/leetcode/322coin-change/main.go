package main

import (
	"fmt"
	"math"
)

const maxCount = math.MaxInt32

func coinChangeTimeOut(coins []int, amount int) int {
	if count := dfsTimeOut(coins, amount, 0, 0); count != maxCount {
		return count
	}
	return -1
}

func dfsTimeOut(coins []int, amount int, cur int, ccount int) int {
	if cur == amount {
		return ccount
	}
	curBestChoice := maxCount
	for idx := range coins {
		if csum := cur + coins[idx]; csum <= amount {
			curBestChoice = min(curBestChoice, dfsTimeOut(coins, amount, csum, ccount+1))
		} else {
			break
		}
	}
	return curBestChoice
}

func coinChange(coins []int, amount int) int {
	return dfs(coins, amount, make([]int, amount))
}

func dfs(coins []int, rem int, dp []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if dp[rem-1] != 0 {
		return dp[rem-1]
	}
	min := math.MaxInt64
	for _, coin := range coins {
		if res := dfs(coins, rem-coin, dp); res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == maxCount {
		dp[rem-1] = -1
	} else {
		dp[rem-1] = min
	}

	return dp[rem-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	res := coinChange([]int{1, 2, 5}, 11)
	fmt.Println(res)
}
