package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Response is the data returned by the server
type Response struct {
	Name string
}

func main() {
	var port string
	flag.StringVar(&port, "p", "1331", "listen on port")
	flag.Parse()

	http.HandleFunc("/", home)

	fmt.Printf("Server listening on port: %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %v \n", r.URL)
	name := strings.Trim(r.URL.Path, "/")

	profile := Response{name}

	response, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Simulates that does something...
	time.Sleep(2 * time.Second)

	w.Write(response)
}
