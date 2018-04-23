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

	var content []byte

	// get config directory from environment
	if configDir, exists := os.LookupEnv("cybr_config_dir"); exists {
		log.Println(configDir)
		// no check if configDir exists!
		files, err := ioutil.ReadDir(configDir)
		if err != nil {
			log.Fatal("Could not read config directory")
		}

		if len(files) == 0 {
			log.Fatal("No configurations found - exiting")
		}

		// list config file names
		for _, file := range files {
			fmt.Println(file.Name())
			content, err = ioutil.ReadFile(configDir +"/"+file.Name())
			if err != nil {
				log.Fatal("Could not read config file")
			}

		}
	} else {
		log.Fatal("No config found")
	}

	// create an Check instance and unmarshal the raw json into
	simpletcp := Check{}
	err := json.Unmarshal(content, &simpletcp)
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
