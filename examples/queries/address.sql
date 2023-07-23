
-- name: GetUserAddress :one
select * from address where user_id = $1;
