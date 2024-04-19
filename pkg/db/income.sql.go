// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: income.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createIncome = `-- name: CreateIncome :one
INSERT INTO income (payee, amount, project_id) VALUES ($1, $2, $3) RETURNING id, payee, amount, project_id, created_at, updated_at
`

type CreateIncomeParams struct {
	Payee     string    `json:"payee"`
	Amount    float32   `json:"amount"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) CreateIncome(ctx context.Context, arg CreateIncomeParams) (Income, error) {
	row := q.db.QueryRow(ctx, createIncome, arg.Payee, arg.Amount, arg.ProjectID)
	var i Income
	err := row.Scan(
		&i.ID,
		&i.Payee,
		&i.Amount,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteIncome = `-- name: DeleteIncome :one
DELETE FROM income WHERE id = $1 RETURNING id, payee, amount, project_id, created_at, updated_at
`

func (q *Queries) DeleteIncome(ctx context.Context, id string) (Income, error) {
	row := q.db.QueryRow(ctx, deleteIncome, id)
	var i Income
	err := row.Scan(
		&i.ID,
		&i.Payee,
		&i.Amount,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getIncome = `-- name: GetIncome :one
SELECT id, payee, amount, project_id, created_at, updated_at FROM income WHERE id = $1
`

func (q *Queries) GetIncome(ctx context.Context, id string) (Income, error) {
	row := q.db.QueryRow(ctx, getIncome, id)
	var i Income
	err := row.Scan(
		&i.ID,
		&i.Payee,
		&i.Amount,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listIncomes = `-- name: ListIncomes :many
SELECT id, payee, amount, project_id, created_at, updated_at FROM income ORDER BY id OFFSET $1 LIMIT $2
`

type ListIncomesParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListIncomes(ctx context.Context, arg ListIncomesParams) ([]Income, error) {
	rows, err := q.db.Query(ctx, listIncomes, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Income{}
	for rows.Next() {
		var i Income
		if err := rows.Scan(
			&i.ID,
			&i.Payee,
			&i.Amount,
			&i.ProjectID,
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

const searchIncomes = `-- name: SearchIncomes :many
SELECT id, payee, amount, project_id, created_at, updated_at FROM income WHERE payee ILIKE '%' || $1 || '%' ORDER BY id OFFSET $2 LIMIT $3
`

type SearchIncomesParams struct {
	Column1 pgtype.Text `json:"column_1"`
	Offset  int32       `json:"offset"`
	Limit   int32       `json:"limit"`
}

func (q *Queries) SearchIncomes(ctx context.Context, arg SearchIncomesParams) ([]Income, error) {
	rows, err := q.db.Query(ctx, searchIncomes, arg.Column1, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Income{}
	for rows.Next() {
		var i Income
		if err := rows.Scan(
			&i.ID,
			&i.Payee,
			&i.Amount,
			&i.ProjectID,
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
