package main

import "fmt"

type fn func(i int) int // see function fnTypes

func main() {
	strValue := fnArgs(11, func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(strValue)
}

func fnArgs(n int, evenFn func(i int) bool) string {

	if evenFn(n) {
		return "yes"
	} else {
		return "no"
	}
}

func fnTypes() {
	var f fn

	f = func(i int) int {
		return 0
	}

	f(0)
}

func anonFunctions() {
	fn := func(i int) int {
		return i + 10
	}

	fg := func(s string) int {
		return 0
	}

	fn(8)
	fg("100")
}

// reserved word 'func'
// name
// arguments -- optional
// return type -- optional in case of return nothing (void return)
// multiple return type (int, string, error)
// named return type (use it with naked return)
func f1() {
	fmt.Println("no args and no return")
}

func f2(i int) int {
	return i + 10
}

func f3(i, j int) (int, int) {
	return i + j, i * j
}

//naked return
func f4(i, j int) (sum int, multi int) {
	sum = i + j
	multi = i * j
	return //naked return
	//return sum, multi named return (valid)
}

// invariants could receive 0 to n args
func invariantsArgs(i ...int) (sum int) {
	for _, v := range i {
		sum += v
	}
	return
}
