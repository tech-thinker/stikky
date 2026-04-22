-- name: GetStikky :one
SELECT * FROM stikky ORDER BY id ASC LIMIT 1;

-- name: GetPublicKeys :many
SELECT * FROM public_keys;
