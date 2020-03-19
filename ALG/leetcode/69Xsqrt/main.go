package main

import "fmt"

func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 0
}

func mySqrtV2(x int) int {
	start := 0
	end := (x / 2) + 1
	var pivot, num int

	for start <= end {
		pivot = (end + start) / 2
		num = pivot * pivot
		if num > x {
			end = pivot - 1
		} else if num < x {
			start = pivot + 1
		} else {
			return pivot
		}
	}
	return end
}

func main() {
	for i := 0; i < 16; i++ {
		fmt.Println(i, mySqrtV2(i))
	}

}
