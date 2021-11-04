-- name: CreateCustomer :execresult
INSERT INTO customer (
  cif_code,level,type,title,full_name,mothers_maiden,gender,birth_date,email,status,created_date,created_by,last_updated_date,last_updated_by,company_code,employers_code,profile_version,cif_code_nkyc
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetCustomerById :one
SELECT * FROM customer WHERE id = ? LIMIT 1;

-- name: GetCustomerByIdForUpdate :one
SELECT * FROM customer WHERE id = ? LIMIT 1 FOR UPDATE;

-- name: ListCustomer :many
SELECT * FROM customer ORDER BY id;

-- name: DeleteCustomerById :exec
DELETE FROM customer WHERE id = ?;