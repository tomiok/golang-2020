package main

import "fmt"

func main() {
	// this is wrong, wont compile
	//a := withReturnValues()

	// correct
	valInt, valString := withReturnValues()

	fmt.Printf("%d - %s", valInt, valString)
}

func callUnExported() {
	unexported(1, "asd", func(u int) {
		fmt.Println(u)
	}, "hello", "world")

	fmt.Println()
}

// unexerported eny type could work as an argument, even functions
// invariant args are 0...n the syntax is [name ...{type}]
func unexported(i int, s string, fn func(u int), t ...string) {
	fmt.Println(i)
	fmt.Println(s)

	fn(8)

	for a, b := range t {
		fmt.Printf("%d - %s", a, b)
		fmt.Println()
	}
}

func withReturnValues() (i int, s string) {
	//i = 10
	s = "naked return"

	return
}
