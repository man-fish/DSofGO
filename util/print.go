package util

import "fmt"

func Print2DArray(arr [][]int) {
	for _, item := range arr {
		fmt.Println(item)
	}
}
