package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Locator struct {
	User     string
	Password string
	Hostname string
	Port     string
}

func NewLocator() *Locator {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "PRODUCTION" {
		return &Locator{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASS"),
			Hostname: os.Getenv("POSTGRES_HOSTNAME"),
			Port:     os.Getenv("POSTGRES_PORT"),
		}
	}

	return &Locator{
		User:     "arachne",
		Password: "arachne",
		Hostname: "localhost",
		Port:     "5432",
	}
}

func Connect() *pgxpool.Pool {
	locator := NewLocator()
	connection, err := pgxpool.Connect(
		context.Background(),
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/postgres",
			locator.User,
			locator.Password,
			locator.Hostname,
			locator.Port,
		),
	)
	if err != nil {
		log.Printf(
			"Unable to connect to PostgreSQL: %v",
			err,
		)
	}

	return connection
}
