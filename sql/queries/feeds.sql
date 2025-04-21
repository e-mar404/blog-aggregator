-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name AS feed_name, feeds.url AS feed_url, users.name AS user_name FROM feeds
JOIN users
ON feeds.user_id=users.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE feeds.url=$1;

-- name: CreateFeedFollow :one
WITH follow_feed_record AS (
  INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
  VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
  )
  RETURNING *
)
SELECT *, users.name AS user_name, feeds.name AS feed_name
FROM follow_feed_record
JOIN users
ON follow_feed_record.user_id=users.id
JOIN feeds 
ON follow_feed_record.feed_id=feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT *, feeds.name AS feed_name FROM feed_follows
JOIN feeds
ON feed_follows.feed_id=feeds.id
WHERE feed_follows.user_id=$1;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows
USING users, feeds
WHERE feed_follows.user_id=$1 AND feed_follows.feed_id=$2;
