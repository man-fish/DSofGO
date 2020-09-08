package main

func multiply(num1 string, num2 string) string {
	n1, n2 := 0
	c1 := 1
	for i := len(num1) - 1; i >= 0; i-- {
		n1 += n1 + int(num1[i])*c1

	}
}

func main() {

}
