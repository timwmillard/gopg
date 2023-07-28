package gen_test

import (
	"sqlgen/gen"
	"testing"
)

const sql = `
UPDATE debtor SET
    first_name = $3,
    last_name = $4,
    email = $5,
    contact_number = $6,
    address_line1 = $7,
    address_state = $8,
    last_updated_at = now()
WHERE debtor_id = $1
    AND company_id = $2
RETURNING *;
`

func TestGen(t *testing.T) {

	debtorID := gen.CodeVar{
		Name: "debtorID",
		Type: gen.CodeType{
			Name:   "uuid.UUID",
			Import: "github.com/gofrs/uuid",
		},
	}
	firmID := gen.CodeVar{
		Name: "firmID",
		Type: gen.CodeType{
			Name:   "uuid.UUID",
			Import: "github.com/gofrs/uuid",
		},
	}

	debtor := gen.CodeVar{
		Name: "debtor",
		Type: gen.CodeType{
			Name:   "model.Debtor",
			Import: "github.com/sqlgen/gen/model",
		},
		Fields: []gen.CodeVar{},
	}
	debtor.AddField(gen.CodeVar{
		Name: "FirstName",
		Type: gen.CodeType{
			Name:   "string",
			Import: "",
		},
	})
	debtor.AddField(gen.CodeVar{
		Name: "LastName",
		Type: gen.CodeType{
			Name:   "string",
			Import: "",
		},
	})

	function := gen.Function{
		Name:    "UpdaetDebtor",
		Comment: "Update the debtor details.",
		SQL:     sql,
		Imports: []string{
			"github.com/sqlgen/gen/model",
			"github.com/gofrs/uuid",
		},
		Args: []gen.CodeVar{
			firmID,
			debtorID,
			debtor,
		},
		Input: []gen.Param{
			{DBField: "$1", CodeVar: firmID},
			{DBField: "$2", CodeVar: debtorID},
			{DBField: "$3", CodeVar: debtor.Fields[0]},
			{DBField: "$4", CodeVar: debtor.Fields[1]},
		},
		Output: []gen.Param{
			{DBField: "first_name", CodeVar: debtor.Fields[0]},
			{DBField: "last_name", CodeVar: debtor.Fields[1]},
		},
		Return: []gen.CodeVar{
			debtor,
		},
	}

	err := gen.Run(function)
	if err != nil {
		t.Error(err)
	}
}
