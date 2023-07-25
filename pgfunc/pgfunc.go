package pgfunc

import (
	"strings"

	pgoid "github.com/lib/pq/oid"
)

type QueryType string

const (
	One      QueryType = "one"
	Many     QueryType = "many"
	Exec     QueryType = "exec"
	ExecRows QueryType = "execrows"
)

type Function struct {
	Name  string
	Query PGQuery
	Type  QueryType

	// If num of fields is greater that query_parameter_limit then use
	// Args.Fields to get the argument list
	Args GoStruct

	Return GoStruct
}

func NewFunction(name string, query PGQuery, qtype QueryType, typemap TypeMap) Function {
	var args, ret GoStruct
	args.Fields = mapTypes(query.Params, typemap)
	ret.Fields = mapTypes(query.Fields, typemap)
	return Function{
		Name:   name,
		Query:  query,
		Type:   qtype,
		Args:   args,
		Return: ret,
	}
}

func mapTypes(pgAttrs []PGAttr, typemap TypeMap) []GoField {
	var goFields []GoField
	for _, attr := range pgAttrs {
		goType := typemap[attr.Name]
		goField := GoField{
			Name:      attr.Name,
			Type:      goType,
			IsPointer: false,
			IsSlice:   attr.Type.IsArray(),
		}
		goFields = append(goFields, goField)
	}
	return goFields
}

// Go

type GoType struct {
	Name   string
	Import string
}

type GoField struct {
	Name      string
	Type      GoType
	IsPointer bool // Should this be on the GoType ??
	IsSlice   bool // Should this be on the GoType ??
}

type GoStruct struct {
	Name   string
	Import string
	Fields []GoField
}

// PostgreSQL

type PGType struct {
	Name string
	OID  uint32
}

func NewPGType(oid uint32) PGType {
	name := pgoid.TypeName[pgoid.Oid(oid)]
	return PGType{
		Name: strings.ToLower(name),
		OID:  oid,
	}
}

func (t *PGType) IsArray() bool {
	return t.Name[0:1] == "_"
}

type PGAttr struct {
	Name     string // $1 or :user_id / id or name
	Type     PGType // text, int4 or uuid
	Nullable bool
}
type PGQuery struct {
	SQL    string
	Params []PGAttr
	Fields []PGAttr
}
