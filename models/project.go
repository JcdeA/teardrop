package models

type Project struct {
	Users          []User
	Budget         int
	APIKey         []APIKey
	ContainerImage Image
}
