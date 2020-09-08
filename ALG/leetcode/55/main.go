package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	visited := append(make([]int, len(nums)), nums...)
	return dfs(nums, visited, 0)
}

func dfs(nums, visited []int, idx int) bool {
	if idx == len(nums)-1 {
		return true
	}
	if nums[idx] == 0 {
		return false
	}
	visited[idx] = 1
	res := false
	for i := 1; i <= nums[idx]; i++ {
		if visited[idx+i] == 1 {
			continue
		}
		res = dfs(nums, visited, idx+i)
		if res {
			return res
		}
	}
	return res
}

func main() {
	arr := []int{3, 2, 1, 0, 4}
	res := canJump(arr)
	fmt.Println(res)
}
