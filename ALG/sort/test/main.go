package main

import (
	"fmt"
	"go_DataStruct/ALG/sort"
)

func main() {
	arr := sort.RandomArray(5, 100)
	fmt.Println(arr)
	sort.InsertSort(arr)
	fmt.Println(arr)
	arr = sort.RandomArray(5, 100)
	fmt.Println(arr)
	sort.ShellSort(arr)
	fmt.Println(arr)
	arr = sort.RandomArray(5, 100)
	fmt.Println(arr)
	sort.BubbleSort(arr)
	fmt.Println(arr)
	arr = sort.RandomArray(5, 100)
	fmt.Println(arr)
	sort.SelectSort(arr)
	fmt.Println(arr)
	arr = sort.RandomArray(5, 100)
	fmt.Println(arr)
	sort.QuickSort(arr)
	fmt.Println(arr)
	arr = sort.RandomArray(5, 100)
	fmt.Println(arr)
	temp := make([]int, len(arr))
	sort.MergeSort(arr, temp, 0, len(arr)-1)
	fmt.Println(arr)
}
