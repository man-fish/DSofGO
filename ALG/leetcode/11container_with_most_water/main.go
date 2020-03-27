package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	lheight, rheight, space := 0, len(height)-1, 0
	for lheight < rheight {
		lh, rh := height[lheight], height[rheight]
		if lh < rh {
			space = int(math.Max(float64(space), float64(lh*(rheight-lheight))))
			lheight++
		} else {
			space = int(math.Max(float64(space), float64(rh*(rheight-lheight))))
			rheight--
		}
	}
	return space
}

func main() {
	heights := []int{2, 3, 4, 5, 18, 17, 6}
	s := maxArea(heights)
	fmt.Println(s)
}
