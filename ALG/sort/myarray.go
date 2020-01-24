package sort

import "math/rand"

func RandomArray(length, max int) []int {
	arr := make([]int, length)
	for idx := range arr  {
		arr[idx] = rand.Intn(max)
	}
	return arr
}