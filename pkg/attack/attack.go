package attack

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Attack struct {
	Target	string	`json:"target"`
	Logger 	log.Logger
}

type AttackInterface interface {
	AttackHandler(w http.ResponseWriter, r *http.Request)
	Launch(attack Attack) error
}

func NewAttack(target string) *Attack {
	attack := Attack{
		Target: target,
		Logger: *log.New(
			os.Stdout, "", log.LstdFlags,
		),
	}

	return &attack
}

func AttackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			decoder := json.NewDecoder(r.Body)
			var attack Attack
			err := decoder.Decode(&attack)
			if err != nil {
				log.Printf("%v", err)
			}
			log.Printf("Target: %v", attack.Target)
		default:
			log.Printf("Request: %v", r.Method)
	}
}
