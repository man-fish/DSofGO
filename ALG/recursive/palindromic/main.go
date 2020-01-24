package main

import "fmt"
/**
	回文序列
*/
func isPalindromic(str []byte) bool {
	start := 0
	end := len(str)-1
	if start == end {
		return true
	}
	if str[start] != str[end] {
		return false

	}
	isP := isPalindromic(str[1:len(str)-1])
	return isP
}

func main() {
	fmt.Println(isPalindromic([]byte("aada")))
}