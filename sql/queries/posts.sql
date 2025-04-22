-- name: CreatePost :one
INSERT INTO posts (id, title, url, description, published_at, feed_id, created_at, updated_at)
Values (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT *, feeds.name AS feed_name FROM posts
LEFT JOIN feeds
ON feeds.id=posts.feed_id
JOIN feed_follows 
ON feed_follows.feed_id=posts.feed_id AND feed_follows.user_id=$1
ORDER BY posts.published_at DESC
LIMIT $2;
