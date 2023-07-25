package main

import (
	"context"
	"testing"
)

const sql = `-- name: GetUser :one
-- @param $1 id uuid.UUID
-- 
-- @return User
select * from users
where id = $1;
`

func Test_parseSQL(t *testing.T) {
	ctx := context.Background()
	config, _ := readConfig("")
	conn, _ := newPGXConn(ctx, config.Connection)

	query, err := parseSQL(conn, sql)
	if err != nil {
		t.Error(err)
	}

	t.Logf("query = %+v", query)

	if len(query.Params) != 1 {
		t.Errorf("num of params should be 1 param but got %v", len(query.Params))
	}
	if len(query.Fields) != 2 {
		t.Errorf("num of fields should be 2 param but got %v", len(query.Fields))
	}
}
