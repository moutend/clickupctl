-- name: GetResponse :one
SELECT * FROM responses WHERE path = ? LIMIT 1;

-- name: CreateResponse :one
INSERT INTO responses (path, body, cached_at, expired_at) VALUES (?, ?, ?, ?) RETURNING *;

-- name: UpdateResponse :one
UPDATE responses SET body = ?, cached_at = ?, expired_at = ? WHERE path = ? RETURNING *;
