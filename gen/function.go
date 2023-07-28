package gen

type Function struct {
	Name    string // UpdateDebtor
	Comment string // Update the debtor details
	SQL     string // update ... from ...

	// input
	// Code IN -> SQL IN
	// SQL OUT -> Code OUT

	Imports []string

	Args []CodeVar

	Input  []Param
	Output []Param

	Return []CodeVar
}

type CodeType struct {
	Name   string
	Import string

	Parent     *CodeType
	Properties []Property
}

type Property struct {
	Name string
	Type CodeType
}

type CodeVar struct {
	Name string   // debtor.FirstName
	Type CodeType // string

	IsPointer bool
	IsArray   bool

	Parent *CodeVar // debtor.MOdel
	Fields []CodeVar
}

func (v *CodeVar) AddField(field CodeVar) {
	field.Parent = v
	v.Fields = append(v.Fields, field)
}

// Param maps a database parameter to a code variable
type Param struct {
	DBField string // $3
	CodeVar CodeVar
}

// type QueryType string

// const (
// 	One      QueryType = "one"
// 	Many     QueryType = "many"
// 	Exec     QueryType = "exec"
// 	ExecRows QueryType = "execrows"
// )

// type QueryFunction struct {
// 	Name string
// 	Type QueryType
// 	SQL  string

// 	Query    Query
// 	Function Function
// }

// type Function interface {
// 	Args() []Object
// 	Return() Object
// }

// type Type struct {
// 	Name   string
// 	Import string
// }

// type Object struct {
// 	Name      string
// 	Type      Type
// 	IsPointer bool
// 	IsSlice   bool
// 	Import    string
// 	Fields    []Object
// }

// // Database
// type DBField interface {
// 	Name() string // $1 or :user_id / id or name
// 	Type() string // text, int4 or uuid
// 	IsArray() bool
// 	Nullable() bool
// }
// type Query interface {
// 	Params() []DBField
// 	Fields() []DBField
// }

// type Mapper interface {
// 	DBFieldToObject(DBField) Object
// }
