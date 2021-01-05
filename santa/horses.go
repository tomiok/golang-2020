package main

import (
	"math/rand"
	"time"
)

const (
	totalHorses = 12
	maxSeconds
)

type horse struct {
	name string
}

var horses = make([]horse, totalHorses)

func arrival() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	secs := r.Intn(maxSeconds)
	for i := 0; i < totalHorses; i++ {
		time.Sleep(time.Duration(secs) * time.Second)
		horse := horse{name: time.Now().String()}
		horses = append(horses, horse)
	}

}
