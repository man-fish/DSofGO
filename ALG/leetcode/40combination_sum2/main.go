package main

import (
	"fmt"
	"sort"
)

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func combinationSum2(candidates []int, target int) [][]int {
	cur := make([]int, 0)
	res := make([][]int, 0)
	sort.Sort(SortBy(candidates))
	handler(candidates, cur, 0, target, 0, &res)
	return res
}

func handler(candidates, cur []int, i, target, sum int, res *[][]int) {
	if target == sum {
		*res = append(*res, cur)
		return
	}
	flag := false
	for ; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			continue
		}
		if flag && candidates[i] == candidates[i-1] {
			continue
		}
		flag = true
		tmp := make([]int, 0)
		tmp = append(tmp, cur...)
		tmp = append(tmp, candidates[i])
		handler(candidates, tmp, i+1, target, sum+candidates[i], res)
	}
}

func main() {
	// candidates = [2,3,6,7], target = 7,
	candidates := []int{1, 1, 2, 5}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}
