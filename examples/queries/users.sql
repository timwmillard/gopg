
-- name: GetUsers :many
-- 
-- @validation isValidUser
-- select *
-- from users;


-- name: GetUser :one
-- @param $1 id uuid.UUID
-- 
-- @return User
select * from users
where id = $1;


-- name: UpdateDebtor
-- @template: "pgx/one"
-- @param $1 = debtor_id uuid.UUID
-- @param $2 = firm_id uuid.UUID
-- @param {$3, $4, $5, $6} = debtor model.Debtor
-- @param $7 = address model.Address 
-- @return model.Debtor
UPDATE debtor SET 
    first_name = $3,
    last_name = $4,
    email = $5,
    contact_number = $6,
    address = $7,
    last_updated_at = now()
WHERE debtor_id = $1
    AND company_id = $2
RETURNING *;

