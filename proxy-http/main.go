package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
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

	remoteConnection, err := net.Dial("tcp", "localhost:3001")

	if err != nil {
		log.Println(err.Error())
		return
	}
	remoteConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))
	remoteConnection.SetReadDeadline(time.Now().Add(10 * time.Second))
	defer remoteConnection.Close()

	go func() {
		io.Copy(remoteConnection, conn)
	}()

	var buff [150]byte
	n, _ := remoteConnection.Read(buff[:])
	s := string(buff[:n])
	rdr := strings.NewReader(s)
	fmt.Println(s)
	io.Copy(conn, rdr)
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
