package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting")
	start()
	go startingSanta()
	time.Sleep(5 * time.Second)
	wakeUpSanta()
	time.Sleep(5 * time.Second)
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
