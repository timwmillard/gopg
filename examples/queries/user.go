package queries

import "github.com/gofrs/uuid"

type User struct {
	UserID uuid.UUID `db:"id"`
	Name   string    `db:"first_name"`
}
