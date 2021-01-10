package main

import (
	"fmt"
	"sync"
	"time"
)

const numOfElf = 12

var (
	elfs  = make([]elf, numOfElf)
	teamA = elfs[0:2]
	teamB = elfs[3:5]
	teamC = elfs[6:8]
	teamD = elfs[9:11]
)

type elf struct {
	name string
}

func elfsAreWorking(extWG *sync.WaitGroup, callingSanta chan bool) {

	var wg sync.WaitGroup

	groupOfElf := make([][]elf, 4)
	groupOfElf = append(groupOfElf, teamA)
	groupOfElf = append(groupOfElf, teamB)
	groupOfElf = append(groupOfElf, teamC)
	groupOfElf = append(groupOfElf, teamD)

	wg.Add(len(groupOfElf))

	for _, group := range groupOfElf {
		go makeToy(group, callingSanta, &wg)
	}

	wg.Wait()
	extWG.Done()
}

func makeToy(elfs []elf, callingSanta chan bool, wg *sync.WaitGroup) {
	fmt.Println("elf crafting a toy")

	for i, _ := range elfs {

		if i == 1 { //hardcoded, change this for a 33% chance of a broken toy
			fmt.Println("toy is broken...")
			callingSanta <- true
			wg.Wait()
			time.Sleep(2 * time.Second)
			wg.Done()
		}
	}
}
