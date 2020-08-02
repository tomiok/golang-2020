package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	for {
		t := temp()
		localTemp := generateTemp(t)
		url := fmt.Sprintf("http://localhost:9192/temperatures/local/%d", localTemp.Temp)
		_, err := http.Post(url, "text/plain", nil)

		if err != nil {
			log.Printf("cannot send temp %s", err.Error())
		}
		log.Printf("temp sent %d", localTemp.Temp)
		time.Sleep(time.Second * 5)
	}
}

type LocalTemp struct {
	Temp int `json:"temp"`
}

func generateTemp(t int) LocalTemp {
	return LocalTemp{Temp: t}
}

func temp() int {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	return r.Intn(40)
}
