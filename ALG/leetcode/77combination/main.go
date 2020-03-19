package main

import "fmt"

func combine(n int, k int) [][]int {
	recs := make([][]int, 0)
	rec := make([]int, k)
	handler(0, 0, n, k, rec, &recs)
	return recs
}

func handler(idx, lst, n, k int, rec []int, recs *[][]int) {
	if idx == k {
		*recs = append(*recs, rec)
		return
	}

	// if k-idx > n-lst {
	// 	return
	// }

	for i := lst + 1; k-(idx+1) <= n-i; i++ {
		rec[idx] = i
		handler(idx+1, i, n, k, append(make([]int, 0), rec...), recs)
	}
}

func main() {
	rec := combine(4, 2)
	fmt.Println(rec)
}
