package main

import (
	"fmt"
	"math"
)

// why go?
// https://yourbasic.org/golang/advantages-over-java-python/

// download go
// https://golang.org/dl/

// IDE visual studio
// https://code.visualstudio.com/docs/languages/go

// scopes
// unexported & exported

// commands
// build -> go build main.go
// build many files -> go build *.go
// build named binary -> go build -o myName *.go

// run without build -> go run main.go
// run without build many files -> go run *.go

// variables & constants
// data types
// exported and unexported
// imports
// shorthand assignation i.e: a := 5
// slices
// built-in functions len() cap() make()
// range with slices
// dif between slices and arrays
// commands in go => go build, go run

const (
	myConst = 888787
)

var Exported = "i am exported"

func main() {
	//dataTypesZeroValues()
	// this is a one line comment

	/*
		this is
		a  multiline comment
	*/
	/*println("")
	dataTypesWithValues()

	println("")
	dataTypesGroup()

	println("")

	printNum(9) */

	//otherTypeOfAssignation()

	//usingExportedConst()
	//otherArrayTypes()

	fmt.Println(seeThisConditional(2))
	fmt.Println(seeThisConditional(11))
}

func dataTypesZeroValues() {
	// data types
	var q int     // zero value = 0
	var e string  // zero value = ""
	var t float32 // zero value = 0
	var b bool    // zero value = false

	fmt.Printf("%v, %v,%v,%v", q, e, t, b)
}

func dataTypesWithValues() {
	// data types
	var q int = 76
	var e string = "hello" // same as var e = "hello"
	var t float32 = 0.32
	var p = 0.64
	var b bool = true

	fmt.Printf("%v, %v,%v,%v, %v", q, e, t, b, p)
}

func dataTypesGroup() {
	var (
		q    = 1
		name = "tomas"
		b    = true
	)
	fmt.Printf("%v,%v,%v, %v", q, name, b, myConst)
}

func otherTypeOfAssignation() {
	myNUmber := 8
	fmt.Println(myNUmber)
	myString := "my name is tomas"
	fmt.Println(myString)
}

func checkingArrays() {
	arr := []int{1, 2, 3, 4, 5}

	// for loop
	for i := 0; i < 10; i++ {
		// ....
		break
	}

	for index, value := range arr {
		fmt.Printf("%v || %v ", index, value)
		fmt.Println("")
	}

	l := len(arr)

	fmt.Println(l)
}

func otherArrayTypes() {
	var sl []int // nil
	sl = []int{1, 2, 3, 4, 5}
	fmt.Printf("len is %d \n", len(sl))
	sl = append(sl, 6)
	fmt.Printf("len is %d \n", len(sl))

	sl2 := make([]int, 0, 5)

	sl2 = append(sl2, 1, 2, 3, 4, 5)

	for _, value := range sl2 {
		fmt.Println(value)
	}
}

func seeThisConditional(i int) bool {
	var b bool

	if i % 2 == 0 {
		b = true
		return b
	}

	b = false
	return b
}

//functions
//errrors
//struct - interfaces
//pointers
//data type nil


func usingExportedConst() {
	myPi := math.Pi
	fmt.Println(myPi)
}
