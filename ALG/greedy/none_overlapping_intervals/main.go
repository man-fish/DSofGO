package main

import (
	"fmt"
	"math"
)

type interval [2]int

func IntervalBubbleSort(arr []interval) {
	exchange := true
	for j := 1;j < len(arr) && exchange;j++  {
		exchange = false
		for i := 0;i < len(arr)-j ;i++  {
			if arr[i][0] != arr[i+1][0] {
				if arr[i][0] > arr[i+1][0] {
					arr[i+1],arr[i] = arr[i],arr[i+1]
					exchange = true
				}
			}else{
				if arr[i][1] > arr[i+1][1] {
					arr[i+1],arr[i] = arr[i],arr[i+1]
					exchange = true
				}
			}
		}
	}
}
func IntervalBubbleSortByback(arr []interval) {
	exchange := true
	for j := 1;j < len(arr) && exchange;j++  {
		exchange = false
		for i := 0;i < len(arr)-j ;i++  {
			if arr[i][1] > arr[i+1][1] {
				arr[i+1],arr[i] = arr[i],arr[i+1]
				exchange = true
			}
		}
	}
}

func eraseOverLappingIntervalDyamiced(intervals []interval) []int {
	IntervalBubbleSort(intervals)
	linear_dp := make([]int,len(intervals))
	linear_dp[0] = 1

	for i := 1 ;i < len(intervals); i++ {
		for j := 0; j < i ; j++  {
			if intervals[i][0] > intervals[j][1] {
				linear_dp[i] = int(math.Max(float64(linear_dp[i]),float64(linear_dp[j]+1)))
			}else{
				linear_dp[i] = int(math.Max(float64(linear_dp[i]),float64(linear_dp[j])))
			}
		}
	}
	return linear_dp
}

func eraseOverLappingIntervalGreedily(intervals []interval) int {
	IntervalBubbleSortByback(intervals)

	if len(intervals) == 0 {
		return 0
	}

	pre := 0
	res := 1

	for i := 1; i < len(intervals); i++ {
		if intervals[pre][1] < intervals[i][0] {
			res++
			pre = i
		}
	}

	return res
}

func main() {
	intervals := []interval{[2]int{3,5},[2]int{3,4},[2]int{1,2}}
	IntervalBubbleSortByback(intervals)
	fmt.Println(intervals)

	IntervalBubbleSort(intervals)
	fmt.Println(intervals)

	fmt.Println(eraseOverLappingIntervalDyamiced(intervals))
	fmt.Println(eraseOverLappingIntervalGreedily(intervals))
}