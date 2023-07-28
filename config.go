package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const ConfigVersion = 1

type Config struct {
	Version int `yaml:"version,omitempty"`

	Connection Connection `yaml:"connection,omitempty"`
	TypeMap    TypeMap    `yaml:"type_map,omitempty"`
	Packages   []Package  `yaml:"packages,omitempty"`
	Options    Options    `yaml:"options,omitempty"`
}

type Connection struct {
	URL    string `yaml:"url,omitempty"`
	EnvVar string `yaml:"env_var,omitempty"`
}

type GenType string

const (
	GenPGX GenType = "pgx"
	GenStd GenType = "std"
)

type Package struct {
	Name    string  `yaml:"name,omitempty"`
	Path    string  `yaml:"path,omitempty"`
	Queries Paths   `yaml:"queries,omitempty"`
	Gen     GenType `yaml:"gen,omitempty"`
}

type Options struct {
	// Maximum number of fields in param list before creating a struct
	MaxParamList int `yaml:"max_param_list,omitempty"`
}

func DefaultConfig() Config {
	return Config{
		Version: ConfigVersion,
		Connection: Connection{
			EnvVar: "DATABASE_URL",
		},
		TypeMap:  DefaultTypeMap,
		Packages: []Package{},
		Options: Options{
			MaxParamList: 0,
		},
	}
}

func readConfig(filename string) (Config, error) {
	if filename == "" {
		filename = "sqlgen.yaml"
	}

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file %s: %w", filename, err)
	}

	config := DefaultConfig()
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding config file %s: %w", filename, err)
	}

	return config, nil
}
