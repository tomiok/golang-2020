package main

import "fmt"

// you have some infinite coins of certain value, how many ways do you have to reach an specific number
func main() {
	fmt.Println(change(10, []int64{2, 5, 3, 6}))
}

func change(amount int32, coins []int64) int64 {
	combinations := make([]int64, amount+1)
	combinations[0] = 1
	for _, coin := range coins {
		for j := 1; j < len(combinations); j++ {
			if int64(j) >= coin {
				combinations[j] += combinations[int64(j)-coin]
			}
		}
	}

	return combinations[amount]
}
