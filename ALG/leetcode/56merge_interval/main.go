/*
	给出一个区间的集合，请合并所有重叠的区间。
	示例 1:
		输入: [[1,3],[2,6],[8,10],[15,18]]
		输出: [[1,6],[8,10],[15,18]]
		解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
	题解：
		按区间起点排序，判断起点是否和上一区间结尾重合，重合就合并。
*/
package main

import (
	"fmt"
	"math"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	// preint := make([][]int, 0)
	quickSortInterval(intervals, 0, len(intervals)-1)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			// if i >= 2 {
			// 	preint =
			// }
			intervals = append(append(intervals[:i-1], [][]int{[]int{intervals[i-1][0], int(math.Max(float64(intervals[i][1]), float64(intervals[i-1][1])))}}...), intervals[i+1:]...)
			// if i >= 2 {
			// 	intervals = append(preint, intervals...)
			// }
			i--
		}
	}
	return intervals
}

func quickSortInterval(intervals [][]int, start, end int) {
	if len(intervals) == 0 {
		return
	}
	l := start
	r := end
	pivot := intervals[(l+r)/2][0]
	for l < r {
		for intervals[l][0] < pivot {
			l++
		}
		for intervals[r][0] > pivot {
			r--
		}

		if r <= l {
			break
		}
		intervals[l], intervals[r] = intervals[r], intervals[l]
		if pivot == intervals[r][0] {
			l++
		}
		if pivot == intervals[l][0] {
			r--
		}

	}
	if l == r {
		l++
		r--
	}
	if l < end {
		quickSortInterval(intervals, l, end)
	}
	if r > start {
		quickSortInterval(intervals, start, r)
	}
}

func main() {
	intervals := [][]int{[]int{0, 1}, []int{1, 3}, []int{2, 6}, []int{8, 10}}
	intervals = [][]int{[]int{2, 3}, []int{2, 2}, []int{3, 3}, []int{1, 3}, []int{5, 7}, []int{2, 2}, []int{4, 6}}
	intval := merge(intervals)
	fmt.Println(intval)
}
