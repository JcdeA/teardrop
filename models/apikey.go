package models

type Scope int64

const (
	Read Scope = iota
	Write
	Admin
)

type APIKey struct {
	Hash  []byte
	Perms []Scope
}
