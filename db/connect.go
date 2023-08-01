package db

import (
	"context"
	"fmt"
	"gin-api/ent"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

// db connect
func Connect() (*ent.Client, error) {

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("SSL_MODE")))

	if err != nil {
		return nil, err
	}
	return client, nil
}

// Run the auto migration tool.
func Schema(client *ent.Client, ctx context.Context) error {
	if err := client.Schema.Create(ctx); err != nil {
		return err
	}
	return nil
}
