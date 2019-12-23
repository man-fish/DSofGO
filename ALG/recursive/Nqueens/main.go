package main

import (
	"fmt"
	"math"
)

/**
	N皇后问题
 */

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	checkboard := make([]int, n)
	calcQueues(&ans,n,checkboard,0)
	return ans
}
	/**
		ans 	*[][]string		全部题解
		n 		int				兼容皇后数量
		queues 	[]int			本轮皇后排列情况
		row 	int				轮数
	 */
func calcQueues(ans *[][]string,n int,queues []int,row int) {
	if row == n {
		q := printQ(queues, n)
		*ans = append(*ans, q)
		return
	}

	for col := 0; col < n; col++ {
		queues[row] = col
		if isValid(queues,row) {
			calcQueues(ans,n,queues,row+1)
		}
	}
}

func isValid(queues []int,row int) bool {
	for i := 0;i < row; i++ {
		if queues[i] == queues[row] || math.Abs(float64(row-i)) == math.Abs(float64(queues[row] - queues[i])){
			return false
		}
	}
	return true
}

func printQ(res []int, n int) []string {
	s := []string{}
	for _, v := range res {
		str := ""
		for i := 0; i < n; i++ {
			if i == v {
				str += "Q"
			} else {
				str += "."
			}
		}
		s = append(s, str)
	}
	return s
}

func main() {
	for _,col := range solveNQueens(4) {
		for _,row := range col{
			fmt.Println(row)
		}
	}
}
