package model

import (
	"time"
)

type Attack struct {
	ID			int64		`db:"id" json:"id"`
	CreatedAt	time.Time 	`db:"created_at" json:"createdAt"`
	UpdatedAt 	time.Time	`db:"updated_at" json:"updatedAt"`
	Type		string		`db:"type" json:"type"`
	Target		string		`db:"target" json:"target"`
}

type AttackResponse struct {
	Attack *Attack `json:"attack"`
}

type CreateAttack struct {
	Type	string	`db:"type" json:"type"`
	Target	string	`db:"target" json:"target"`
}
