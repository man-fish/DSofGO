func jump(nums []int) int {
	count := 0
	max := math.MaxInt16
	dfs(nums, count, 0, &max)
	return max
}

func dfs(nums []int, count, idx int, max *int) {
	if idx == len(nums)-1 {
		if *max > count {
			*max = count
		}
		return
	}
	if nums[idx] == 0 {
		return
	}
	for i := 1; i <= nums[idx] && idx+i < len(nums); i++ {
		dfs(nums, count+1, idx+i, max)
	}
}
