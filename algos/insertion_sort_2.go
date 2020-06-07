package main

import "fmt"

// rearrange all the array with insertion sort
func main() {
	arr := []int{2, 4, 5, 7, 3, 8}
	fmt.Println(insertion(arr))
}

func insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
