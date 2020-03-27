package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	//10110
	//00010
	idxA := len(a) - 1
	idxB := len(b) - 1
	if idxA > idxB {
		for idxB != idxA {
			idxB++
			b = "0" + b
		}
	} else if idxA < idxB {
		for idxB != idxA {
			idxA++
			a = "0" + a
		}
	}
	res := ""
	lst := 0
	for idxA >= 0 {
		iA, _ := strconv.Atoi(a[idxA : idxA+1]) // 0 0 1
		iB, _ := strconv.Atoi(b[idxA : idxA+1]) // 1 0 1
		sub := iA ^ iB                          // 1 0 0
		fmt.Printf("第%v位：iA:%v, iB:%v, sub1:%v， ", len(a)-idxA, iA, iB, sub)
		sub = sub ^ lst // 1 1 1
		fmt.Printf("sub2:%v, lst:%v\n", sub, lst)
		idxA--
		if sub == 0 {
			res = "0" + res

			if iA == 1 || iB == 1 || lst == 1 {
				lst = 1
				continue
			}
			lst = 0
		} else {
			res = "1" + res
			if iA == 1 && iB == 1 && lst == 1 {
				lst = 1
				continue
			}
			lst = 0
		}
	}
	if lst == 1 {
		res = "1" + res
	}
	return res
}

func main() {
	sub := addBinary("10110", "10")
	fmt.Println(sub)
}
