package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	maxS := 0
	for i := 0; i < n; i++ {
		if i != 0 && heights[i] < heights[i-1] {
			continue
		}
		minY := heights[i]
		for j := i; j < n; j++ {
			x := j - i + 1
			if minY > heights[j] {
				minY = heights[j]
			}
			maxS = max(maxS, x*minY)
		}
	}
	return maxS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {}
