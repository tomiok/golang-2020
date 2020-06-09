package main

import "fmt"

func main() {
	steps := 4
	res := stairs(int32(steps))
	fmt.Println(res)
}

func stairs(steps int32) int32 {
	counter := make([]int32, steps)
	// jumps
	counter[0] = 1
	counter[1] = 2
	counter[2] = 4
	for j := 3; j < len(counter); j++ {
		counter[j] += counter[j-1] + counter[j-2] + counter[j-3]
	}

	return counter[steps-1]
}
