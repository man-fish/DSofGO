package main

import (
	"fmt"
	"math"
)

/*
//mst(dp,0) 初始化DP数组
for(int i=1;i<=n;i++)
{
    dp[i][i]=初始值
}
for(int len=2;len<=n;len++)  //区间长度
for(int i=1;i<=n;i++)        //枚举起点
{
    int j=i+len-1;           //区间终点
    if(j>n) break;           //越界结束
    for(int k=i;k<j;k++)     //枚举分割点，构造状态转移方程
    {
        dp[i][j]=max(dp[i][j],dp[i][k]+dp[k+1][j]+w[i][j]);
    }
}
*/

func mergeStone(n int,w []int) int {
	dp := make([][]int,n)
	for  i := 0; i < n ; i ++  {
		dp[i] = make([]int,n)
	}

	sum := make([]int,n)

	for idx := range sum{
		for i := 0; i <= idx ; i++  {
			sum[idx] += w[i]
		}
	}

	for len := 2; len < n; len++  {		/* 区间长度，在这道题里指的就是第1个到第n个。 */
		for i := 1; i < n ; i++  {		/* 枚举起点，从1开始到n，i表示起点下标，j表示终点下标， */
			j := i + len -1
			if j > n {
				continue
			}
			for k := i; k < j ; k++  {
				dp[i][j] = int(math.Min(float64(dp[i][j]),float64(dp[i][k]+dp[k+1][j]+sum[j]-sum[i-1])))
			}
		}
	}
	fmt.Println(dp[1][n])
	return dp[1][n]
}

func main() {
	stones := []int{2,4,6,78,1}
	mergeStone(len(stones),stones)
}