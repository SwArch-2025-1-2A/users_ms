-- name: GetUserById :one
SELECT *
FROM "User"
WHERE "id" = $1;

-- name: CreateUser :one
INSERT INTO "User" ("id", "name", "profile_pic")
VALUES ($1, $2, $3)
RETURNING *;

-- name: ChangeUserProperties :one
UPDATE "User"
SET "name" = $2,
    "profile_pic" = $3
WHERE "id" = $1
RETURNING *;

-- name: ChangeUserName :one
UPDATE "User"
SET "name" = $2
WHERE "id" = $1
RETURNING *;

-- name: ChangeUserProfilePic :one
UPDATE "User"
SET "profile_pic" = $2
WHERE "id" = $1
RETURNING *;