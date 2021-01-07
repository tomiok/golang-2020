package main

import (
	"fmt"
	"sync"
)

const numOfElf = 12

var (
	elfs  = make([]elf, numOfElf)
	teamA = elfs[0:2]
	teamB = elfs[3:5]
	teamC = elfs[6:8]
	teamD = elfs[9:11]
	chToy = make(chan bool)
)

type elf struct {
	name string
}

func elfsAreWorking() {

	var wg sync.WaitGroup
}

func makeToy(elfs []elf, ch chan bool) {
	fmt.Println("elf crafting a toy")

	for i, _ := range elfs {

		if i == 1 { //hardcoded, change this for a 33% chance of a broken toy
			fmt.Println("toy is broken...")

		}
	}

}
