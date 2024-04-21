-- name: CreatePayOut :one
INSERT INTO pay_out (owner, amount, subject) VALUES ($1, $2, $3) RETURNING *;

-- name: ListPayOuts :many
SELECT * FROM pay_out ORDER BY id OFFSET $1 LIMIT $2;

-- name: GetPayOut :one
SELECT * FROM pay_out WHERE id = $1;

-- name: SearchPayOuts :many
SELECT * FROM pay_out WHERE owner ILIKE $1 ORDER BY id OFFSET $2 LIMIT $3;

-- name: DeletePayOut :one
DELETE FROM pay_out WHERE id = $1 RETURNING *;
