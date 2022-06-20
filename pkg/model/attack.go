package model

import (
	"time"
)

type Attack struct {
	ID			int64		`db:"id" json:"id"`
	CreatedAt	time.Time 	`db:"created_at" json:"createdAt"`
	UpdatedAt 	time.Time 	`db:"updated_at" json:"updatedAt"`
	Target		string		`db:"target" json:"target"`
}

type AttackResponse struct {
	Attack *Attack `json:"attack"`
}

type CreateAttack struct {
	Target	string	`db:"target" json:"target"`
}
