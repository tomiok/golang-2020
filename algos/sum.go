package main

import "fmt"

func main() {
	//nums := []int{1,2,4,4} //true case
	nums := []int{1,2,4,5} //false case
	res := 8

	val := findSum(nums, res)

	fmt.Println(val)
}

func findSum(nums []int, res int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[res-v] = 1
	}
	return false
}
