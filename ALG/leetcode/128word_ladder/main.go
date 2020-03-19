/*
思路：
	 双向广度优先遍历
*/
 package main

import (
	"fmt"
	"go_dataStruct/QUEUE"
)

func ladderLengthBackTracking(beginWord string, endWord string, wordList []string) int {
	ans := make([]int, 1)
	for i, v := range wordList {
		if v == beginWord {
			wordList = append(wordList[:i], wordList[i+1:]...)
		}
	}
	helper(beginWord, endWord, wordList, 1, ans)
	bestMethod := ans[0]
	return bestMethod
}

func helper(beginWord string, endWord string, wordList []string, count int, ans []int) {
	if beginWord == endWord {
		if ans[0] == 0 || count < ans[0] {
			ans[0] = count
		}
		return
	}
	currentCount := count
	for i, v := range wordList {
		count = currentCount
		if conversable(beginWord, v) { //如果说可以转换
			newList := make([]string, len(wordList)) //复制
			copy(newList, wordList)
			if len(wordList) != 1 {
				if len(wordList)-1 != i {
					newList = append(newList[:i], newList[i+1:]...)
				} else {
					newList = newList[:i]
				}
			}
			count++

			if ans[0] != 0 && count >= ans[0] {
				return
			}
			helper(v, endWord, newList, count, ans)
		}
	}
}

func conversable(originWord string, toWord string) bool {
	count := 0
	for i, v := range originWord {
		if toWord[i] == byte(v) {
			count++
		}
	}
	if count == len(originWord)-1 {
		return true
	}
	return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	count := 0
	for i, v := range wordList {
		if v == beginWord {
			wordList = append(wordList[:i], wordList[i+1:]...)
			count++
		}
	}
	visited := make([]bool, len(wordList)) //初始化标记数组
	queue := QUEUE.NewLinkedQueue()
	queue.Add(beginWord)

	for !queue.IsEmpty() {
		size := queue.Size()
		fmt.Println(size)
		count++
		for size > 0 {
			size--
			curEle := queue.Poll()
			for i, v := range wordList {
				if visited[i] {
					continue
				}
				if !conversable(curEle.(string), v) {
					continue
				}

				if v == endWord {
					return count + 1
				}

				queue.Add(v)
				visited[i] = true
			}
		}

	}
	return 0
}

func ladderLengthPro(beginWord string, endWord string, wordList []string) int {
	var (
		hasBeg = false
		count  = 0
		idx1   = -1
		idx2   = -1
	)

	var (
		visited1 []bool
		visited2 []bool
		queue1   = QUEUE.NewLinkedQueue()
		queue2   = QUEUE.NewLinkedQueue()
	)

	for i, v := range wordList {
		if beginWord == v {
			hasBeg = true
			idx1 = i
		}

		if endWord == v {
			idx2 = i
		}
	}

	if idx2 == -1 {
		return 0
	}

	if !hasBeg {
		wordList = append(wordList, beginWord)
		idx1 = len(wordList) - 1
	}

	visited1 = make([]bool, len(wordList))
	visited2 = make([]bool, len(wordList))

	queue1.Add(beginWord)
	queue2.Add(endWord)
	visited1[idx1] = true
	visited2[idx2] = true

	for !queue1.IsEmpty() && !queue2.IsEmpty() {
		count++

		if queue1.Size() > queue2.Size() {
			queue1, queue2 = queue2, queue1
			visited1, visited2 = visited2, visited1
		}

		size := queue1.Size()

		for size > 0 {
			size--
			curEle := queue1.Poll()
			for i, v := range wordList {
				if visited1[i] {
					continue
				}

				if !conversable(curEle.(string), v) {
					continue
				}
				fmt.Println(i)
				if visited2[i] {
					return count + 1
				}

				queue1.Add(wordList[i])
				visited1[i] = true
			}
		}
	}
	return 0
}

/*
	"hit"
	"cog"
	["hot","dot","dog","lot","log","cog"]
*/

func main() {
	beginWord := "a"
	endWord := "c"
	wordList := []string{"a", "b", "c"}
	ans := ladderLengthPro(beginWord, endWord, wordList)

	fmt.Println(ans)
}
