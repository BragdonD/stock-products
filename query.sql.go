// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package stockproducts

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO "Products" 
("name", "reference", "line", "supplier", "description", "price", "image-url") 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING uuid, name, reference, line, supplier, description, "image-url", price, "created-at", "updated-at"
`

type CreateProductParams struct {
	Name        string
	Reference   string
	Line        string
	Supplier    string
	Description string
	Price       pgtype.Numeric
	ImageUrl    string
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Name,
		arg.Reference,
		arg.Line,
		arg.Supplier,
		arg.Description,
		arg.Price,
		arg.ImageUrl,
	)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.Reference,
		&i.Line,
		&i.Supplier,
		&i.Description,
		&i.ImageUrl,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT uuid, name, reference, line, supplier, description, "image-url", price, "created-at", "updated-at" FROM "Products"
WHERE "name" = $1
`

func (q *Queries) GetProduct(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, name)
	var i Product
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.Reference,
		&i.Line,
		&i.Supplier,
		&i.Description,
		&i.ImageUrl,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT uuid, name, reference, line, supplier, description, "image-url", price, "created-at", "updated-at" FROM "Products"
ORDER BY "updated-at" DESC
LIMIT $2 OFFSET (($1 - 1) * $2)
`

type ListProductsParams struct {
	Column1 interface{}
	Limit   int32
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts, arg.Column1, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.Reference,
			&i.Line,
			&i.Supplier,
			&i.Description,
			&i.ImageUrl,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
