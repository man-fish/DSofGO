/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例:
	board =
	[
	['A','B','C','E'],
	['S','F','C','S'],
	['A','D','E','E']
	]
	给定 word = "ABCCED", 返回 true.
	给定 word = "SEE", 返回 true.
	给定 word = "ABCB", 返回 false.
*/

package main

import "fmt"

func exist(board [][]byte, word string) bool {
	visited := make([][]int, len(board))
	for i := range visited {
		visited[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtracking(board, visited, word, j, i, 0) {
				return true
			}
		}
	}
	return false
}

func backtracking(board [][]byte, visitor [][]int, word string, x, y, idx int) bool {
	visited := make([][]int, len(visitor))
	for i := range visitor {
		visited[i] = make([]int, len(visitor[i]))
		copy(visited[i], visitor[i])
	}

	if board[y][x] != word[idx] {
		return false
	}

	if idx == len(word)-1 {
		return true
	}

	strategy := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
	// 顺时针试探策略
	visited[y][x] = 1
	flag := false

	for _, direction := range strategy {
		nextX := x + direction[0]
		nextY := y + direction[1]
		if nextX >= 0 && nextY >= 0 && nextX < len(board[0]) && nextY < len(board) && visited[nextY][nextX] == 0 {
			if backtracking(board, visited, word, nextX, nextY, idx+1) {
				flag = true
			}
		}
	}
	return flag
}

func existV2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, j, i, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, x, y, idx int) bool {
	if board[y][x] != word[idx] {
		return false
	}

	if idx == len(word)-1 {
		return true
	}

	tmp := board[y][x]
	board[y][x] = ' '
	// 不需要改变数组的小技巧
	if 0 <= y-1 && dfs(board, word, x, y-1, idx+1) {
		return true
	}
	if y+1 < len(board) && dfs(board, word, x, y+1, idx+1) {
		return true
	}
	if 0 <= x-1 && dfs(board, word, y, x-1, idx+1) {
		return true
	}
	if x+1 < len(board[0]) && dfs(board, word, y, x+1, idx+1) {
		return true
	}

	// strategy := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
	// for _, direction := range strategy {
	// 	nextX := x + direction[0]
	// 	nextY := y + direction[1]
	// 	if nextX >= 0 && nextY >= 0 && nextX < len(board[0]) && nextY < len(board) {
	// 		if dfs(board, word, nextX, nextY, idx+1) {
	// 			return true
	// 		}
	// 	}
	// }

	board[y][x] = tmp
	return false
}

/*
board =
	[
		['A','B','C','E'],
		['S','F','C','S'],
		['A','D','E','E']
	]
*/

func main() {
	// board := [][]byte{
	// 	[]byte{'A', 'B', 'C', 'E'},
	// 	[]byte{'S', 'F', 'C', 'S'},
	// 	[]byte{'A', 'D', 'E', 'E'},
	// }

	// board := [][]byte{
	// 	[]byte{'A'},
	// }

	board := [][]byte{
		[]byte{'C', 'A', 'A'},
		[]byte{'A', 'A', 'A'},
		[]byte{'B', 'C', 'D'},
	}
	res := existV2(board, "AAB")
	fmt.Println(res)
}
