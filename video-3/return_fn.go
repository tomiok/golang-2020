package main

import (
	"fmt"
)

func main() {
	fn := rFn("hello")
	fn() //	fmt.Printf("%T", fn)

	fnOperation := test(1, 2, operation)
	//fnOperation := test(1, 2, operation)
	//fnOperation2 := test(1, 2, operationExample2())

	i := fnOperation()
	//i := fnOperation2()
	fmt.Println(fmt.Printf("the result of the operation is: %d", i))
}

func operationExample2() func(int, int) int {
	return func(i, j int) int {
		return (i + 150) + (j + 200)
	}
}

func operation(i, j int) int {
	return (i + 150) + (j + 200)
}

func test(i, j int, fn func(i, j int) int) func() int {
	res := fn(i, j)
	return func() int {
		return res * 100
	}
}

func rFn(s string) func() {
	fmt.Println(s)

	return func() {
		fmt.Println("inside the return " + s)
	}
}
