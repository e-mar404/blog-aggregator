-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE,
  url TEXT UNIQUE,
  user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  feed_id UUID REFERENCES feeds(id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT unique_follow UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feeds;
DROP TABLE feed_follows;
