-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (
    id,
    created_at,
    updated_at,
    feed_id,
    user_id
  )
  VALUES (
    $1,  -- id
    $2,  -- created_at
    $3,  -- updated_at
    $4,  -- feed_id
    $5   -- user_id
  )
  RETURNING *
)
SELECT
  iff.*,
  f.name AS feed_name,
  u.name AS user_name
FROM inserted_feed_follow AS iff
INNER JOIN feeds AS f ON f.id = iff.feed_id
INNER JOIN users AS u ON u.id = iff.user_id;

-- name: GetFeedFollowsForUser :many
SELECT
  ff.id,
  ff.created_at,
  ff.updated_at,
  ff.feed_id,
  ff.user_id,
  f.name AS feed_name
FROM feed_follows AS ff
INNER JOIN feeds AS f ON f.id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY ff.created_at DESC;

-- name: UnfollowFeed :one
DELETE FROM feed_follows
WHERE user_id = $1
AND feed_id = $2
RETURNING id, created_at, updated_at, feed_id, user_id;