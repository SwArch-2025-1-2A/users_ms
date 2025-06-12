-- name: CreateCategory :one
INSERT INTO "Category"("category")
VALUES ($1)
RETURNING *;

-- name: SoftDeleteCategory :one
UPDATE "Category"
SET "deleted_at" = now()
WHERE "id" = $1
RETURNING *;

-- name: GetCategoryById :one
SELECT * FROM "Category"
WHERE "deleted_at" IS NULL AND "id" = $1;

-- name: GetCategoriesById :many
SELECT * FROM "Category"
WHERE "deleted_at" IS NULL
ORDER BY "category";
