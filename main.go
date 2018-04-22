package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Check struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {

	// raw json config
	raw := `
{
  "Name": "Simple TCP",
  "type": "TCP",
  "host": "localhost",
  "port": 8081
}
`

	// create an Check instance and unmarshal the raw json into
	simpletcp := Check{}
	err := json.Unmarshal([]byte(raw), &simpletcp)
	if err != nil {
		log.Fatal("Could not read config")
	}

	// Simple TCP Check
	conn, err := net.Dial(simpletcp.Type, fmt.Sprintf("%s:%s", simpletcp.Host, simpletcp.Port))
	if err != nil {
		log.Printf("CYBR (%s) - You cannot be reached!\n", simpletcp.Name)
	} else {
		log.Printf("CYBR (%s) - You can be reached!", simpletcp.Name)
		conn.Close()
	}
}
