/*
给出一个无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
示例 1:
	输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
	输出: [[1,5],[6,9]]
方法：贪心
算法：
	将 newInterval 之前开始的区间添加到输出。
	添加 newInterval 到输出，若 newInterval 与输出中的最后一个区间重合则合并他们。
	一个个添加区间到输出，若有重叠部分则合并他们。
*/
package main

import "fmt"

// —— —————— —— ————————    ————
//       ——————
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var (
		result = make([][]int, 0)
		i      = 0
		s, e   = newInterval[0], newInterval[1]
	)

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])
			i++
			continue
		}
		if intervals[i][0] < newInterval[0] {
			s = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			e = intervals[i][1]
		}
		i++
	}

	result = append(append(result, []int{s, e}), intervals[i:]...)
	return result
}

// [][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}

func main() {
	var (
		intervals    = [][]int{[]int{1, 5}}
		newInterval  = []int{2, 3}
		newIntervals [][]int
	)
	newIntervals = insert(intervals, newInterval)
	fmt.Println(newIntervals)
}
