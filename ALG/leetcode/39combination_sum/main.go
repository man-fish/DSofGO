package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	cur := make([]int, 0)
	res := make([][]int, 0)
	handler(candidates, cur, 0, target, 0, &res)
	return res
}

func handler(candidates, cur []int, i, target, sum int, res *[][]int) {
	if target == sum {
		*res = append(*res, cur)
		return
	}

	for ; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			continue
		}
		tmp := make([]int, 0)
		tmp = append(tmp, cur...)
		tmp = append(tmp, candidates[i])
		handler(candidates, tmp, i, target, sum+candidates[i], res)
	}
}

func main() {
	// candidates = [2,3,6,7], target = 7,
	candidates := []int{2, 3, 5}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
