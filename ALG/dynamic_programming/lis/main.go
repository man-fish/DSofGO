package main

import (
	"fmt"
	"math"
)

/*
	longest increasing substring 最长递增字串
		给定数组arr，返回arr的最长递增子序列。
*/


func lislis (arr []int) int {
	bestAns := math.MinInt64
	for i := len(arr)-1;i >= 0 ;i--  {
		bestAns = int(math.Max(float64(bestAns),float64(lis(arr,i))))
	}

	return bestAns
}

func lis(arr []int,n int) int {
	if n == 0 {
		return 1
	}

	bestAns := math.MinInt64

	for i := n ;i > 0 ;i-- {
		if arr[n] > arr[n-i] {
			bestAns = int(math.Max(float64(bestAns),float64(lis(arr,n-i)+1)))
		}
	}

	return bestAns
}

func main(){
	var testArr []int = []int{1,5,3,6,19,2}
	fmt.Println(lislis(testArr))
}