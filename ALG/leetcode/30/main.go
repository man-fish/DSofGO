package main

import "fmt"

func findSubstring(s string, words []string) []int {
	start := make([]int, 0)
	wordsLen := len(words)
	if s == "" || wordsLen == 0 {
		return start
	}
	wordLen := len(words[0])
	strlen := len(s)
	for i := 0; i < strlen && wordsLen <= (strlen-i)/wordLen; i++ {
		visited := make([]int, wordsLen)
		for j := i; j <= strlen-wordLen; j = j + wordLen {
			matchedIdx := match(s[j:j+wordLen], words, visited)

			if matchedIdx == -1 || visited[matchedIdx] != 0 {
				break
			}
			visited[matchedIdx] = 1
			if isAllMatched(visited) {
				start = append(start, i)
				i = i + wordLen
				break
			}
		}
	}
	return start
}

func match(w string, words []string, visited []int) int {
	for i, v := range words {
		if visited[i] == 1 {
			continue
		}
		if v == w {
			return i
		}
	}
	return -1
}

func isAllMatched(visited []int) bool {
	for _, v := range visited {
		if v == 0 {
			return false
		}
	}
	return true
}

//   s = ,
//  words = ["foo","bar"]
// "wordgoodgoodgoodbestword"
// ["word","good","best","good"]

func main() {
	// "foobarfoobar"
	// ["foo","bar"]
	s := "aaa"
	words := []string{"aa", "aa"}
	res := findSubstring(s, words)
	fmt.Println(res)
}
