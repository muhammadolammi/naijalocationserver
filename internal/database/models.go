// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/google/uuid"
)

type City struct {
	ID         uuid.UUID
	Name       string
	Population string
	StateID    uuid.UUID
}

type Lga struct {
	ID      uuid.UUID
	Name    string
	StateID uuid.UUID
}

type State struct {
	ID      uuid.UUID
	Name    string
	Capital string
}
