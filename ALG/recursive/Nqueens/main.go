package main

import "fmt"

type checkboard [4][4]int

/**
	0 没有试过
	1 表示成功放置皇后
	2 表示位置冲突
 */

func setQueen(board checkboard,i int,j int) bool {
	for f := 0;f < len(board[3]) ;f++  {
		if board[3][f] == 1 {
			return  true
		}
	}

	if board[i][j] == 0 {

	}else{
		return false
	}
}

func main() {
	var board checkboard
	for i:=0;i<4 ;i++ {
		for j:=0;j<4 ;j++  {
			fmt.Print(board[i][j],"\t")
		}
		fmt.Println()
	}
}
