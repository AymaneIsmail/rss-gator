-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUsers :many
SELECT *
FROM users 
;

-- name: GetOneUserByName :one
SELECT * 
FROM USERS
WHERE name = $1;

-- name: DropUsers :exec
DELETE FROM users;