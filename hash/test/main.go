package main

import (
	"fmt"
)

var hashset map[interface{}]interface{}

func remRep(arr []interface{}) []interface{} {
	uniques := make(map[interface{}]int)
	for i, v := range arr {
		uniques[v] = i
	}
	arr = make([]interface{}, len(arr))
	for k, v := range uniques {
		arr[v] = k
	}
	uniqueArr := make([]interface{}, len(uniques))
	i := 0
	for _, v := range arr {
		if v == nil {
			continue
		}
		uniqueArr[i] = v
		i++
	}
	return uniqueArr
}

func main() {
	hashset = make(map[interface{}]interface{})
	hashset["hello"] = nil
	if _, ok := hashset["hello"]; ok {
		fmt.Println(hashset["hello"])
	}

	arr := []interface{}{1, 2, 3, 4, 4}
	arr = remRep(arr)
	fmt.Println(arr)
}
