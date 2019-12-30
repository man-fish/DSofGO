package main

import "math"
/*https://juejin.im/post/5d556b7ef265da03aa2568d5#heading-6*/
func coinCharge(coins []int, account int) int {
	if account == 0 {
		return 0
	}

	ans := math.MaxInt64

	for _, coin := range coins {
		if account - coin < 0 {continue}
		subPro := coinCharge(coins,account-coin)
		if subPro == -1 {continue}
		ans = int(math.Min(float64(ans),float64(subPro)+1))
	}

	if ans == math.MaxInt64 {
		return -1
	}else{
		return ans
	}
}

func coinMemorizeCharge(coins []int, memory []int, account int) int {
	if memory[account] != -2 {
		return memory[account]
	}
	if account == 0 {
		return 0
	}
	ans := math.MaxInt64
	for _,coin := range coins {
		if account - coin < 0 {
			continue
		}
		pre := coinMemorizeCharge(coins,memory,account-coin)
		if pre == -1 {
			continue
		}
		ans = int(math.Max(float64(pre)+1,float64(ans)))
	}

	if ans != math.MaxInt64 {
		memory[account] = ans
		return ans
	}else{
		return -1
	}
}

func main() {
}
