-- +goose Up
CREATE TABLE books (
	id UUID PRIMARY KEY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title TEXT NOT NULL,
	pub_date TIMESTAMP
);

-- +goose Down
DROP TABLE books;
