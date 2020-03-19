/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串，返回 s 所有可能的分割方案。
示例:
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
*/

package main

import (
	"fmt"
)

func partition(s string) [][]string {
	res := make([][]string, 0)
	segs := make([]int, 1)
	handler(s, segs, &res)
	return res
}

func handler(s string, segs []int, res *[][]string) {
	lastSeg := segs[len(segs)-1]
	if lastSeg == len(s) {
		tmp := make([]string, 0)
		for i := 1; i < len(segs); i++ {
			start := segs[i-1]
			end := segs[i]
			tmp = append(tmp, s[start:end])
		}
		*res = append(*res, tmp)
	}

	for i := lastSeg + 1; i <= len(s); i++ {
		if isPalindromePro(s[lastSeg:i]) {
			tmp := append(segs, i)

			handler(s, tmp, res)
		}
	}
}

func isPalindromePro(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end { //当start >= end 有两种情况， 两指针重合于奇数中点或者两指针错位，这个时候都应该返回。
		if !((s[start] >= 97 && s[start] <= 122) || (s[start] >= 65 && s[start] <= 90) || (s[start] >= 48 && s[start] <= 57)) {
			start++ //过滤
			continue
		}
		if !((s[end] >= 97 && s[end] <= 122) || (s[end] >= 65 && s[end] <= 90) || (s[end] >= 48 && s[end] <= 57)) {
			end-- //过滤
			continue
		}
		if s[start] != s[end] { //判断是否相等
			if s[start] >= 65 && s[end] >= 65 {
				if (s[start]+32) != s[end] && (s[start]-32) != s[end] {
					return false
				}
				// else {
				// 	fmt.Println("判断是否为大小写")
				// }
			} else {
				return false
			}
		}
		start++
		end--
	}
	return true
}

func main() {
	res := partition("efe")
	fmt.Println(res)
}
