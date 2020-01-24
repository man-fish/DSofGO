package factorial

import "fmt"

/**
	n的阶乘计算
 */
func getFactorial(n int) int {
	if n <= 0 {
		 return 0
	}
	return n + getFactorial(n-1)
}

func main() {
	fmt.Println(getFactorial(5))
}