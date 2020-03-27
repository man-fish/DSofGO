package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	y := 0

	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}

	if x < 0 {
		y = -y
	}

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	return y
}

func main() {
	res := reverse(-100)
	fmt.Println(res)
}
