-- name: CreateIncome :one
INSERT INTO income (payee, amount, project_id) VALUES ($1, $2, $3) RETURNING *;

-- name: ListIncomes :many
SELECT * FROM income ORDER BY id OFFSET $1 LIMIT $2;

-- name: GetIncome :one
SELECT * FROM income WHERE id = $1;

-- name: SearchIncomes :many
SELECT * FROM income WHERE payee ILIKE '%' || $1 || '%' ORDER BY id OFFSET $2 LIMIT $3;

-- name: DeleteIncome :one
DELETE FROM income WHERE id = $1 RETURNING *;