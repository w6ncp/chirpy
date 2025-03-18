-- +goose Up
CREATE TABLE chirps (
    id UUID PRIMARY KEY,
    body TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id UUID NOT NULL
    REFERENCES users (id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE chirps;
