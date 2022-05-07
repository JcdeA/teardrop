// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProjectsTable is the table that holds the projects relation/edge. The primary key declared below.
	ProjectsTable = "project_users"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "accounts"
	// AccountsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountsInverseTable = "accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "user_accounts"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldImage,
	FieldCreateAt,
	FieldUpdateAt,
}

var (
	// ProjectsPrimaryKey and ProjectsColumn2 are the table columns denoting the
	// primary key for the projects relation (M2M).
	ProjectsPrimaryKey = []string{"project_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
