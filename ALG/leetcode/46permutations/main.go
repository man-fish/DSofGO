package main

import "fmt"

func permute(nums []int) [][]int {
	visited := make([]int, len(nums))
	for i := range visited {
		visited[i] = 0x0000ffff
	}
	recs := make([][]int, 0)
	for i := range nums {
		v := append(make([]int, 0), visited...)
		v[i] = nums[0]
		handler(nums, v, 1, &recs)
	}
	return recs
}

func handler(nums []int, visited []int, idx int, recs *[][]int) {
	isOver := true
	for i := 0; i < len(nums); i++ {
		if visited[i] == 0x0000ffff {
			isOver = false
			v := append(make([]int, 0), visited...)
			v[i] = nums[idx]
			handler(nums, v, idx+1, recs)
			continue
		}

		if isOver && i == len(nums)-1 && visited[i] != 0x0000ffff {
			*recs = append(*recs, visited)
			return
		}
	}
}

func permuteV2(nums []int) [][]int {
	recs := make([][]int, 0)
	handlerV2(nums, 0, len(nums), &recs)
	return recs
}

func handlerV2(nums []int, first int, n int, recs *[][]int) {
	if n == first {
		*recs = append(*recs, nums)
		return
	}
	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]

		handlerV2(append(make([]int, 0), nums...), first+1, n, recs)
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permuteV2(nums))
}
