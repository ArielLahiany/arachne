package repository

import (
	"context"

	"github.com/ArielLahiany/arachne/pkg/database"
	"github.com/ArielLahiany/arachne/pkg/model"
)

type AttackRepository struct{}

func (r *AttackRepository) SelectSingle(id int) (*model.Attack, error) {
	connection := database.Connect()
	defer connection.Close()

	query := connection.QueryRow(
		context.Background(),
		"SELECT * FROM attack WHERE id=$1;",
		id,
	)

	var attack model.Attack
	err := query.Scan(
		&attack.ID,
		&attack.CreatedAt,
		&attack.UpdatedAt,
		&attack.Type,
		&attack.Target,
	)

	return &attack, err
}
