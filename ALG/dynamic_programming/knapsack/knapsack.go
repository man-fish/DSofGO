package main

import (
	"fmt"
	"math"
)

/*
	明天做背包问题的递归方法，切木棒还有这个https://juejin.im/post/5d556b7ef265da03aa2568d5#heading-6。
	解决游历问题https://blog.csdn.net/wqk1014/article/details/8283325。
*/

const kind = 6
const weight = 21

func Knapsack(w,v []int,weight int,index int) int {
	if index == -1 || weight <= 0 {
		//如果说背包不能装下最轻的物品，当前物品种类也为0了，返回的价值就是0。
		return 0
	}

	bestChoice := math.MinInt64

	if weight < w[index] {
		/*如果说背包装不下当前物品,就返回公式为B(k-1,n)*/
		bestChoice = Knapsack(v,w,weight,index-1)
	}else{
		valueIn := v[index] + Knapsack(v,w,weight-w[index],index-1)
		valueOut := Knapsack(v,w,weight,index-1)
		bestChoice = int(math.Max(float64(valueIn),float64(valueOut)))
	}

	return bestChoice
}

func KnapsackMemorized(w,v []int,weight int,ans []int,index int) int {
	if index == -1 || weight <= 0 {
		//如果说背包不能装下最轻的物品，当前物品种类也为0了，返回的价值就是0。
		return 0
	}

	bestChoice := math.MinInt64

	if ans[weight] != -1 {
		return ans[weight]
	}

	if weight < w[index] {
		/*如果说背包装不下当前物品,就返回公式为B(k-1,n)*/
		bestChoice = Knapsack(v,w,weight,index-1)
	}else{
		valueIn := v[index] + Knapsack(v,w,weight-w[index],index-1)
		valueOut := Knapsack(v,w,weight,index-1)
		bestChoice = int(math.Max(float64(valueIn),float64(valueOut)))
	}
	ans[weight] = bestChoice
	return bestChoice
}

func SolveQuestion(kind, weight int) {
	var knappsack01Map = make([][]int, kind)
	for idx := range knappsack01Map {
		knappsack01Map[idx] = make([]int, weight)
	}
	w := []int{0,2,3,4,5,9}
	v := []int{0,3,4,5,8,10}

	fmt.Println(Knapsack(w,v,20,len(v)-1))
	SolveKnapsack01(knappsack01Map,v,w,weight)
}

func SolveKnapsack01 (knapsackValueMap [][]int, V,W []int, weigt int){
	for k := 1;k < kind ; k++  {
		for w := 1; w < weight  ; w++ {
			if W[k] > w {
				/* 如果说第k种物品重量大于当前背包剩余容量 */
				knapsackValueMap[k][w] = knapsackValueMap[k-1][w]
			}else{
				valueIn := knapsackValueMap[k-1][w-W[k]]+V[k]
				valueOut := knapsackValueMap[k-1][w]
				if  valueIn > valueOut  {
					/*如果说不装当前物品的价值比装当前物品的价值高*/
					knapsackValueMap[k][w] = valueIn
				}else{
					knapsackValueMap[k][w] = valueOut
				}
			}
		}
	}
	for idx := range knapsackValueMap {
		for _,item := range knapsackValueMap[idx] {
			fmt.Print(item,"\t")
		}
		fmt.Println()
	}
}

func main() {
	SolveQuestion(kind,weight)
	w := []int{2,3,4,5,9}
	v := []int{3,4,5,8,10}
	arr := make([]int, weight)
	for a := range arr {
		arr[a] = -1
	}

	fmt.Println(Knapsack(v,w,20,len(v)-1))
	fmt.Println(KnapsackMemorized(v,w,20,arr,len(v)-1))
}