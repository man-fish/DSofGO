package main

import "fmt"

func demo(n int){
	if n > 2 {
		n--
		demo(n)
	}
	fmt.Println(n)
}

func demo2(n int){
	if n > 2 {
		n--
		demo2(n)
	}else{
		fmt.Println(n)
	}
}


func main() {
	demo(4)
	demo2(4)
}