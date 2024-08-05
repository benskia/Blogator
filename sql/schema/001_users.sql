-- +goose up
CREATE TABLE users (
	id UUID PRIMARY KEY,
	created_at temporal NOT NULL,
	updated_at temporal NOT NULL,
	name character NOT NULL,
);

-- +goose down
DROP TABLE users;
