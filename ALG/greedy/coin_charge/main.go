package main

import (
	"fmt"
	"math"
)

/*
	钱币找零问题https://blog.csdn.net/qq_32400847/article/details/51336300
	假设1元、2元、5元、10元、20元、50元、100元的纸币分别有c0, c1, c2, c3, c4, c5, c6张。
	现在要用这些钱来支付K元，至少要用多少张纸币？用贪心算法的思想，很显然，每一步尽可能用面值大的纸币即可。
	在日常生活中我们自然而然也是这么做的。在程序中已经事先将Value按照从小到大的顺序排好。

	贪心算法的思路就是每次我们都取最大值，用相对较小的值来填充。
*/

func coinCharge(coins []int, count []int, n int) int {
	if n == 0 {
		return 0
	}
	num := 0
	for i := len(coins)-1 ; i >=  0  ; i-- {
		c := int(math.Min(float64(n/coins[i]),float64(count[i])))
		/* 最大纸币是否充足 */
		n -= c*coins[i]
		/* 削减金额 */
		num += c
		/* 计算消耗纸币次数 */
		if n == 0 {
			break
		}
	}
	if n != 0 {
		return -1
	}
	return num
}

func main () {
	value := []int{1,5,10,20,50,100}
	count := []int{3,0,2,1,3,5}
	fmt.Println(coinCharge(value,count,78))
}
