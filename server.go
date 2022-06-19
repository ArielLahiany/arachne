package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addr = flag.String(
		"addr",
		"0.0.0.0",
		"Listening address of the server on the",
	)
	port = flag.Int(
		"port",
		8080,
		"Listening port of the server",
	)
)

type JSONResponse struct {
	Body string `json:"body"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := &JSONResponse{
		Body: "Hello, world!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func main() {
	flag.Parse()
	log.Printf("Listening and serving on %s:%v",
		*addr,
		*port,
	)
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(
		fmt.Sprintf(
			"%s:%v",
			*addr,
			*port,
		),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}
