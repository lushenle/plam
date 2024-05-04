// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: loan.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createLoan = `-- name: CreateLoan :one
INSERT INTO loan (borrower, amount, subject) VALUES ($1, $2, $3) RETURNING id, borrower, amount, subject, created_at, updated_at
`

type CreateLoanParams struct {
	Borrower string  `json:"borrower"`
	Amount   float32 `json:"amount"`
	Subject  string  `json:"subject"`
}

func (q *Queries) CreateLoan(ctx context.Context, arg CreateLoanParams) (Loan, error) {
	row := q.db.QueryRow(ctx, createLoan, arg.Borrower, arg.Amount, arg.Subject)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.Borrower,
		&i.Amount,
		&i.Subject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteLoan = `-- name: DeleteLoan :one
DELETE FROM loan WHERE id = $1 RETURNING id, borrower, amount, subject, created_at, updated_at
`

func (q *Queries) DeleteLoan(ctx context.Context, id uuid.UUID) (Loan, error) {
	row := q.db.QueryRow(ctx, deleteLoan, id)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.Borrower,
		&i.Amount,
		&i.Subject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLoan = `-- name: GetLoan :one
SELECT id, borrower, amount, subject, created_at, updated_at FROM loan WHERE id = $1
`

func (q *Queries) GetLoan(ctx context.Context, id uuid.UUID) (Loan, error) {
	row := q.db.QueryRow(ctx, getLoan, id)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.Borrower,
		&i.Amount,
		&i.Subject,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listLoans = `-- name: ListLoans :many
SELECT id, borrower, amount, subject, created_at, updated_at FROM loan ORDER BY id OFFSET $1 LIMIT $2
`

type ListLoansParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListLoans(ctx context.Context, arg ListLoansParams) ([]Loan, error) {
	rows, err := q.db.Query(ctx, listLoans, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Loan{}
	for rows.Next() {
		var i Loan
		if err := rows.Scan(
			&i.ID,
			&i.Borrower,
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

const searchLoans = `-- name: SearchLoans :many
SELECT id, borrower, amount, subject, created_at, updated_at FROM loan WHERE borrower ILIKE $1 ORDER BY id OFFSET $2 LIMIT $3
`

type SearchLoansParams struct {
	Borrower string `json:"borrower"`
	Offset   int32  `json:"offset"`
	Limit    int32  `json:"limit"`
}

func (q *Queries) SearchLoans(ctx context.Context, arg SearchLoansParams) ([]Loan, error) {
	rows, err := q.db.Query(ctx, searchLoans, arg.Borrower, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Loan{}
	for rows.Next() {
		var i Loan
		if err := rows.Scan(
			&i.ID,
			&i.Borrower,
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
