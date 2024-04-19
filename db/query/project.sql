-- name: CreateProject :one
INSERT INTO project (name, description, amount) VALUES ($1, $2, $3) RETURNING *;

-- name: ListProjects :many
SELECT * FROM project ORDER BY id OFFSET $1 LIMIT $2;

-- name: GetProject :one
SELECT * FROM project WHERE id = $1;

-- name: SearchProjects :many
SELECT * FROM project WHERE name ILIKE '%' || $1 || '%' ORDER BY id OFFSET $2 LIMIT $3;

-- name: DeleteProject :one
DELETE FROM project WHERE id = $1 RETURNING *;