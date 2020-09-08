package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	res := 0
	flag := false

	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}

		if dividend < 0 {
			return -dividend
		}

		return -dividend
	}

	if divisor == 1 {
		return dividend
	}

	if dividend < 0 || divisor < 0 {
		if !(dividend < 0 && divisor < 0) {
			flag = true
		}
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	for dividend >= divisor {
		dividend = dividend - divisor
		res++
	}

	if flag {
		res = -res
	}

	return res
}

func main() {
	res := divide(-2147483648,
		-1)
	fmt.Println(res)
}
