package main

import (
	"fmt"
	"sync"
	"time"
)

const numOfElf = 12

var (
	elfs   = make([]elf, numOfElf)
	broken int
)

type elf struct {
	name string
}

func elfsAreWorking(extWG *sync.WaitGroup, callingSanta chan bool) {

	var (
		wg sync.WaitGroup
	)

	chBrokenToy := make(chan bool)

	wg.Add(len(elfs))

	for range elfs {
		go makeToy(&wg, chBrokenToy)
	}

	wg.Wait()
	extWG.Done()
}

func makeToy(wg *sync.WaitGroup, ch chan bool) {
	fmt.Println("elf crafting a toy")
	if time.Now().Nanosecond()%2 == 1 { //hardcoded, change this for a 33% chance of a broken toy
		fmt.Println("toy is broken...")
		wg.Wait()
		go santaHelp(ch)
		<-ch
		wg.Done()
	} else {
		fmt.Println("toy is done")
		wg.Done()
	}

}
var i int
func santaHelp(ch chan bool) {
	i ++

	if i == 3 {
		ch <- true
	}
}