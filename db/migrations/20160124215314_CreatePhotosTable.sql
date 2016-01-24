
-- +goose Up
CREATE TABLE images (
	id INTEGER PRIMARY KEY,
	user_id INTEGER,

	title VARCHAR(255),
	file VARCHAR(255),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted UNSIGNED TINYINT(1) NOT NULL DEFAULT 0
);


-- +goose Down
DROP TABLE images;

