package main

import "fmt"

func getFibNum(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return getFibNum(n-1) + getFibNum(n-2)
}


func main () {
	fmt.Println(getFibNum(4))
}