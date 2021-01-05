package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	leave := make(chan os.Signal, 1)

	fmt.Println("starting")
	start()
	go startingSanta()
	time.Sleep(5 * time.Second)
	go horseArriving()
	time.Sleep(5 * time.Second)

	go func() {
		signal.Notify(leave, os.Interrupt)
	}()

	select {
	case <-leave:
		fmt.Println("outta")
	}
}

func horseArriving() {
	arrival()
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
