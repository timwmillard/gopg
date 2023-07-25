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
	Pointer   bool
	Slice     bool
}

func (gt *GoType) Definition() string {
	var name = gt.Name
	if gt.Pointer {
		name = "*" + name
	}
	if gt.Slice {
		name = "[]" + name
	}
	return name
}

type GoField struct {
	Name     string
	GoType   GoType
	DBColumn string
	Nilable  bool // Can the type hold nil, eg pointer or slice

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
	PgParam string // eg. $1
	PgType  string
	GoField GoField // eg id
}

type QueryColumn struct {
	Name     string // eg. user_id
	ColumnID int    // Which is the column index eg. 1, 2, 3
	PgType   string // text or numeric

	columnNullable bool
	joinNullable   bool
}

func (col *QueryColumn) Nullable() bool {
	return col.columnNullable || col.joinNullable
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
