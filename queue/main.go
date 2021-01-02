package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//user; guest, password: guest
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("hello world")
}

