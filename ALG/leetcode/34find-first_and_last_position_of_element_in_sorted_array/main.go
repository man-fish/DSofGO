package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	middle, start, end := 0, 0, n-1
	for start <= end {
		middle = (start+end)/2
		if nums[middle] == target {
			start, end = middle, middle
			break
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	if n == 0 || nums[middle] != target {
		return []int{-1, -1}
	}
	for start >= 0 && nums[start] == target {
		start--
	}
	for end <= n - 1 && nums[end] == target {
		end++
	}

	return []int{start+1, end-1}
}

func main() {
	res := searchRange([]int{4, 5, 6, 7, 7, 8}, 7)
	fmt.Println(res)
}