-- +goose Up
CREATE TABLE authors (
	id UUID PRIMARY KEY NOT NULL,
	created_at Timestamp NOT NULL,
	updated_at Timestamp NOT NULL,
	name TEXT NOT NULL
);

-- +goose Down
DROP TABLE authors;
