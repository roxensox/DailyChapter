-- name: CreateUser :one
INSERT INTO users (
	id,
	email,
	hashed_password,
	created_at,
	updated_at
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
) RETURNING id, created_at, email;

-- name: GetUserByEmail :one
SELECT id, email, hashed_password, updated_at
FROM users
WHERE email = $1;

-- name: Reset :exec
DELETE FROM users;
