package main

import (
	"log"
	"net"
)

func main() {
	port := "8888"
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("error listening port %v: %v\n", port, err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Printf("error listening port %v: %v\n", port, err)
		conn.Close()
	}
}
