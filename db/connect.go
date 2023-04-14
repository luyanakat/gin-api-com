package db

import (
	"context"
	"gin-api/ent"

	_ "github.com/lib/pq"
)

// db connect
func Connect() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=student password=abc@123 sslmode=disable")
	if err != nil {
		return nil, err
	}
	

	return client, nil
}

// Run the auto migration tool.
func Schema(client *ent.Client) error {
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}
