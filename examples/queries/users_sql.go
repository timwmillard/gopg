package queries

import (
	"context"

	"github.com/gofrs/uuid"
)

const getUsers = `
-- name: GetUsers :many
-- @param $1 id uuid.UUID
-- @param :user_id uuid.UUID
-- 
-- @return User
-- @validation isValidUser
select * from users
where id = $1;
`

// type GetUsersParams struct {
// 	ID uuid.UUID
// 	Name   string
// }

type GetUsersRow struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `
-- name: GetUser :one
-- @param $1 id uuid.UUID
-- 
-- @return User
select *
from users
where id = $1;
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.UserID, // ID
		&i.Name,   // Name
	)
	return i, err
}
