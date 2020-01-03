package main

import (
	"fmt"
	"go_DataStruct/ALG/sort"
)
/*
	贪心算法
		https://blog.csdn.net/effective_coder/article/details/8736718
*/

/*
	问题：
		假设你是一个很棒的父母，并希望给你的孩子一些饼干。但是，你最多应该给每个孩子一个饼干。
		每个孩子都有一个贪婪指数g(i)，这是指一个饼干的最小大小。
		孩子将满足：
			每个 cookie j 都有一个大小 s(j)。如果 sj >= gi，我们可以将 cookie j 分配给孩子 i，孩子就会很开心。
		你的目标是求出最多开心孩子的个数。
	注意：
		你可以假设胃口值为正。
		一个小朋友最多只能拥有一块饼干。
	例子：
		输入: [1,2,3], [1,1]

		输出: 1

		解释:
		你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
		虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
		所以你应该输出1。
		示例 2:

		输入: [1,2], [1,2,3]

		输出: 2

		解释:
		你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
		你拥有的饼干数量和尺寸都足以让所有孩子满足。
		所以你应该输出2.
*/

func findContentChildren(g []int, s []int) int {
	sort.BubbleSort(g)
	sort.BubbleSort(s)

	var si, gi,res int
	for si < len(s) && gi < len(g) {
		if s[si] >= g[gi] {
			res++
			si++
			gi++
		}else{
			gi++
		}
	}
	return res
}

func main(){
	fmt.Println(findContentChildren([]int{1,2},[]int{3}))
}