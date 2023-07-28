package gen

type Function struct {
	Name    string // UpdateDebtor
	Comment string // Update the debtor details
	SQL     string // update ... from ...

	// input
	// Code IN -> SQL IN
	// SQL OUT -> Code OUT

	Package string
	Imports []string

	Args []CodeVar

	Input  []Param
	Output []Param

	Return []CodeVar
}

type CodeType struct {
	Name   string
	Import string

	Properties []Property
}

type Property struct {
	Name string
	Type CodeType

	Parent *CodeType
}

func (c *CodeType) AddProperty(property Property) {
	property.Parent = c
	c.Properties = append(c.Properties, property)
}

func (c *CodeType) IsScalar() bool {
	return c.Properties == nil
}

func (c *CodeType) IsObject() bool {
	return c.Properties != nil
}

type CodeVar struct {
	Name string   // debtor.FirstName
	Type CodeType // string

	IsPointer bool
	IsArray   bool
}

func (c *CodeVar) Property(index int) Property {
	return c.Type.Properties[index]
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
