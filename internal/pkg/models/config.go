package models

import "github.com/palantir/go-githubapp/githubapp"

type Config struct {
	Github githubapp.Config `yaml:"github"`
}
