-- name: GetUserById :one
SELECT *
FROM "User"
WHERE "id" = $1;

-- name: CreateUser :one
INSERT INTO "User" ("id", "name", "profilePic")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ChangeUserProperties :one
UPDATE "User"
SET "name" = $2,
    "profilePic" = $3
WHERE "id" = $1
RETURNING *;

-- name: ChangeUserName :one
UPDATE "User"
SET "name" = $2
WHERE "id" = $1
RETURNING *;

-- name: ChangeUserProfilePic :one
UPDATE "User"
SET "profilePic" = $2
WHERE "id" = $1
RETURNING *;