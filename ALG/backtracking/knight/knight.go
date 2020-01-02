package knight

import "fmt"

/**
	骑士游历问题:
		棋盘大小是8*8，骑士在棋盘任一方格开始游历。要求骑士游历棋盘的每一个方格且每个方格只游历一次。输出骑士的游历路径。
	解题思路：
		定义8*8的棋盘
		0:表示没有探索
		1-n：表示已经走过了，并且用n记录步骤。
		外围的墙也用1表示
*/

func SolveKnight() {
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
	knight(checkBoard,1,1,1)
	printBoard(checkBoard)
	ans := recordKnight(checkBoard)
	fmt.Println(ans)
}
/*构建棋盘开始游戏*/
/**
	checkboard	棋盘
	i,j			马的位置
*/
func knight(checkBoard [][]int, i,j int,n int) bool {
	if n == (len(checkBoard)-2)*(len(checkBoard)-2) + 1 {
		return true
	}

	if i < 0 || i >= len(checkBoard) || j < 0 || j >= len(checkBoard) {
		return false
	}

	if checkBoard[i][j] == 0 {
		checkBoard[i][j] = n
		n = n + 1
		if knight(checkBoard,i+2,j+1,n) {
			return true
		}else if knight(checkBoard,i+1,j+2,n) {
			return true
		}else if knight(checkBoard,i-1,j+2,n) {
			return true
		}else if knight(checkBoard,i-2,j+1,n) {
			return true
		}else if knight(checkBoard,i-2,j-1,n) {
			return true
		}else if knight(checkBoard,i-1,j-1,n) {
			return true
		}else if knight(checkBoard,i+1,j-2,n) {
			return true
		}else if knight(checkBoard,i+2,j-1,n) {
			return true
		}else{
			return false
		}
	}else{
		return false
	}
}
/*如果全部位置都已经走过了的话游戏结束,已经进行了优化*/
func printBoard(checkBoard [][]int) {
	for _,line := range checkBoard {
		for _,block := range line {
			fmt.Print(block,"\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func recordKnight(checkboard [][]int) []string {
	ans := make([]string,(len(checkboard)-2)*(len(checkboard)-2))
	for i := 1; i < len(checkboard)-2; i++  {
		for j := 1; j < len(checkboard)-2 ; j++  {
			ans[checkboard[i][j]-1] = fmt.Sprintf("(%v,%v)",j,i)
		}
	}
	return ans
}
/*打印棋盘*/
