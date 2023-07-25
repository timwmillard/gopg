package pgfunc

type TypeMap map[string]GoType // key = PG type, value = Go type

var DefaultTypeMap = TypeMap{
	"bool": GoType{
		Name:   "bool",
		Import: "",
	},
	"int2": GoType{
		Name:   "int16",
		Import: "",
	},
	"int4": GoType{
		Name:   "int32",
		Import: "",
	},
	"int8": GoType{
		Name:   "int64",
		Import: "",
	},
	"char": GoType{
		Name:   "int",
		Import: "",
	},
	"float4": GoType{
		Name:   "float32",
		Import: "",
	},
	"float8": GoType{
		Name:   "float64",
		Import: "",
	},
	"numberic": GoType{
		Name:   "decimal.Decimal",
		Import: "github.com/shopspring/decimal",
	},
	"text": GoType{
		Name:   "string",
		Import: "",
	},
	"uuid": GoType{
		Name:   "uuid.UUID",
		Import: "github.com/gofrs/uuid",
	},
}
