// Code generated by sqlc. DO NOT EDIT.
// source: status.sql

package db

import (
	"context"
	"database/sql"
)

const createStatus = `-- name: CreateStatus :execresult
INSERT INTO status (
  type, code, name, description
) VALUES (
  ?, ?, ?, ?
)
`

type CreateStatusParams struct {
	Type        string         `json:"type"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateStatus(ctx context.Context, arg CreateStatusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createStatus,
		arg.Type,
		arg.Code,
		arg.Name,
		arg.Description,
	)
}

const deleteStatusById = `-- name: DeleteStatusById :exec
DELETE FROM status WHERE id = ?
`

func (q *Queries) DeleteStatusById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStatusById, id)
	return err
}

const getStatusById = `-- name: GetStatusById :one
SELECT id, type, code, name, description FROM status WHERE id = ? LIMIT 1
`

func (q *Queries) GetStatusById(ctx context.Context, id int64) (Status, error) {
	row := q.db.QueryRowContext(ctx, getStatusById, id)
	var i Status
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Code,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const listStatus = `-- name: ListStatus :many
SELECT id, type, code, name, description FROM status ORDER BY id
`

func (q *Queries) ListStatus(ctx context.Context) ([]Status, error) {
	rows, err := q.db.QueryContext(ctx, listStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Status
	for rows.Next() {
		var i Status
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Code,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
