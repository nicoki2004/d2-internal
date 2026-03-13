-- name: GetUser :one
SELECT * FROM users
WHERE membership_id = ? LIMIT 1;

-- name: UpsertUser :exec
INSERT INTO users (
    membership_id, display_name, access_token, refresh_token, 
    expires_in, refresh_expires_in, received_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(membership_id) DO UPDATE SET
    access_token = excluded.access_token,
    refresh_token = excluded.refresh_token,
    expires_in = excluded.expires_in,
    refresh_expires_in = excluded.refresh_expires_in,
    received_at = excluded.received_at,
    updated_at = excluded.updated_at;
