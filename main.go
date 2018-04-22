package main

import (
	"net"
	"log"
)

func main() {

	// Simple TCP Check
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Println( "CYBR - You cannot be reached!")
	} else {
		log.Println("CYBR - You can be reached!")
		conn.Close()
	}
}
