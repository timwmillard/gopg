package postgres

import (
	"sqlgen/gen"
	"strings"

	pgoid "github.com/lib/pq/oid"
)

type Type struct {
	Name string
	OID  uint32
}

func NewType(oid uint32) Type {
	name := pgoid.TypeName[pgoid.Oid(oid)]
	return Type{
		Name: strings.ToLower(name),
		OID:  oid,
	}
}

func (t *Type) IsArray() bool {
	return t.Name[0:1] == "_"
}

type Attr struct {
	Name     string // $1 or :user_id / id or name
	Type     Type   // text, int4 or uuid
	Nullable bool
}

type Query struct {
	SQL        string
	ParamAttrs []Attr
	FieldAttrs []Attr
}

func (q *Query) Params() []gen.DBField {
	return nil
}
func (q *Query) Fields() []gen.DBField {
	return nil
}

// type Query interface {
// 	Params() []DBField
// 	Fields() []DBField
// }
