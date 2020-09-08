package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i := start
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	// reverse(nums, 1)
	fmt.Println(nums)
}
