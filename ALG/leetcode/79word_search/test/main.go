package main

import "fmt"

func main() {
	slice := [][]int{[]int{1}}
	s := make([][]int, len(slice))
	for i := range slice {
		s[i] = make([]int, len(slice[i]))
		copy(s[i], slice[i])
	}
	fmt.Println(slice, s)
	slice[0][0] = 11
	fmt.Println(s)
}
