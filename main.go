package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var (
	content []byte
	Checks  []Check
)

type Check struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	// Initialize the gorilla/mux router
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/api/listChecks", listChecksHandler)
	r.HandleFunc("/api/addCheck", addChecksHandler)

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

	log.Fatal(http.ListenAndServe(":8082", r))

}

// rootHandler
// Will list all status checks by iterating over the []Checks slice
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Number of checks defined: %d\n", len(Checks))
	for _, check := range Checks {

		// Simple TCP Check
		conn, err := net.DialTimeout(strings.ToLower(check.Type), fmt.Sprintf("%s:%d", check.Host, check.Port), 2*time.Second)
		if err != nil {
			fmt.Fprintf(w,"CYBR (%s) - You cannot be reached: %v!\n", check.Name, err)
		} else {
			fmt.Fprintf(w,"CYBR (%s) - You can be reached!\n", check.Name)
			conn.Close()
		}
	}
}

// listChecksHandler
// Will print a simplified List of all configured checks
func listChecksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Number of checks defined: %d\n", len(Checks))
	for _, check := range Checks {
		fmt.Fprintf(w, "%s - %s\n", check.Type, check.Name)
	}
}

// addCheckHandler
// Will provide a way to add more checks by calling this API endpoint
func addChecksHandler(w http.ResponseWriter, r *http.Request) {
	// Intentionally empty
}