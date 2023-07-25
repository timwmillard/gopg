package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopsql/pgfunc"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func readQueries(locations Paths) (queries []string, err error) {
	files, err := Glob(locations)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		blob, readErr := os.ReadFile(filename)
		if err != nil {
			err = errors.Join(err, readErr)
			continue
		}
		queries = append(queries, string(blob))
	}

	return queries, err
}

func parseSQL(conn *pgx.Conn, sql string) (pgfunc.PGQuery, error) {

	var query pgfunc.PGQuery

	num := rand.Int()
	name := fmt.Sprintf("gopg_prepare_%d", num)

	var desc *pgconn.StatementDescription
	var err error
	desc, err = conn.Prepare(context.Background(), name, sql)
	if err != nil {
		return pgfunc.PGQuery{}, err
	}

	// TODO: oid lookup could from quering pg_type table instead.
	// This would give a dynamic lookup and reflect custom types for the
	// database.
	for i, oid := range desc.ParamOIDs {
		param := pgfunc.PGAttr{
			Name: fmt.Sprintf("$%d", i+1),
			Type: pgfunc.NewPGType(oid),
		}
		query.Params = append(query.Params, param)
	}
	for _, field := range desc.Fields {
		attr := pgfunc.PGAttr{
			Name:     field.Name,
			Type:     pgfunc.NewPGType(field.DataTypeOID),
			Nullable: false,
		}
		query.Fields = append(query.Fields, attr)
	}

	return query, nil
}

type Paths []string

func (p *Paths) UnmarshalJSON(data []byte) error {
	if string(data[0]) == `[` {
		var out []string
		if err := json.Unmarshal(data, &out); err != nil {
			return nil
		}
		*p = Paths(out)
		return nil
	}
	var out string
	if err := json.Unmarshal(data, &out); err != nil {
		return nil
	}
	*p = Paths([]string{out})
	return nil
}

func (p *Paths) UnmarshalYAML(unmarshal func(interface{}) error) error {
	out := []string{}
	if sliceErr := unmarshal(&out); sliceErr != nil {
		var ele string
		if strErr := unmarshal(&ele); strErr != nil {
			return strErr
		}
		out = []string{ele}
	}

	*p = Paths(out)
	return nil
}

// Return a list of SQL files in the listed paths. Only includes files ending
// in .sql. Omits hidden files, directories, and migrations.
func Glob(paths []string) ([]string, error) {
	var files []string
	for _, path := range paths {
		f, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("path %s does not exist", path)
		}
		if f.IsDir() {
			listing, err := os.ReadDir(path)
			if err != nil {
				return nil, err
			}
			for _, f := range listing {
				files = append(files, filepath.Join(path, f.Name()))
			}
		} else {
			files = append(files, path)
		}
	}
	var sqlFiles []string
	for _, file := range files {
		if !strings.HasSuffix(file, ".sql") {
			continue
		}
		if strings.HasPrefix(filepath.Base(file), ".") {
			continue
		}
		sqlFiles = append(sqlFiles, file)
	}
	return sqlFiles, nil
}
