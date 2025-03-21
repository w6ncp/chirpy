-- name: CreateChip :one
INSERT INTO chirps (id, created_at, updated_at, user_id, body)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;
---

-- name: GetChirps :many
SELECT * FROM chirps ORDER BY created_at ASC;
---

-- name: GetChirpByID :one
SELECT * FROM chirps WHERE id = $1;
---

-- name: DeleteChirpByID :exec
DELETE FROM chirps
WHERE id = $1;
---

-- name: GetChirpsByUser :many
SELECT * from chirps
WHERE user_id = $1
ORDER BY created_at ASC;
---
