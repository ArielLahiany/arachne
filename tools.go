//go:build tools
// +build tools

package tools

import (
	_ "github.com/jackc/pgx/v4"
	_ "github.com/streadway/amqp"
)
