package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

const query = `
select
	users.id,
	users.name,
	address.line1,
	address.state
from users
left join address on users.id = address.user_id
`

func main() {

	// query, err := os.ReadFile("examples/queries/users.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tree, err := pg_query.ParseToJSON(string(query))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", tree)
}
