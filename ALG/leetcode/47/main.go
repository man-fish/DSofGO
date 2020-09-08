package main

import (
	"fmt"
	"sort"
)

type SortByInt []int

func (a SortByInt) Len() int           { return len(a) }
func (a SortByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByInt) Less(i, j int) bool { return a[i] < a[j] }

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Sort(SortByInt(nums))
	dfs(nums, 0, &res)
	return res
}

func dfs(nums []int, idx int, res *[][]int) {
	if idx == len(nums)-1 {
		*res = append(*res, nums)
		return
	}
	for i := idx; i < len(nums); i++ {
		if i != idx && nums[i] == nums[idx] {
			continue
		}
		nums[idx], nums[i] = nums[i], nums[idx]
		tmp := append(make([]int, 0), nums...)
		fmt.Println(tmp)
		dfs(tmp, idx+1, res)
	}
}

func main() {
	res := permuteUnique([]int{1, 1, 2})
	fmt.Println(res)
}
