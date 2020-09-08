package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	bestAns := 1
	dp(n-1, nums, &bestAns)
	return bestAns
}

func dp(n int, nums []int, bestAns *int) int {
	if n == 0 {
		return 1
	}
	if nums[n] > nums[n-1] {
		cur := dp(n-1, nums, bestAns) + 1
		*bestAns = max(*bestAns, cur)
		return cur
	}
	dp(n-1, nums, bestAns)
	return 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
}
