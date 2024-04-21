// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: pay_out.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPayOut = `-- name: CreatePayOut :one
INSERT INTO pay_out (owner, amount, subject) VALUES ($1, $2, $3) RETURNING id, owner, amount, subject, created_at, updated_at
`

type CreatePayOutParams struct {
	Owner   string  `json:"owner"`
	Amount  float32 `json:"amount"`
	Subject string  `json:"subject"`
}

func (q *Queries) CreatePayOut(ctx context.Context, arg CreatePayOutParams) (PayOut, error) {
	row := q.db.QueryRow(ctx, createPayOut, arg.Owner, arg.Amount, arg.Subject)
	var i PayOut
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Amount,
		&i.Subject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePayOut = `-- name: DeletePayOut :one
DELETE FROM pay_out WHERE id = $1 RETURNING id, owner, amount, subject, created_at, updated_at
`

func (q *Queries) DeletePayOut(ctx context.Context, id uuid.UUID) (PayOut, error) {
	row := q.db.QueryRow(ctx, deletePayOut, id)
	var i PayOut
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Amount,
		&i.Subject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPayOut = `-- name: GetPayOut :one
SELECT id, owner, amount, subject, created_at, updated_at FROM pay_out WHERE id = $1
`

func (q *Queries) GetPayOut(ctx context.Context, id uuid.UUID) (PayOut, error) {
	row := q.db.QueryRow(ctx, getPayOut, id)
	var i PayOut
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Amount,
		&i.Subject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPayOuts = `-- name: ListPayOuts :many
SELECT id, owner, amount, subject, created_at, updated_at FROM pay_out ORDER BY id OFFSET $1 LIMIT $2
`

type ListPayOutsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListPayOuts(ctx context.Context, arg ListPayOutsParams) ([]PayOut, error) {
	rows, err := q.db.Query(ctx, listPayOuts, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PayOut{}
	for rows.Next() {
		var i PayOut
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Amount,
			&i.Subject,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchPayOuts = `-- name: SearchPayOuts :many
SELECT id, owner, amount, subject, created_at, updated_at FROM pay_out WHERE owner ILIKE $1 ORDER BY id OFFSET $2 LIMIT $3
`

type SearchPayOutsParams struct {
	Owner  string `json:"owner"`
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
}

func (q *Queries) SearchPayOuts(ctx context.Context, arg SearchPayOutsParams) ([]PayOut, error) {
	rows, err := q.db.Query(ctx, searchPayOuts, arg.Owner, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PayOut{}
	for rows.Next() {
		var i PayOut
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Amount,
			&i.Subject,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
