package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numOfElf             = 12
	brokenToys           = 3
	totalProbability     = 100
	brokenToyProbability = 33
)

var (
	elfs    = make([]elf, numOfElf)
	counter int
	ticker  = time.NewTicker(20 * time.Second)
)

type elf struct {
	name string
}

func elfsAreWorking(extWG *sync.WaitGroup, chCallingSanta chan bool) {

	var wg sync.WaitGroup

	chBrokenToy := make(chan bool)
	wg.Add(len(elfs))

	for range elfs {
		go makeToy(&wg, chBrokenToy, chCallingSanta)
	}

	wg.Wait()
	extWG.Done()
}

func makeToy(wg *sync.WaitGroup, chBrokenToy, chCallingSanta chan bool) {
	fmt.Println("elf crafting a toy")
	rand := randomNumber(totalProbability)

	if rand < brokenToyProbability {
		counter++
		go santaHelp(chBrokenToy, chCallingSanta)
		select {
		case <-ticker.C:
			chBrokenToy <- true
		}
		<-chBrokenToy
	} else {
		fmt.Println("toy is done")

	}
	wg.Done()
}

func santaHelp(chBrokenToy, chCallingSanta chan bool) {
	fmt.Println(fmt.Sprintf("toy is broken... and counter: %d and need %d", counter, brokenToys))
	if counter == 3 {
		fmt.Println("santa fixing the toys")
		time.Sleep(2 * time.Second)
		chBrokenToy <- true
		chCallingSanta <- true
		counter = 0
	}
}
