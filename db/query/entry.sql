-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 limit 1;