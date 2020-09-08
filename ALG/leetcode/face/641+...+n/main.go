package main

import "fmt"

func quickMulti(A, B int) int {
	ans := 0
	for ; B != 0; B >>= 1 {
		if B&1 == 1 {
			ans += A
		}
		A <<= 1
	}
	return ans
}

func sumNums(n int) int {
	ans := 0
	qm(n, n+1, &ans)
	return ans >> 1
}

func qm(a, b int, ans *int) bool {
	cond := b&1 > 0

	addGreatZero := func() bool {
		*ans += a
		return *ans > 0
	}

	_ = cond && addGreatZero()
	a <<= 1
	b >>= 1

	_ = (b != 0) && qm(a, b, ans)
	return true
}

func main() {
	res := quickMulti(100, 100)
	res = sumNums(3)
	fmt.Println(res)
}
