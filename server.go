package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ArielLahiany/arachne/pkg/attack"
)

var (
	addr = flag.String(
		"addr",
		"0.0.0.0",
		"Listening address of the server",
	)
	port = flag.Int(
		"port",
		8080,
		"Listening port of the server",
	)
)

func main() {
	flag.Parse()
	log.Printf("Listening and serving on %s:%v",
		*addr,
		*port,
	)
	http.HandleFunc("/attack", attack.AttackHandler)
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
