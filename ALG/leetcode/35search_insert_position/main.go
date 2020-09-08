package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		}
		if start == end || start == end-1 {
			if nums[end] < target {
				return end + 1
			} else if nums[start] > target {
				return start
			} else {
				return end
			}
		}

		if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return 0
}

func main() {
	nums := []int{3, 5, 7, 9, 10}

	idx := searchInsert(nums, 8)

	fmt.Println(idx)
}
