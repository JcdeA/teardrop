package models

import "github.com/natesales/aarch64-client-go"

type VMRole int

const (
	Master VMRole = iota
	Worker
)

type VM struct {
	aarch64.VM
	Role VMRole
}
