
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

