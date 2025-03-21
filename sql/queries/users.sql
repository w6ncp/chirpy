-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;
---

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
---

-- name: UpdateUserByID :one
UPDATE users SET 
    (email, hashed_password, updated_at) = 
    ($1, $2, NOW())
WHERE id = $3
RETURNING *;
---

-- name: UpgradeByUser :one
UPDATE users SET
    is_chirpy_red = true,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
---

-- name: DowngradeByUser :one
UPDATE users SET
    is_chirpy_red = false,
    updated_at = NOW()
WHERE id = $1
RETURNING *;
---