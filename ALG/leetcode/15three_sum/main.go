/*
	threeSum
	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
	注意：答案中不可以包含重复的三元组。
	1. 排序（O(NlogN)O(NlogN)）
		为什么要排序呢？我想是不是这样：
		排序后相同的数会挨在一起，所以方便接下来的操作来避免重复；
		右边的数比左边的大，方便比较。
	2. 遍历（O(N^2)O(N2)）
		我们定义当前下标 k，i, j 分别位于 k + 1, len(nums) - 1的位置。若 k = 1，如下图所示：
		从左到右遍历，对于每个 k 值：
		如果三数之和大于 0，j 左移会使得和减小，甚至等于 0，故 j 左移；
		如果三数之和小于 0，i 右移会使得和增大，甚至等于 0，故 i 右移；
		如果三数之和等于 0，将 [nums[i],nums[j],nums[k]] 加入到结果中，i 右移，j 左移，寻找下一组可能结果。

*/
package main

import (
	"fmt"
)

// [-3, -1, 0, 1, 2]
//	k    i        j
func threeSum(nums []int) [][]int {
	var (
		k int
		i int
		j int

		threeSums [][]int
	)
	BubbleSort(nums)
	for ; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i, j = k+1, len(nums)-1
		for i < j {
			if nums[k]+nums[i]+nums[j] > 0 {
				j--
				for nums[j] == nums[j+1] && i < j {
					j--
				}
			} else if nums[k]+nums[i]+nums[j] < 0 {
				i++
				for nums[i] == nums[i-1] && i < j {
					i++
				}
			} else {
				threeSums = append(threeSums, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for nums[i] == nums[i-1] && i < j {
					i++
				}
				for nums[j] == nums[j+1] && i < j {
					j--
				}
			}
		}
	}
	return threeSums
}

func BubbleSort(arr []int) {
	exchange := true
	for j := 1; j < len(arr) && exchange; j++ {
		exchange = false
		for i := 0; i < len(arr)-j; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				exchange = true
			}
		}
	}
}

func main() {
	fmt.Println(threeSum([]int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}))
}
