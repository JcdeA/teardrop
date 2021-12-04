package models

type Organization struct {
	Projects []Project
	Users    []User
	Name     string
	Budget   int
}
