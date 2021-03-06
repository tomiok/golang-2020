package main

import "fmt"

func main() {
	c := Car{Model: 2000}
	//r(c)
	//fmt.Println(c.Model)
	fmt.Println("****")

	rr(&c)
	fmt.Println(c.Model)
}

func otherFunc() {
	car := new(Car) //pointer
	car2 := Car{
		Model: 2019,
		Color: "RED",
		Engine: CarEngine{
			Version: 8,
		},
		Line: &Line{LineName: "trend line"},
	} // reference

	//car3 := make([]Car, 1) // array of car

	fmt.Printf("%v", car)
	fmt.Println()
	fmt.Printf("%v", car2)

}

//Car is an exported structure
type Car struct {
	Model     int
	Color     string
	Engine    CarEngine
	Line      *Line // a point holds a memory address as a value
	Insurance Insurance
}

type Line struct {
	LineName string
}

type CarEngine struct {
	Version int
}

type Insurance interface {
}

// receiving an structure and change it, will not change the original value since we receive a copy of the original
func r(c Car) {
	fmt.Println(c.Model)
	c.Model = 2020
	fmt.Println(c.Model)
}

// receiving a pointer and change it inside the function will ALSO change the first one that the function received as a parameter
func rr(c *Car) {
	fmt.Println(c.Model)
	c.Model = 3030
	fmt.Println(c.Model)

}
