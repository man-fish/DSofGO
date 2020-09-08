package main

import (
	"fmt"
	"strconv"
)

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	hanlder(digits, 0, "", &res)
	return res
}

func hanlder(digits string, idx int, cur string, res *[]string) {
	if idx == len(digits) {
		*res = append(*res, cur)
		return
	}
	curNum, _ := strconv.Atoi(digits[idx : idx+1])

	start := (curNum-1)*3 + 94
	end := (curNum-1)*3 + 96

	if curNum == 7 {
		end = end + 1
	}

	if curNum == 8 {
		start = start + 1
		end = end + 1
	}

	if curNum == 9 {
		start = start + 1
		end = end + 2
	}

	for i := start; i <= end; i++ {
		hanlder(digits, idx+1, cur+string(i), res)
	}
}

//			2
// 			2			3
// a	b	c	d	e	f
// 97	98	99	100	101	102
// mappingletter =[i - 2 + 97 : i + 97\

func main() {
	res := letterCombinations("9")
	fmt.Println(res)
}
