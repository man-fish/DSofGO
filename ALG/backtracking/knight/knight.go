package main

import "fmt"

/**
	骑士游历问题
	棋盘大小是8*8，骑士在棋盘任一方格开始游历。要求骑士游历棋盘的每一个方格且每个方格只游历一次。输出骑士的游历路径。
	解题思路：
	定义8*8的棋盘
	0:表示没有探索
	1-n：表示已经走过了，并且用n记录步骤。
	外围的墙也用1表示
 */
func btknight() {
	checkBoard := make([][]int, 9)
	for row := range checkBoard {
		checkBoard[row] = make([]int, 9)
	}
	for i:=0; i<9 ;i++  {
		checkBoard[0][i] = 1
		checkBoard[8][i] = 1
	}
	for j:=0 ;j<9 ;j++  {
		checkBoard[j][0] = 1
		checkBoard[j][8] = 1
	}
	printBoard(checkBoard)
	solveKnight(checkBoard,1,1,1)
	printBoard(checkBoard)
}
/*构建棋盘开始游戏*/
/**
	checkboard	棋盘
	i,j			马的位置
 */
func solveKnight(checkBoard [][]int, i,j int,n int) bool {
	if isOver(checkBoard) {
		return true
	}

	if i < 0 || i >= len(checkBoard) || j < 0 || j >= len(checkBoard) {
		return false
	}

	if checkBoard[i][j] == 0 {
		checkBoard[i][j] = n
		n = n + 1
		if solveKnight(checkBoard,i+2,j+1,n) {
			return true
		}else if solveKnight(checkBoard,i+1,j+2,n) {
			return true
		}else if solveKnight(checkBoard,i-1,j+2,n) {
			return true
		}else if solveKnight(checkBoard,i-2,j+1,n) {
			return true
		}else if solveKnight(checkBoard,i-2,j-1,n) {
			return true
		}else if solveKnight(checkBoard,i-1,j-1,n) {
			return true
		}else if solveKnight(checkBoard,i+1,j-2,n) {
			return true
		}else if solveKnight(checkBoard,i+2,j-1,n) {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}
/*回溯法解决问题*/
func isOver(checkBoard [][]int) bool {
	for _,row := range checkBoard {
		for _,col := range row {
			if col == 0 {
				return false
			}
		}
	}
	return true
}
/*如果全部位置都已经走过了的话游戏结束*/
func printBoard(checkBoard [][]int) {
	for _,line := range checkBoard {
		for _,block := range line {
			fmt.Print(block,"\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
/*打印期盼*/
func main () {
	btknight()
}
