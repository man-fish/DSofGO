package main

import (
	"fmt"
	"go_DataStruct/util"
	"math"
)

func mergeStoneRecursive(sum []int, i, j int) int {
	if i == j {
		return 0
	}

	if j == i + 1 {
		return sum[j+1] - sum[i]
	}

	bestChoice := math.MaxInt64

	for k := i+1; k < j ; k++ {
		l := mergeStoneRecursive(sum,i,k-1) + mergeStoneRecursive(sum,k,j) + sum[j+1] - sum[i]
		r := mergeStoneRecursive(sum,i,k) + mergeStoneRecursive(sum,k+1,j) + sum[j+1] - sum[i]
		cbs := int(math.Min(float64(l),float64(r)))
		// 就以2，3，1为例，先合并2和3与先合并3和1的结果是不同的我们要取最小值。
		bestChoice = int(math.Min(float64(cbs),float64(bestChoice)))
	}

	if bestChoice == math.MaxInt64 {
		return 0
	}else{
		return bestChoice
	}
}

func solveProblem() {
	stones := []int{2,3,1}
	sum := make([]int, len(stones)+1)

	for idx := 1; idx < len(sum); idx++ {
		for i := 1; i <= idx ; i++  {
			sum[idx] += stones[i-1]
		}
	}

	memo := make([][]int, len(stones))

	for idx := range memo{
		memo[idx] = make([]int, len(stones))
	}
	fmt.Println(mergeStoneRecursiveMemorized(sum,memo,0,len(stones)-1))
	util.Print2DArray(memo)
}

func mergeStoneRecursiveMemorized(sum []int, memory [][]int, i, j int) int {
	if i == j {
		return 0
	}

	if memory[i][j] != 0 {
		return memory[i][j]
	}

	if j == i + 1 {
		memory[i][j] = sum[j+1] - sum[i]
		return memory[i][j]
	}

	bestChoice := math.MaxInt64

	for k := i+1; k < j ; k++ {
		l := mergeStoneRecursiveMemorized(sum,memory,i,k-1) + mergeStoneRecursiveMemorized(sum,memory,k,j) + sum[j+1] - sum[i]
		r := mergeStoneRecursiveMemorized(sum,memory,i,k) + mergeStoneRecursiveMemorized(sum,memory,k+1,j) + sum[j+1] - sum[i]
		cbs := int(math.Min(float64(l),float64(r)))
		bestChoice = int(math.Min(float64(cbs),float64(bestChoice)))
	}

	if bestChoice == math.MaxInt64 {
		return 0
	}else{
		memory[i][j] = bestChoice
		return bestChoice
	}
}

func mergeStone(n int,w []int) int   {
	dp := make([][]int,n+1)
	for  i := 0; i < n+1 ; i ++  {
		dp[i] = make([]int,n+1)
	}
	/*声明dp数组*/
	sum := make([]int, n+1)
	for idx := 1; idx < len(sum); idx++ {
		for i := 1; i <= idx ; i++  {
			sum[idx] += w[i-1]
		}
	}
	/*计算组合石头花费（从0开始，0表示没有石头）*/
	for len := 2; len < n+1; len++  {		/* 区间长度，在这道题里指的就是第1个到第n个。 */
		for i := 1; i < n+1 ; i++  {		/* 枚举起点，从1开始到n，i表示起点下标，j表示终点下标， */
			j := i + len -1					/* 计算终点坐标 */
			if j > n {						/* 省去越界的部分 */
				continue
			}
			for k := i; k < j ; k++  {		/* 枚举区间分割点 */
				if dp[i][j] == 0 {			/* 这里我们要避免前值为0 */
					dp[i][j] = dp[i][k]+dp[k+1][j]+sum[j]-sum[i-1]
				}else{
					dp[i][j] = int(math.Min(float64(dp[i][j]),float64(dp[i][k]+dp[k+1][j]+sum[j]-sum[i-1])))
				}

			}
		}
	}
	util.Print2DArray(dp)
	return dp[1][n]
}

func mergeStonev2(n int,w []int) int   {
	dp := make([][]int,n)
	for  i := 0; i < n ; i ++  {
		dp[i] = make([]int,n)
	}
	/*声明dp数组*/
	sum := make([]int, n+1)
	for idx := 1; idx < len(sum); idx++ {
		for i := 1; i <= idx ; i++  {
			sum[idx] += w[i-1]
		}
	}
	/*计算组合石头花费（从0开始，0表示没有石头）*/
	for len := 2; len <= n; len++  {		/* 区间长度，在这道题里指的就是第1个到第n个。 */
		for i := 0; i < n ; i++  {			/* 枚举起点，从1开始到n，i表示起点下标，j表示终点下标， */
			j := i + len -1					/* 计算终点坐标 */
			if j >= n {						/* 省去越界的部分 */
				continue
			}
			for k := i; k < j ; k++  {		/* 枚举区间分割点 */
				if dp[i][j] == 0 {			/* 这里我们要避免前值为0 */
					dp[i][j] = dp[i][k]+dp[k+1][j]+sum[j+1]-sum[i]
				}else{
					dp[i][j] = int(math.Min(float64(dp[i][j]),float64(dp[i][k]+dp[k+1][j]+sum[j+1]-sum[i])))
				}

			}
		}
	}
	util.Print2DArray(dp)
	return dp[1][n-1]
}

func main() {
	stones := []int{2,3,1}
	sum := make([]int, len(stones)+1)

	for idx := 1; idx < len(sum); idx++ {
		for i := 1; i <= idx ; i++  {
			sum[idx] += stones[i-1]
		}
	}

	fmt.Println(sum)

	fmt.Println(mergeStoneRecursive(sum,0,len(stones)-1))
	solveProblem()
	mergeStone(len(stones),stones)
	mergeStonev2(len(stones),stones)
}