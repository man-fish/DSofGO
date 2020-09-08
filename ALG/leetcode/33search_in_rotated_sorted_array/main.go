package main

import "fmt"

func search(nums []int, target int) int {
	rotatedPoint := findRotatedPoint(nums)
	var start, end int
	if target < nums[0] {
		start = rotatedPoint+1
		end = len(nums)-1
		if rotatedPoint == end {
			start = 0
		}
	} else {
		start = 0
		end = rotatedPoint
	}
	fmt.Printf("start:%d, end:%d\n", start, end)
	for start <= end {
		middle := (start + end)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1 
		}
	}

	return -1
}

func findRotatedPoint(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		middle := (start + end)/2
		if nums[middle] > nums[middle+1] {
			return middle
		}
		if nums[0] > nums[middle] {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return len(nums)-1
}

func main() {
	nums := []int{4,5,6,7}
	idx := search(nums, 4)
	fmt.Println(idx)
}