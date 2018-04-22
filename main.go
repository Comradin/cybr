package main

import (
	"net"
	"log"
)

func main() {

	// Simple TCP Check
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Println( "Connection refused!")
	} else {
		log.Println("Connection successful!")
		conn.Close()
	}
	
}
