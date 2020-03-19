package main

import "fmt"

func removeElement(nums []int, val int) int {
	idx := 0
	for _, ele := range nums {
		if ele != val {
			nums[idx] = ele
			idx++
		}
	}
	return idx
}

func main() {
	arr := []int{1, 2, 5, 8, 2, 2}
	v := removeElement(arr, 2)
	fmt.Println(v)
	fmt.Println(arr)
}
