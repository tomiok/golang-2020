package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	leave := make(chan os.Signal, 1)
	callingSanta := make(chan bool, 1)
	var wg sync.WaitGroup
	fmt.Println("starting")
	start()

	wg.Add(2)
	go startingSanta()
	go horseArriving(&wg)
	go elfsAreWorking(&wg, callingSanta)
	wg.Done()
	time.Sleep(5 * time.Second)

	go func() {
		signal.Notify(leave, os.Interrupt)
	}()

	select {
	case <-callingSanta:
		fmt.Println("santa is called")
		wakeUpSanta()
	case <-leave:
		fmt.Println("outta")
	}
	wg.Wait()
	fmt.Println("Santa is ready to go!")
}

func horseArriving(wg *sync.WaitGroup) {
	horseArrival(wg)
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
