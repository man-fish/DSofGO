package main

import (
	"fmt"
	"math"
)

func lcsRecursive(m,n []int,i,j int) int {
	if i == 0 || j == 0 {
		return 1
	}
	if m[i] == n[j] {
		return 1 + lcsRecursive(m,n,i-1,j-1)
	}else{
		return int(math.Max(float64(lcsRecursive(m,n,i-1,j)),float64(lcsRecursive(m,n,i,j-1))))
	}
}

func lcsRynamic(m,n []int) int {
	if len(m) == 0 || len(n) == 0 {
		return 0
	}

	linear_dp := make([][]int, len(m)+1)

	for idx := range linear_dp {
		linear_dp[idx] = make([]int, len(n)+1)
	}

	for i := 1; i < len(linear_dp); i++  {
		for j := 1; j < len(linear_dp[i]); j++  {
			if m[i-1] == n[j-1] {
				linear_dp[i][j] = linear_dp[i-1][j-1] + 1
			}else{
				linear_dp[i][j] = int(math.Max(float64(linear_dp[i-1][j]),float64(linear_dp[i][j-1])))
			}
		}
	}
	for idx := range linear_dp {
		fmt.Println(linear_dp[idx])
	}
	return linear_dp[len(m)][len(n)]
}

func main() {
	testArr2 := []int{1,5,3,6,19,2}
	testArr := []int{3,6,2}
	fmt.Println(lcsRecursive(testArr,testArr2, len(testArr)-1, len(testArr2)-1))
	fmt.Println(lcsRynamic(testArr,testArr2))
}
