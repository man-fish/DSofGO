package main

import "fmt"

type mymap [8][7]int

func dailMaze(maze *mymap, i int, j int) bool {
	/**
	下：i+1
	右：j+1
	上：i-1
	左：j-1
	1 ：墙
	2：通路
	3：死路
	4：未探索
	*/
	if maze[6][5] == 2 {
		fmt.Println("congratulation for finish the maze ！")
		return true
	}

	if maze[i][j] == 0 {
		maze[i][j] = 2
		if dailMaze(maze,i+1,j) {
			return true
		}else if dailMaze(maze,i,j+1) {
			return true
		}else if dailMaze(maze,i-1,j) {
			return true
		}else if dailMaze(maze,i,j-1) {
			return true
		}else{
			maze[i][j] = 3
			return false
		}
	}else{
		return false
	}
}

func main() {
	var maze mymap
	for i:=0; i<7 ;i++  {
		maze[0][i] = 1
		maze[7][i] = 1
	}

	for j:=0 ;j<8 ;j++  {
		maze[j][0] = 1
		maze[j][6] = 1
	}

	maze[4][1] = 1
	maze[4][2] = 1
	//maze[1][2] = 1
	maze[2][2] = 1
	maze[3][2] = 1

	fmt.Println(dailMaze(&maze,1,1))

	for i:=0;i<8 ;i++  {
		for j:=0;j<7;j++ {
			fmt.Print(maze[i][j],"\t")
		}
		fmt.Println()
	}
}