package main

import (
	"fmt"
	"time"
)

// ** panic **
// panic returns the "panic" to the callee, ending in the one that generated the goroutine.

// ** defer **
// is the last call in the function body
// doesnt matter where I put the deferred call
// defer alone is not going to handle panic at all

// ** recover **
// recovers from a panic in a function scope.
// you can use recover with defer to ensure the last call

func main() {
	fmt.Println("hi from main")
	gWithRecover()
}

func gWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("panic message is: %v", r))
			fmt.Println("i'm still alive, recovered")
		}
	}()

	gWithPanic(99)
}

func gWithPanic(i int) {
	defer fmt.Println("not defer yet")
	fmt.Println(i)
	panic("going down dude")
}

func deferTest() {
	defer timeTracker(time.Now(), "main")

	fmt.Println("first")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("second")
}

func timeTracker(t time.Time, name string) {
	elapsed := time.Since(t)

	fmt.Println(fmt.Sprintf("the function %s took %s", name, elapsed.String()))
}
