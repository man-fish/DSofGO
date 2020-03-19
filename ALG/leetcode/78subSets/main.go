package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= len(nums); i++ {
		backtracking(nums, i, 0, &res)
	}
	return res
}

func backtracking(nums []int, capacity int, idx int, res *[][]int) {
	if idx == capacity {
		*res = append(*res, nums[0:capacity])
		return
	}

	for i := idx; i < len(nums); i++ {
		cur := make([]int, len(nums))
		copy(cur, nums)
		swap(idx, i, cur)
		backtracking(cur, capacity, idx+1, res)
	}
}

func swap(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func subsetsV2(nums []int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)
	result = append(result, item)
	generatel(0, nums, &item, &result)
	return result
}

func generatel(idx int, nums []int, item *[]int, result *[][]int) {
	if idx >= len(nums) { // 结束
		return
	}
	*item = append(*item, nums[idx])

	// 加入解
	temp := make([]int, len(*item))
	copy(temp, *item)
	*result = append(*result, temp)
	// 纵向拓展
	generatel(idx+1, nums, item, result)
	// 横向拓展
	*item = (*item)[:len(*item)-1]
	generatel(idx+1, nums, item, result)
}

// func subsetsV2(nums []int) [][]int {
// 	result := make([][]int, 0)
// 	item := make([]int, 0)
// 	result = append(result, item)
// 	generatel(0, nums, item, &result)
// 	return result
// }

// func generatel(idx int, nums []int, item []int, result *[][]int) {
// 	if idx >= len(nums) { // 结束
// 		return
// 	}
// 	item = append(item, nums[idx])

// 	// 加入解
// 	temp := make([]int, len(item))
// 	copy(temp, item)
// 	*result = append(*result, temp)
// 	// 纵向拓展
// 	generatel(idx+1, nums, item, result)
// 	// 横向遍历
// 	item = item[:len(item)-1]
// 	generatel(idx+1, nums, item, result)
// }

func slice(s []int) {
	s = append(s, 1) // 底层数组改变，不发生变化。
}

func slice2(s *[]int) {
	slice4(s)          // 传递指针，发生变化
	*s = append(*s, 1) // 传递指针，发生变化。
}

func slice3(s []int) {
	s = s[:len(s)-1] // 底层数组改变，发生变化。
}

func slice4(s *[]int) {
	*s = (*s)[:len(*s)-1] // 底层数组改变，发生变化。
}

func main() {
	set := []int{1, 2, 3}
	_ = subsetsV2(set)
	// for i := range res {
	// 	fmt.Println(res[i])
	// }
	s := []int{1}
	slice2(&s)
	fmt.Println(s)
}
