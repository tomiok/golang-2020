package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	connect()
}

func connect() {
	log.Println("starting server")
	l, err := net.Listen("tcp", "localhost:9995")

	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go copyOnErr(conn)
	}
}

func copyOnErr(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)

	if err != nil {
		log.Println(fmt.Sprintf("written = %d, error: %s", n, err.Error()))
	}

}
