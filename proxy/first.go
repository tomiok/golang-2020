package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connectF()
}

func connectF() {
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
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		// then copy on err
		n, err := io.Copy(os.Stderr, conn)
		if err != nil {
			log.Println(fmt.Sprintf("written = %d, error: %s", n, err.Error()))
		}
	}
}

func copyOnErrF(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)
	if err != nil {
		log.Println(fmt.Sprintf("written = %d, error: %s", n, err.Error()))
	}
}
