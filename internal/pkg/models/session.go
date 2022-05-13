package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	User    User      `json:"user"`
	Expires time.Time `json:"expires"`
}

type User struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
	Email string    `json:"email"`
	Admin bool      `json:"admin"`
}
