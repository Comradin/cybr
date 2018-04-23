package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
	"strings"
)

type Check struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {

	var content []byte
	var Checks []Check

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
			content, err = ioutil.ReadFile(configDir + "/" + file.Name())
			if err != nil {
				log.Fatal("Could not read config file")
			}

			// create an Check instance and unmarshal the raw json into
			simpletcp := Check{}
			err := json.Unmarshal(content, &simpletcp)
			if err != nil {
				log.Fatal("Could not read config")
			}

			Checks = append(Checks, simpletcp)

		}
	} else {
		log.Fatal("No config found")
	}



	for _, check := range Checks {

		// Simple TCP Check
		conn, err := net.DialTimeout(strings.ToLower(check.Type), fmt.Sprintf("%s:%d", check.Host, check.Port), 2*time.Second)
		if err != nil {
			log.Printf("CYBR (%s) - You cannot be reached: %v!\n", check.Name, err)
		} else {
			log.Printf("CYBR (%s) - You can be reached!", check.Name)
			conn.Close()
		}

	}
}
