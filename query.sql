-- name: CreateProduct :one
INSERT INTO "Products" 
("name", "reference", "line", "supplier", "description", "price", "image-url") 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetProduct :one
SELECT * FROM "Products"
WHERE "name" = $1;

-- name: ListProducts :many
SELECT * FROM "Products"
ORDER BY "updated-at" DESC
LIMIT $2 OFFSET (($1 - 1) * $2);