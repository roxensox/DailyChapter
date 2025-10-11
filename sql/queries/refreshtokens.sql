-- name: CheckRefreshToken :one
SELECT user_id 
FROM refreshtokens
WHERE 
	token = $1 AND 
	revoked_at = NULL AND
	valid_until > NOW()
;

-- name: CreateRefreshToken :one
INSERT INTO refreshtokens (
	id,
	created_at,
	updated_at,
	token,
	user_id,
	valid_until,
	revoked_at
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
) RETURNING token;

-- name: ExpireRefreshToken :exec
UPDATE refreshtokens
SET	revoked_at = NOW(),
	updated_at = NOW()
WHERE token = $1;
