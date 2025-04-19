-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE,
  url TEXT UNIQUE,
  user_id UUID UNIQUE REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feeds;
