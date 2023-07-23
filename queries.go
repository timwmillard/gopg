package main

type FunctionType string

const (
	One      FunctionType = "one"
	Many     FunctionType = "many"
	Exec     FunctionType = "exec"
	ExecRows FunctionType = "execrows"
)

type GoType struct {
	Name      string // eg. uuid
	ImportPkg string // eg. github.com/gorfs/uuid
}

type GoField struct {
	Name     string
	GoType   GoType
	DBColumn string

	// Optional transformation and validation functions
	Pre      TransformFunc
	Post     TransformFunc
	Validate ValidateFunc
}

type GoStruct struct {
	Name   string
	Fields []string
}

type Param struct {
	PgParam string  // eg. $1
	GoField GoField // eg id
}

type QueryColumn struct {
	Name     string // eg. user_id
	ColumnID int    // Which is the column index eg. 1, 2, 3
	PgType   string // text or numeric
	Nullable bool
}

type QueryFunction struct {
	SQL string

	Name string
	Type FunctionType

	Params []Param

	QueryReturning []QueryColumn
	FunctionReturn GoStruct
}

type TransformFunc func(any) any
type ValidateFunc func(any) (any, error)
