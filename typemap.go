package main

type TypeMap map[string]GoType // key = PG type, value = Go type

var DefaultTypeMap = TypeMap{
	"bool": GoType{
		Name:      "bool",
		ImportPkg: "",
	},
	"int2": GoType{
		Name:      "int16",
		ImportPkg: "",
	},
	"int4": GoType{
		Name:      "int32",
		ImportPkg: "",
	},
	"int8": GoType{
		Name:      "int64",
		ImportPkg: "",
	},
	"char": GoType{
		Name:      "int",
		ImportPkg: "",
	},
	"float4": GoType{
		Name:      "float32",
		ImportPkg: "",
	},
	"float8": GoType{
		Name:      "float64",
		ImportPkg: "",
	},
	"numberic": GoType{
		Name:      "decimal.Decimal",
		ImportPkg: "github.com/shopspring/decimal",
	},
	"text": GoType{
		Name:      "string",
		ImportPkg: "",
	},
	"uuid": GoType{
		Name:      "uuid.UUID",
		ImportPkg: "github.com/gofrs/uuid",
	},
}
