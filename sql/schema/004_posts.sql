-- +goose Up
CREATE TABLE IF NOT EXISTS posts(
id uuid PRIMARY KEY, 
created_at TIMESTAMP NOT NULL, 
updated_at TIMESTAMP NOT NULL, 
title VARCHAR NOT NULL, 
url VARCHAR NOT NULL UNIQUE, 
description VARCHAR, 
published_at TIMESTAMP, 
feed_id uuid NOT NULL,
FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;