-- name: GetStock :one
SELECT * FROM stocks
WHERE ID = $1 LIMIT 1;

-- name: GetStockByName :one
SELECT * FROM stocks WHERE name = $1 LIMIT 1;

-- name: ListStocks :many
SELECT * FROM stocks
ORDER BY Name;

-- name: CreateStock :one
INSERT INTO stocks (
    Name, Price, Company
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateStock :exec
UPDATE stocks
SET 
    name = COALESCE($2, name),
    price = COALESCE($3, price),
    company = COALESCE($4, company),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteStock :exec
DELETE FROM stocks
WHERE ID = $1;