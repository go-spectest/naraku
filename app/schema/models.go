// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package schema

import (
	"time"
)

// user table
type User struct {
	// ID is user id that expressed in ULID（Universally Unique Lexicographically Sortable Identifier）
	ID string
	// Name is user name
	Name string
	// Biography is user self introduction
	Biography string
	// Email is user email address
	Email string
	// CreatedAt is the date that record was created
	CreatedAt time.Time
	// UpdatedAt is the date record was updated
	UpdatedAt time.Time
}
