-- name: CreateUser :one
INSERT INTO uaccount (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM uaccount
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM uaccount
WHERE email = $1
LIMIT 1;

--name: ListUsers :many
SELECT * FROM uaccount
LIMIT $1
OFFSET $2;

-- name: UpdateFirstName :exec
UPDATE uaccount
SET first_name = $2
WHERE id = $1;

-- name: UpdateLastName :exec
UPDATE uaccount
SET last_name = $2
WHERE id = $1;

-- name: UpdateEmail :exec
UPDATE uaccount
SET email = $2
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE uaccount
SET password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM uaccount
WHERE id = $1;
