/*
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。
	示例 1：
		输入： 2
		输出： 2
		解释： 有两种方法可以爬到楼顶。
		1.  1 阶 + 1 阶
		2.  2 阶
	解题思路：
		状态转移方程 -> 暴力递归 -> 超出时间限制 -> 备忘录数组优化 -> 动态规划写法 -> dp数组优化(n->3->2) ->双100%
		f(n) = f(n-1) + f(n-2) | if (n <= 2) { return n }
*/

package main

import "fmt"

func climbStairs(n int) int {
	var dparr = make([]int, n+1)
	methods := climb(n, dparr)
	fmt.Println(dparr)
	return methods
}

func climb(n int, dparr []int) int {
	if dparr[n] != 0 {
		return dparr[n]
	}

	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	methods := climb(n-2, dparr)
	methods += climb(n-1, dparr)
	dparr[n] = methods
	return methods
}

func climbStairsDy(n int) int {
	if n <= 2 {
		return n
	}
	dpArr := make([]int, n+1)
	dpArr[1] = 1
	dpArr[2] = 2
	for i := 3; i <= n; i++ {
		dpArr[i] = dpArr[i-1] + dpArr[i-2]
	}

	return dpArr[n]
}

func climbStairsDyPro(n int) int {
	if n <= 2 {
		return n
	}
	dpArr := make([]int, 3)
	dpArr[1] = 1
	dpArr[2] = 2
	for i := 3; i <= n; i++ {
		curMs := dpArr[1] + dpArr[2]
		dpArr[0] = dpArr[1]
		dpArr[1] = dpArr[2]
		dpArr[2] = curMs
	}
	fmt.Println(dpArr)
	return dpArr[2]
}

func climbStairsDyProPro(n int) int {
	if n <= 2 {
		return n
	}
	dpArr := make([]int, 2)
	dpArr[0] = 1
	dpArr[1] = 2
	for i := 3; i <= n; i++ {
		dpArr[1], dpArr[0] = dpArr[0]+dpArr[1], dpArr[1]
	}
	return dpArr[1]
}

func main() {
	fmt.Println(climbStairs(44))
	fmt.Println(climbStairsDy(1))
	fmt.Println(climbStairsDyPro(44))
	fmt.Println(climbStairsDyProPro(44))
}
