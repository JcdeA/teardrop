package db

import (
	"context"
	"log"

	"github.com/fosshostorg/teardrop/ent"
)

var DBClient *ent.Client
var Ctx = context.Background()

// Returns an ent client; shares a client using a global variable
func Connect() *ent.Client {
	var err error

	if DBClient == nil {
		DBClient, err = ent.Open("sqlite3", "file:lmao?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to db: %v", err)
		}
	}

	return DBClient

}
