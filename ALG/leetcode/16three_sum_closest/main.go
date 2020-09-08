package main

import (
	"fmt"
	"sort"
)

// Len is the number of elements in the collection.
type arr []int

func (n arr) Len() int {
	return len(n)
}

func (n arr) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n arr) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	var nms arr = nums
	sort.Sort(nms)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		for start != end {
			curResult := nums[i] + nums[start] + nums[end]
			curOffset := curResult - target
			if Abs(curOffset) < Abs(result-target) {
				result = curResult
			}
			if curOffset > 0 {
				end--
			} else if curOffset < 0 {
				start++
			} else {
				return curResult
			}
		}

	}
	return result
}

// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}
