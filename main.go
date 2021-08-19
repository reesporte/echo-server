package main

import (
	"log"
	"net"
)

const MAX_BYTES = 600

func handle(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, MAX_BYTES)
	n, err := conn.Read(data)
	if err != nil {
		log.Println("could not read data")
	}

	if n < MAX_BYTES {
		data = data[0:n]
	}

	log.Println(conn.RemoteAddr().String(), string(data))

	_, err = conn.Write(data)

	if err != nil {
		log.Println("could not write data")
	}

}

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("Could not listen on port 9999")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error accepting connection")
		}
		go handle(conn)
	}
}
