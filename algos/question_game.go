package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	done := make(chan bool)
	var count int
	scanner := bufio.NewScanner(os.Stdin)
	questions := questions()
	go func() {
		for _, q := range questions {
			fmt.Print(q[0])
			_ = scanner.Scan()
			answer := scanner.Text()
			if answer == q[1] {
				count++
			}
		}
		done <- true
	}()

	game(time.NewTicker(3*time.Second), done)
	fmt.Println(count)
}

func game(t *time.Ticker, done chan bool) {
	select {
	case <-t.C:
		fmt.Println("time's up dude")
	case <-done:
	//default will not check any one and exit the select immediately
	}
}

func questions() [][]string {
	return [][]string{{"1+1: ", "2"}, {"1+2: ", "3"}, {"1+3: ", "4"}}
}
