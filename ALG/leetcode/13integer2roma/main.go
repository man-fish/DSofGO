package main

import "fmt"

func romanToInt(s string) int {
	romaMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	sum := 0
	i := 0
	for ; i < len(s)-2; i++ {
		if v, ok := romaMap[s[i:i+2]]; ok {
			sum += v
			i++
		} else {
			sum += romaMap[s[i:i+1]]
		}

	}
	if i == len(s)-1 {
		return sum + romaMap[s[i:]]
	}
	if v, ok := romaMap[s[i:]]; ok {
		sum += v
	} else {
		sum += romaMap[s[i:i+1]]
		sum += romaMap[s[i+1:]]
	}

	return sum
}

func main() {
	result := romanToInt("MDCXCV")
	fmt.Println(result)
}
