-- name: AddUserInterest :one
INSERT INTO "UserInterests"("user_id", "interest_id")
VALUES ($1, $2)
RETURNING *;

-- name: RemoveUserInterest :exec
DELETE FROM "UserInterests"
WHERE "user_id" = $1
  AND "interest_id" = $2;