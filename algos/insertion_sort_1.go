package main

import (
	"fmt"
	"strings"
)

// insertion sort, only the last element
func main() {
	n := 5
	arr := []int32{2, 4, 6, 8, 1}
	insertionSort1(int32(n), arr)
}

func insertionSort1(n int32, arr []int32) {
	num := arr[len(arr)-1]
	n = n - 2
	for ; n >= 0; n-- {
		printArr(arr, n, num)
	}
}

// TODO remove last character
func printArr(arr []int32, index, n int32) {
	s := ""
	flag := false
	for i, v := range arr {
		if int32(i) == index {
			if v > n {
				s = s + fmt.Sprint(v) + fmt.Sprint(v)
			} else {
				s = s + fmt.Sprint(v) + fmt.Sprint(n)
				flag = true
			}
		} else {
			s = s + fmt.Sprint(v)
		}
	}
	if !flag && index == 0 {
		fmt.Println(s)
		f := s[0]
		fmt.Println(strings.Replace(s, string(f), fmt.Sprint(n), 1))
	} else {
		fmt.Println(s)
	}

}
