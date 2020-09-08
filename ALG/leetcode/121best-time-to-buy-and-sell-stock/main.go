package main

import "math"

func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt32, 0
	for i := range prices {
		curProfit := 0
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			curProfit = prices[i] - minPrice
		}
		maxProfit = max(maxProfit, curProfit)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
