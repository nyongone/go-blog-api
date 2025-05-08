package main

import (
	"context"
	"go-blog-api/ent/migrate"
	"go-blog-api/internal/config"
	"go-blog-api/internal/datastore"
	"log"
)

func main() {
	config.LoadEnv()

	client, err := datastore.NewClient()

	if err != nil {
		log.Fatalf("Failed connecting mysql database: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
}