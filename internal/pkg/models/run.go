package models

import (
	"time"
)

type Run struct {
	Name        string
	Started     time.Time
	Id          int
	Project     Project
	StartedBy   User
	Environment Environment
	EnvVars     map[string]string
}
