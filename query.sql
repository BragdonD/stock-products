-- name: ListProducts :many
SELECT * FROM "Products"
ORDER BY "updated-at" DESC
LIMIT $2 OFFSET (($1 - 1) * $2);
