package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	totalHorses = 9
	maxSeconds  = 6
)

type horse struct {
	name string
}

var horses = make([]horse, totalHorses)

func horseArrival(wg *sync.WaitGroup) bool {
	for i := 0; i < totalHorses; i++ {
		secs := randomNumber(maxSeconds)
		time.Sleep(time.Duration(secs) * time.Second)
		horse := horse{name: time.Now().String()}
		horses = append(horses, horse)
		fmt.Println("horse arrived!")
	}
	fmt.Println("all the horses are here homie")
	wg.Done()
	return true
}
