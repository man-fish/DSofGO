package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lenS := len(s)
	dp := make([]int, lenS+1)
	m := make(map[uint8]int, 0)
	max := 0
	for idx := 0; idx < lenS; idx++ {
		char := s[idx]
		dpIdx := idx + 1
		if i, ok := m[char]; ok {
			dp[dpIdx] = idx - i
			m = make(map[uint8]int, 0)
			for i < idx {
				m[s[i]] = i
				i++
			}
		} else {
			dp[dpIdx] = dp[idx] + 1
		}
		m[char] = idx
		if dp[dpIdx] > max {
			max = dp[dpIdx]
		}
	}
	return max
}

func lengthOfLongestSubstringV2(s string) int {
	m := make(map[uint8]int, 0)

	lst := 0
	max := 0

	for idx := 0; idx < len(s); idx++ {
		char := s[idx]
		if i, ok := m[char]; ok {
			lst = idx - i
			m = make(map[uint8]int, 0)
			for i < idx {
				m[s[i]] = i
				i++
			}
		} else {
			lst = lst + 1
		}
		m[char] = idx
		if lst > max {
			max = lst
		}
	}
	return max
}

func lengthOfLongestSubstringV3(s string) int {
	var (
		alps [128]int	// 记录数组
		left int 		// 左界
		max  int		// 记录最大值
	)
	//for i := range alps {
	//	alps[i] = -1
	//}
	for idx := 0; idx < len(s); idx++ {
		charId := s[idx]	        // 拿到当前字符 id

		if i := alps[charId]; i >= left || (i > left && i != 0) {
			left = i + 1			// 左界更新
		}
		alps[charId] = idx       	// 记录字符出现的位置

		if idx + 1 - left > max {
			max =  idx + 1 - left
		}
	}
	return max
}


func main() {
	ml := lengthOfLongestSubstring("")
	ml = lengthOfLongestSubstringV3("abca")
	fmt.Println(ml)
}
