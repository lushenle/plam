-- name: CreateLoan :one
INSERT INTO loan (borrower, amount, subject) VALUES ($1, $2, $3) RETURNING *;

-- name: ListLoans :many
SELECT * FROM loan ORDER BY id OFFSET $1 LIMIT $2;

-- name: GetLoan :one
SELECT * FROM loan WHERE id = $1;

-- name: SearchLoans :many
SELECT * FROM loan WHERE borrower ILIKE $1 ORDER BY id OFFSET $2 LIMIT $3;

-- name: DeleteLoan :one
DELETE FROM loan WHERE id = $1 RETURNING *;