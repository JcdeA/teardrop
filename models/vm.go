package models

import "github.com/natesales/aarch64-client-go"

type VMRole int

const (
	Server VMRole = iota // nomad server node - orchestrates clients
	Client               // nomad client role - executes jobs
)

type VM struct {
	aarch64.VM
	Role VMRole
}
