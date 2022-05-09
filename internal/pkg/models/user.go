package models

import (
	"github.com/google/uuid"
)

type SessionUser struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
	Email string    `json:"email"`
	Admin bool      `json:"admin"`
}
