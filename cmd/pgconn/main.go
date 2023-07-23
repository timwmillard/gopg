package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

const query = `
-- name: GetUsers :many
-- @param $1 id uuid.UUID
-- @param :user_id uuid.UUID
-- 
-- @return User
-- @validation isValidUser
select
	users.id,
	users.name,
	address.line1,
	address.state
from users
left join address on users.id = address.user_id
-- where id = $1
;
`

type User struct {
	ID   uuid.UUID
	Name string
}

type Address struct {
	UserID uuid.UUID
	Line1  string
	State  *string
}

// QueryResultFormatsByOID controls the result format (text=0, binary=1) of a query by the result column OID.
func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// DATABASE_URL=postgresql://root:postgres@localhost:55432/gopg
	databaseURL := "postgresql://root:postgres@localhost:55432/gopg"
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// id := uuid.FromStringOrNil("5e367264-d991-4274-a851-72370fe1dada")

	desc, err := conn.Prepare(context.Background(), "get_users", "select * from users where id = $1 and name = $2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prepare statement error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Prepare = %+v\n", desc)

	// 	var user User
	// 	var address Address
	// 	// rows := conn.QueryRow(context.Background(), query)
	// 	rows, err := conn.Query(context.Background(), query)
	// 	if err != nil {
	// 		fmt.Println("Query error: ", err)
	// 	}
	// 	fmt.Println("----------------")
	// 	fmt.Printf("row(%T)=%+v\n", rows, rows)
	// 	fmt.Println("----------------")
	// 	fields := rows.FieldDescriptions()
	// 	for _, field := range fields {
	// 		fmt.Printf("Field=%+v\n", field)
	// 	}
	// 	fmt.Println("----------------")

	// 	// err = rows.Scan(
	// 	// 	&user.ID,
	// 	// 	&user.Name,
	// 	// 	&address.Line1,
	// 	// 	&address.State,
	// 	// )
	// 	// if err != nil {
	// 	// 	fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
	// 	// 	os.Exit(1)
	// 	// }

	// fmt.Println(user.ID, user.Name, address.Line1, address.State)
}
