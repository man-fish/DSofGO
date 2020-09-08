package main

func maxSubArray(nums []int) int {
	length := len(nums)
	sum := nums[0]
	max := sum
	for i := 1; i < length; i++ {
		sum = sum + nums[i] // sum 维护实际上开始的序列现在遍历到index的和。
		if sum < nums[i] {  // 顾前，如果当前元素大于前面所有元素的和，更新实际开始的位置。
			sum = nums[i]
		}
		if sum > max { // 顾后，实际开始的位置不在变动的时候，我们就考虑每次向后遍历的时候当前是不是最大值。
			max = sum
		}
	}
	return max
}
