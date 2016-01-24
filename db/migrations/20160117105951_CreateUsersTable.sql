
-- +goose Up
CREATE TABLE users (
	id INTEGER PRIMARY KEY,

	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(100),
	password CHAR(60),
	status UNSIGNED TINYINT(1) NOT NULL DEFAULT 1,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted UNSIGNED TINYINT(1) NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE users;
