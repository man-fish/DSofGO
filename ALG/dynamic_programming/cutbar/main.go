package main

func cut(p []int, n int) int {
	if n == 0 {
		return 0
	}
	valid := 0
	for idx := range p {
		if valid > p[idx] + cut(p,n-1) {

		}
	}
}

func main() {

}