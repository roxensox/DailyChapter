-- name: Subscribe :exec
INSERT INTO userbooks (
	id,
	created_at,
	updated_at,
	user_id,
	book_id
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
);
