package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5"
)

func parseQueries(locations Paths) error {
	files, err := Glob(locations)
	if err != nil {
		return err
	}

	var errs error
	for _, filename := range files {
		blob, err := os.ReadFile(filename)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		parseSQL(string(blob))
	}
	// info, err := os.Stat(location)
	// if err != nil {
	// 	return err
	// }

	// if info.IsDir() {
	// 	return errors.New("directory queries currently not supported")
	// }

	// data, err := os.ReadFile(location)
	// if err != nil {
	// 	return fmt.Errorf("read file %s error: %v", location, err)
	// }

	return errs
}

func parseSQL(queries string) {
	fmt.Printf("sql = %v\n", queries)
}

func generate(config Config) {
	fmt.Printf("config=%v\n", config)

	var databaseURL string

	if config.Connection.URL == "" {
		envVar := "DATABASE_URL"
		if config.Connection.EnvVar != "" {
			envVar = config.Connection.EnvVar
		}
		databaseURL = os.Getenv(envVar)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "database error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	fmt.Printf("Connected to database: %s\n", databaseURL)

	for _, pkg := range config.Packages {
		err := parseQueries(pkg.Queries)
		if err != nil {
			fmt.Fprintf(os.Stderr, "generate error, package %s: %v\n", pkg.Name, err)
		}
	}
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
