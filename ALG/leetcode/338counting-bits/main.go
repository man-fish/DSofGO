package main

func countBits(num int) []int {
	sum := make([]int, num+1)
	sum[0] = 0
	for i := 1; i <= num; i++ {
		sum[i] = 1 + sum[(i-1)&i]
	}
	return sum
}

func main() {
	countBits(5)
}
