-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: GetUserFeeds :many
SELECT 
    feeds.id AS id,
    feeds.created_at AS created_at,
    feeds.updated_at AS updated_at,
    feeds.name AS name,
    feeds.url AS url,
    feeds.user_id AS feed_user_id,
    users.id AS user_id,
    users.created_at AS user_created_at,
    users.updated_at AS user_updated_at,
    users.name AS user_name
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT *
FROM feeds 
WHERE url = $1;