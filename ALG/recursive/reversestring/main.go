package main

import "fmt"

//1,2,3,4,5 -> 5,4,3,2,1

func reverseString(str []byte) []byte {
	if len(str) == 0 {
		return nil
	}
	return append(reverseString(str[1:]),str[0])
}

func main() {
	fmt.Println(reverseString([]byte{1,2,3,4,5}))
}