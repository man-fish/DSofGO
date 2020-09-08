package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var (
		i   int
		res string
	)

	if len(strs) != 0 {
		return res
	}

	for {
		if i > len(strs[0])-1 {
			return res
		}
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != strs[0][i] {
				return res
			}
		}
		res = res + string(strs[0][i])
		i++
	}
}

func main() {
	strs := []string{}
	fmt.Println(longestCommonPrefix(strs))
}
