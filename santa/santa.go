package main

import (
	"fmt"
	"sync"
)

var (
	santaChUp    = make(chan bool)
	santaChSleep = make(chan bool)
	s            = santa{
		mutex: sync.Mutex{},
	}
)

func santaStateZero() {
	go func() {
		santaChSleep <- true
	}()
}

type santa struct {
	mutex sync.Mutex
}

//santaRoutine Santa when all the reindeer are at home and 3 elf have troubles with the toy's production
func santaRoutine() {
	for {
		select {
		case <-santaChUp:
			fmt.Println("Santa is up now...")
			s.mutex.Lock()
		case <-santaChSleep:
			fmt.Println("Santa is sleeping...")
		}
	}

}
