package main

import (
	"fmt"
)

func main() {

	//forever := make(chan bool)
	fmt.Println("starting")
	start()
	go startingSanta()
	//time.Sleep(5 * time.Second)
	wakeUpSanta()
	fmt.Println("finish")
	//time.Sleep(5 * time.Second)
	//<-forever
}

func wakeUpSanta() {
	go func() { santaChUp <- true }()
}

func startingSanta() {
	santaRoutine()
}

func start() {
	santaStateZero()
}
