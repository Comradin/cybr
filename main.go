package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"io/ioutil"
)

type Check struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {

	// get config directory from environment
	if config_dir, exists := os.LookupEnv("cybr_config_dir"); exists {
		log.Println(config_dir)
		// no check if config_dir exists!
		files, err := ioutil.ReadDir(config_dir)
		if err != nil {
			log.Fatal("Could not read config directory")
		}

		// list config file names
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		log.Fatal("No config found")
	}

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
