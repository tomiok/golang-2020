package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	chDelivery, err := ch.Consume(
		"gophers",
		"",
		true,
		false,
		false,
		false, nil)


	if err != nil {
		log.Fatal(err)
	}

	noStop := make(chan bool)

	go func() {
		for delivery := range chDelivery {
			fmt.Println("msg: " + string(delivery.Body))
		}
	}()

	<- noStop
}
