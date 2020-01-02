package knight

/**
骑士游历问题
棋盘大小是8*8，骑士在棋盘任一方格开始游历。要求骑士游历棋盘的每一个方格且每个方格只游历一次。输出骑士的游历路径。
解题思路：
定义8*8的棋盘
0:表示没有探索
1-n：表示已经走过了，并且用n记录步骤。
外围的墙也用1表示
*/
//func btknight() {
//	checkBoard := make([][]int, 9)
//	for row := range checkBoard {
//		checkBoard[row] = make([]int, 9)
//	}
//	for i:=0; i<9 ;i++  {
//		checkBoard[0][i] = 1
//		checkBoard[8][i] = 1
//	}
//	for j:=0 ;j<9 ;j++  {
//		checkBoard[j][0] = 1
//		checkBoard[j][8] = 1
//	}
//	printBoard(checkBoard)
//	solveKnight(ans,checkBoard,1,1,1)
//	printBoard(checkBoard)
//}
/*构建棋盘开始游戏*/
/**
checkboard	棋盘
i,j			马的位置
*/
//func solveKnight(ans *[][]string, checkBoard [][]int, i,j int,n int) bool {
//	if n == (len(checkBoard)-2)*(len(checkBoard)-2) + 1 {
//		r := recordKnight(checkBoard)
//		*ans = append(*ans,r)
//		return true
//	}
//
//	if i < 0 || i >= len(checkBoard) || j < 0 || j >= len(checkBoard) {
//		return false
//	}
//
//	for a := -2; a <= 2 ;a++  {
//		b := 0
//		if math.Abs(float64(a)) == 2 {
//			b = 1
//		}else{
//			b = 2
//		}
//		moveKnight(a,b)
//		moveKnight(a,-b)
//	}
//
//
//}

//func moveKnight(i,j,x,y,n int,checkBoard [][]int,ans *[][]string) bool {
//	if checkBoard[i][j] == 0 {
//		checkBoard[i][j] = n
//		n = n + 1
//		if solveKnight(ans,checkBoard,i+y,j+x,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i+1,j+2,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i-1,j+2,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i-2,j+1,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i-2,j-1,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i-1,j-1,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i+1,j-2,n) {
//			return true
//		}else if solveKnight(ans,checkBoard,i+2,j-1,n) {
//			return true
//		}else{
//			return false
//		}
//	}else{
//		return false
//	}
//}


/*回溯法解决问题*/
