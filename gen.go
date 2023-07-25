package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func newPGXConn(ctx context.Context, config Connection) (*pgx.Conn, error) {
	databaseURL := config.URL

	if databaseURL == "" {
		envVar := "DATABASE_URL"
		if config.EnvVar != "" {
			envVar = config.EnvVar
		}
		databaseURL = os.Getenv(envVar)
	}

	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connected to database: %s\n", databaseURL)
	return conn, nil
}

func generate(config Config) {
	ctx := context.Background()
	conn, err := newPGXConn(ctx, config.Connection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "database error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	for _, pkg := range config.Packages {
		queries, err := readQueries(pkg.Queries)
		if err != nil {
			fmt.Fprintf(os.Stderr, "generate error, package %s: %v\n", pkg.Name, err)
		}
		for _, query := range queries {
			_, err := parseSQL(conn, query)
			if err != nil {
				fmt.Fprintf(os.Stderr, "parse query error: %v\n", err)
				os.Exit(1)
			}
		}
	}
}
