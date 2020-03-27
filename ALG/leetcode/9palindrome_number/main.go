package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := strconv.Itoa(x)
	reverseIdx := len(a) - 1
	for i := 0; i <= reverseIdx/2; i++ {
		if a[i] != a[reverseIdx-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1001))
}
