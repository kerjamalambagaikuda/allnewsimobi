-- name: CreateStatus :execresult
INSERT INTO status (
  type, code, name, description
) VALUES (
  ?, ?, ?, ?
);

-- name: GetStatusById :one
SELECT * FROM status WHERE id = ? LIMIT 1;

-- name: ListStatus :many
SELECT * FROM status ORDER BY id;

/* name: DeleteStatusById :exec */
DELETE FROM status WHERE id = ?;