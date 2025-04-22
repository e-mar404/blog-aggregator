-- +goose Up
CREATE TABLE posts (
  id UUID PRIMARY KEY,
  title TEXT,
  url TEXT UNIQUE,
  description TEXT,
  published_at TIMESTAMP,
  feed_id UUID REFERENCES feeds(id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts CASCADE;
