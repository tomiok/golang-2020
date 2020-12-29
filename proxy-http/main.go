package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connect()
}

func connect() {
	log.Println("connection starting")

	l, err := net.Listen("tcp", "localhost:4545")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()

	remoteConnection, err := net.Dial("tcp", "google.com:443")

	if err != nil {
		log.Println(err.Error())
		return
	}
	defer remoteConnection.Close()

	go io.Copy(remoteConnection, conn)
	io.Copy(conn, remoteConnection)
}

func copyWhenErr(conn net.Conn) {
	defer conn.Close()
	 for {
	 	var buf [128]byte
	 	n, err := conn.Read(buf[:])

	 	if err != nil {
	 		log.Println(err.Error())
			return
		}
		 _ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		 os.Stderr.Write(buf[:n])
	 }
}
