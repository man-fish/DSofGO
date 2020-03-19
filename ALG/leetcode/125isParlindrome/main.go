/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
	输入: "A man, a plan, a canal: Panama"
	输出: true
	示例 2:

	输入: "race a car"
	输出: false
*/
package main

import "fmt"

// a c d e f
//     2
// a c d e f h
//     2

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
	var flg bool
	flg = isPalindromePro("A man, a plan, a canal: Panama")
	fmt.Println(flg)

	flg = isPalindromePro("0P")
	fmt.Println(flg)
}
