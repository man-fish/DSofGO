package main

import (
	"fmt"
)

type rArr []rune

//Len()
func (s rArr) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s rArr) Less(i, j int) bool {
	return s[i] < s[j]
}

//Swap()
func (s rArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func QuickSort(arr []rune) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []rune, left int, right int) {
	if len(arr) <= 0 {
		return
	}
	iLeft := left
	iRight := right
	pivot := arr[(left+right)/2]
	for iLeft < iRight {
		for arr[iLeft] < pivot {
			iLeft++
		}

		for arr[iRight] > pivot {
			iRight--
		}

		if iLeft >= iRight {
			break
		}

		arr[iRight], arr[iLeft] = arr[iLeft], arr[iRight]

		if arr[iLeft] == pivot {
			iRight--
		}
		if arr[iRight] == pivot {
			iLeft++
		}
	}

	if iLeft == iRight {
		iLeft++
		iRight--
	}

	if iLeft < right {
		quickSort(arr, iLeft, right)
	}

	if left < iRight {
		quickSort(arr, left, iRight)
	}
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	if len(strs) == 0 {
		return res
	}
	groups := make(map[string][]string)

	for _, v := range strs {
		arr := []rune(v)
		// sort.Sort(arr)
		QuickSort(arr)
		fmt.Println(arr)
		key := string(arr)
		if _, ok := groups[key]; ok {
			groups[key] = append(groups[key], v)
		} else {
			groups[key] = append(make([]string, 0), v)
		}
	}
	for _, value := range groups {
		res = append(res, value)
	}
	return res
}

func main() {
	res := groupAnagrams([]string{"abc", "bac", "efg", "ijk"})
	fmt.Println(res)
}
