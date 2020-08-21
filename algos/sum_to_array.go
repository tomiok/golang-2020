package main

import "fmt"

// you have an array [2,4,5] and you need to sum 1 to that, return the new value
// as an array like [2,4,6]

func main() {
	fmt.Println(sum([]int{9, 9, 8}))
}

func sum(arr []int) []int {
	l := len(arr)
	res := make([]int, l)
	carry := 1
	for i := len(arr) - 1; i >= 0; i-- {
		value := arr[i] + carry
		res[i] = value % 10
		carry = value / 10
	}

	if carry == 1 {
		res2 := []int{}
		for i, v := range res {
			if i == 0 {
				res2 = append(res2, 1)
				res2 = append(res2, v)
			} else {
				res2 = append(res2, v)
			}
		}
		return res2
	}

	return res
}
