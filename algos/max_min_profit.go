package main

import (
	"fmt"
)

func main() {
	values := []int{90, 100, 80, 90, 25}
	res := maxMin(values)
	fmt.Println(res)
}

func maxMin(values []int) string {
	buy := 0
	sell := false

	var profit int
	for _, price := range values {
		if buy == 0 {
			buy = price
			continue
		}

		if price > buy {
			sell = true
		}

		if sell {
			profit += price - buy
			sell = false
			buy = 0
		}
	}

	return fmt.Sprintf("the profit is %d", profit)
}
