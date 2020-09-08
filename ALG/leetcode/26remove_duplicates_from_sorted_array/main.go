package main

import "fmt"

func removeDuplicates(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 6, 8}
	idx := removeDuplicates(nums)
	fmt.Println(idx, nums)
}
