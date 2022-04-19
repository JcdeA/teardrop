package models

type User struct {
	Name         string   `bson:"name"`
	Email        string   `bson:"email"`
	IsAdmin      bool     `bson:"isAdmin"`
	ProjectUUIDs []string `bson:"projects"`
}
