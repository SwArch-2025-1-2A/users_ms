// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID        uuid.UUID
	Category  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type User struct {
	ID         uuid.UUID
	Name       string
	ProfilePic []byte
}

type UserInterest struct {
	UserID     uuid.UUID
	InterestID uuid.UUID
}
