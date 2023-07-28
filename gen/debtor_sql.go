package gen

import (
	"context"
	"database/sql"
	"sqlgen/gen/model"

	"github.com/gofrs/uuid"
)

// -- name: UpdateDebtor :one
// -- Update the debtor details
// -- @param $1 = debtorID uuid.UUID
// -- @param $2 = firmID uuid.UUID
// -- @param {$3, $4, $5, $6} = debtor model.Debtor
// --
// -- @param $7 = address model.Address
// -- @return model.Debtor, model.Address
// --
// -- @import github.com/sqlgen/gen/model
// -- @import github.com/gofrs/uuid
// -- @type uuid = uuid.UUID
// --
// UPDATE debtor SET
//     first_name = $3,
//     last_name = $4,
//     email = $5,
//     contact_number = $6,
//     address_line1 = $7,
//     address_state = $8,
//     last_updated_at = now()
// WHERE debtor_id = $1
//     AND company_id = $2
// RETURNING *;

const updateDebtorSQL = `
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

// UpdateDebtor - Update the debtor details
func UpdateDebtor(ctx context.Context, db *sql.DB,
	debtorID uuid.UUID,
	firmID uuid.UUID,
	debtor model.Debtor,
	address model.Address,
) (model.Debtor, model.Address, error) {
	// input
	row := db.QueryRowContext(ctx, updateDebtorSQL,
		debtorID,             // $1
		firmID,               // $2
		debtor.FirstName,     // $3
		debtor.LastName,      // $4
		debtor.Email,         // $5
		debtor.ContactNumber, // $6
		address.Line1,        // $7
		address.Suburb,       // $8
	)
	// output
	var (
		retDebtor  model.Debtor
		retAddress model.Address
		err        error
	)
	err = row.Scan(
		&retDebtor.FirstName,     // first_name
		&retDebtor.LastName,      // last_name
		&retDebtor.Email,         // email
		&retDebtor.ContactNumber, // contact_number
		&retAddress.Line1,        // address_line1
		&retAddress.Suburb,       // address_state
	)
	return retDebtor, retAddress, err
}
