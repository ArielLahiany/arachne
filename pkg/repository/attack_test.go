package repository

import (
	"testing"
)

func TestSelectSingle(t *testing.T) {
	r := AttackRepository{}
	attack, err := r.SelectSingle(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(
		"Identifier: %v, Created at: %v, Updated at: %v, Target: %v",
		attack.ID,
		attack.CreatedAt,
		attack.UpdatedAt,
		attack.Target,
	)
}
