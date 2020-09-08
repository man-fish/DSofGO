package main

import (
	"go_dataStruct/util"
	"sort"
)

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func fourSum(nums []int, target int) [][]int {
	ress := make([][]int, 0)
	sort.Sort(SortBy(nums))
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			start := j + 1
			end := len(nums) - 1
			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					ress = append(ress, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
					for start < end && start != j+1 && nums[start] == nums[start-1] {
						start++
					}
					for start < end && end != len(nums)-1 && nums[end] == nums[end+1] {
						end--
					}
				} else if sum > target {
					end--
				} else {
					start++
				}

			}
		}
	}
	return ress
}

/*
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	res := fourSum(nums, 0)
	util.Print2DArray(res)
}
