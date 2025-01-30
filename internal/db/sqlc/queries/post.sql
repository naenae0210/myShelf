-- name: CreatePost :exec
INSERT INTO posts (user_id, title, content)
VALUES (?, ?, ?);

-- name: GetPostByID :one
SELECT * FROM posts
WHERE id = ?;

-- name: UpdatePost :exec
UPDATE posts
SET content = ?
WHERE id = ?;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?