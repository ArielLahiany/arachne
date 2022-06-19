package router

import (
	"log"
	"net/http"
)

type AttackRouter struct {}

func NewAttackRouter() *AttackRouter {
	router := &AttackRouter{}

	return router
}

type AttackRouterInterface interface {
	GetAttack(w http.ResponseWriter, r *http.Request)
}

func (a *AttackRouter) GetAttack(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.Method)
}
