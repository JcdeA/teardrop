package main

import (
	"context"
	"log"

	"github.com/fosshostorg/teardrop/api"
	"github.com/fosshostorg/teardrop/ent"
	"github.com/fosshostorg/teardrop/ent/migrate"
)

func main() {
	client, err := ent.Open("sqlite3", "file:lmao?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	api.StartAPI()
}
