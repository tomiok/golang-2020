package main

import (
	"fmt"
	"math/rand"
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
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	secs := r.Intn(maxSeconds)
	for i := 0; i < totalHorses; i++ {
		time.Sleep(time.Duration(secs) * time.Second)
		horse := horse{name: time.Now().String()}
		horses = append(horses, horse)
		fmt.Println("horse arrived!" + string(i))
	}
	fmt.Println("all the horses are here homie")
	wg.Done()
	return true
}
