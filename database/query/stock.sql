-- name: GetStock :one
SELECT * FROM stocks
WHERE ID = $1 LIMIT 1;

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
SET Name = $2,
        Price = $3,
        Company = $4,
        UpdatedAt = CURRENT_TIMESTAMP
WHERE ID = $1
RETURNING *;

-- name: DeleteStock :exec
DELETE FROM stocks
WHERE ID = $1;