// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: project.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createProject = `-- name: CreateProject :one
INSERT INTO project (name, description, amount) VALUES ($1, $2, $3) RETURNING id, name, amount, description, created_at, updated_at
`

type CreateProjectParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject, arg.Name, arg.Description, arg.Amount)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Amount,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :one
DELETE FROM project WHERE id = $1 RETURNING id, name, amount, description, created_at, updated_at
`

func (q *Queries) DeleteProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRow(ctx, deleteProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Amount,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProject = `-- name: GetProject :one
SELECT id, name, amount, description, created_at, updated_at FROM project WHERE id = $1
`

func (q *Queries) GetProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRow(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Amount,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, amount, description, created_at, updated_at FROM project ORDER BY id OFFSET $1 LIMIT $2
`

type ListProjectsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListProjects(ctx context.Context, arg ListProjectsParams) ([]Project, error) {
	rows, err := q.db.Query(ctx, listProjects, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Amount,
			&i.Description,
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

const searchProjects = `-- name: SearchProjects :many
SELECT id, name, amount, description, created_at, updated_at FROM project WHERE name ILIKE $1 ORDER BY id OFFSET $2 LIMIT $3
`

type SearchProjectsParams struct {
	Name   string `json:"name"`
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
}

func (q *Queries) SearchProjects(ctx context.Context, arg SearchProjectsParams) ([]Project, error) {
	rows, err := q.db.Query(ctx, searchProjects, arg.Name, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Amount,
			&i.Description,
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
