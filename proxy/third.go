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
	connectT()
}

func connectT() {
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
		go proxyT(conn)
	}
}

func proxyT(conn net.Conn) {
	defer conn.Close()

	upstream, err := net.Dial("tcp", "google.com:443")

	if err != nil {
		log.Println(err.Error())
		return
	}
	defer upstream.Close()

	go io.Copy(upstream, conn)
	io.Copy(conn, upstream)
}

func copyOnErrT(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println(fmt.Sprintf("written = %d, error: %s", n, err.Error()))
			return
		}
		_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		_, _ = os.Stderr.Write(buf[:n])
	}

}
