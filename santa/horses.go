package main

import (
	"fmt"
	"math/rand"
	"time"
)

type horse struct {
	name string
}

var horses = make([]horse, 12)

func arrival() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	secs := r.Intn(6)
	fmt.Println(secs)

}
