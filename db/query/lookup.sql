-- name: CreateLookup :execresult
INSERT INTO lookup (
  type,code,name,priority,description,shortname,status,filter
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetLookupById :one
SELECT * FROM lookup WHERE id = ? LIMIT 1;

-- name: ListLookup :many
SELECT * FROM lookup ORDER BY id;

-- name: DeleteLookupById :exec
DELETE FROM lookup WHERE id = ?;