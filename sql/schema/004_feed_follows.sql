-- +goose Up
CREATE TABLE feed_follows (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL,
	feed_id UUID NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	FOREIGN KEY(user_id)
		REFERENCES users(id)
		ON DELETE CASCADE,
	FOREIGN KEY(feed_id)
		REFERENCES feeds(id)
		ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_follows;

