package main

import "fmt"

func fib (n int) int {
	if n <= 0 {
		return n
	}
	arr := make([]int ,n+1)
	for idx := range arr {
		arr[idx] = -1
	}
	return fibDypBackward(arr,n)
}

func fibDypBackward(arr []int, n int) int {
	if arr[n] != -1 {
		return arr[n]
	}
	if n <= 2 {
		return 1
	}
	arr[n] = fibDypBackward(arr, n-1) + fibDypBackward(arr, n-2)
	return arr[n]
}

func fibDypForward(n int) int {
	if n <= 0 {
		return n
	}
	arr := make([]int ,n+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func main() {
	fmt.Println(fibDypForward(3))
	fmt.Println(fib(3))
}
