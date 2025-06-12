-- name: AddUserInterest :one
INSERT INTO "UserInterests"("user_id", "interest_id")
VALUES ($1, $2)
RETURNING *;

-- name: RemoveUserInterest :exec
DELETE FROM "UserInterests"
WHERE "user_id" = $1
  AND "interest_id" = $2;

-- name: GetUserInterests :many
SELECT c.*
FROM "UserInterests" as ui
  JOIN "Category" as c
  ON ui.interest_id = c.id
WHERE ui.user_id = $1
  AND c.deleted_at IS NULL;
