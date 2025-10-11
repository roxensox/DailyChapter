-- +goose Up
CREATE TABLE refreshtokens (
	id UUID PRIMARY KEY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
	token TEXT NOT NULL,
	valid_until TIMESTAMP NOT NULL,
	revoked_at TIMESTAMP
);

-- +goose Down
DROP TABLE refreshtokens;
