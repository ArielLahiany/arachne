package model

type Attack struct {
	ID		int64	`db:"id" json:"id"`
	Target	string	`db:"target" json:"target"`
}

type AttackResponse struct {
	Attack *Attack `json:"attack"`
}

type CreateAttack struct {
	Target	string	`db:"target" json:"target"`
}
