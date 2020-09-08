package main

import "fmt"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	width := min(C, G) - max(A, E);
	height := min(H, D) - max(B, F);
	commonSpace := max(width, 0) * max(height, 0)
    eachSpace := (C - A) * (D - B) + (G - E) * (H - F);
    return eachSpace - commonSpace;
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}