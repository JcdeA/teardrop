package db

import (
	"context"
	"log"

	"github.com/fosshostorg/teardrop/ent"
)

var DBClient *ent.Client
var Ctx = context.Background()

func Connect() {
	DBClient, err := ent.Open("sqlite3", "file:lmao?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	defer DBClient.Close()
}
