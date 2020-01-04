package main

import (
	"fmt"
	"math"
)

/*
	钢条切割：https://www.jianshu.com/p/32798a535f5f || https://blog.csdn.net/u013309870/article/details/75193592
	给定一段长度为n英寸的钢条和一个价格表pi(i=1，2，...，n)，求解切割钢条方案，使得销售收益rn最大。
	注意，如果长度为n英寸的钢条价格pn足够大，最优解可能就是完全不需要切割。
*/

func CutSteel(value []int, length int) int {
	if length == 0 {
		return 0
	}

	bestValue := math.MinInt64

	for i := 1; i <= length; i++ {
		preValue := value[i-1]+CutSteel(value,length-i)
		bestValue = int(math.Max(float64(preValue),float64(bestValue)))
	}
	return bestValue
}


func MemorizedCutSteel(v []int,r []int,n int) int {
	if r[n] != -1 {
		return r[n]
	}
	bestChoice := -1
	if n == 0 {
		bestChoice = 0
	}else{
		for i := 1; i < n; i++ {
			temp := v[i] + MemorizedCutSteel(v,r,n-i)
			bestChoice = int(math.Max(float64(bestChoice), float64(temp)))
		}
	}
	r[n] = bestChoice
	fmt.Println(r)
	return bestChoice
}

/*

	@v []int	价值数组
	@n int		钢条长度
*/
func BottomUpCutSteel(v []int, n int) int {
	r := make([]int, n+1)
	/*
		创建备忘录，用于记录长度为i的时候的最优解。
		我们只要拿到了上一轮的最优解，下一轮只需要比较之前最优解的组合外加上当前长度不切割的价值就可以了，依此递推，
	*/
	for i := 1; i <= n; i++  {
		bestChoice := -1
		for j := 1; j <= i; j++ {
			bestChoice = int(math.Max(float64(bestChoice), float64(v[j-1] + r[i-j])))
		}
		r[i] = bestChoice
	}
	return r[len(r)-1]
}


func main() {
	fmt.Println(CutSteel([]int{1,5,8,9},4))
	fmt.Println(BottomUpCutSteel([]int{1,5,8,9},4))
}