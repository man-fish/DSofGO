package main

import "fmt"

const kind = 6
const weight = 21

func SolveQuestion(kind, weight int) {
	var knappsack01Map = make([][]int, kind)
	for idx := range knappsack01Map {
		knappsack01Map[idx] = make([]int, weight)
	}
	w := []int{0,2,3,4,5,9}
	v := []int{0,3,4,5,8,10}
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
}