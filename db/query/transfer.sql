-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id, to_account_id, amount
    )
VALUES (
    $1, $2, $3
    )RETURNING *;

-- name: UpdateTransfer :one
UPDATE transfers
SET from_account_id = $2, to_account_id = $3, amount = $4
WHERE id = $1
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1;

-- name: ListTransfer :many
SELECT * FROM transfers
ORDER BY id DESC;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
