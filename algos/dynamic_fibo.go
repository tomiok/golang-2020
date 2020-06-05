package main

import "fmt"

func main() {
	fmt.Println(dFib(12))
}

func dFib(n int) int {
	var (
		n1  int
		n2  int
		sum int
	)
	for i := 0; n >= i; i++ {
		if i == 0 {
			continue
		}
		if i == 1 {
			n1 = 1
			continue
		}
		sum = n1 + n2
		n2 = n1
		n1 = sum
	}
	return sum
}
