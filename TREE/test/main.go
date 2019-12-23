package main

import (
	"fmt"
	"go_DataStruct/TREE"
)

//func reverseString(s []byte) {
//	s[0],s[len(s)-1] = s[len(s)-1],s[0]
//	if len(s) > 1 {
//		reverseString(s[1:len(s)-1])
//	}else{
//		fmt.Println(s)
//	}
//}

func reverseString(s []byte)  {
	a := s[0]
	if len(s) > 1 {
		reverseString(s[1:len(s)])
	}

	fmt.Print(string(a))
}

var i int = 0
func reverseStringvol2(s []byte)  {

	s[i],s[len(s)-i-1] = s[len(s)-i-1],s[i]
	if i == len(s)/2 {
		fmt.Printf("%s",s)
	}else{
		i++
		reverseStringvol2(s)
	}
}

//
//func reverse (s []byte) []byte {
//	//[]byte{s[len(s)-1]} + reverse(s[1:])
//	return []byte(string(s[len(s)-1]reverse(s[:len(s)-1])))
//}


func main() {
	//fmt.Println(string(rune('A')+2))
	//reverseString([]byte{'H','a','n','n','a','h'})

	haff := TREE.NewHaffTree([]int{7,5,1,2})
	fmt.Println(haff.GetCode(0))
	fmt.Println(haff)
	fmt.Println(haff.Decode(haff.Encode("ACA")))
}
