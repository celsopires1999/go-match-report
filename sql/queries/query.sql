-- name: CreateTenant :exec
INSERT INTO tenants (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4);

-- name: UpdateTenant :exec
UPDATE tenants SET name = $1, updated_at = $2 WHERE id = $3;

-- name: DeleteTenant :exec
DELETE FROM tenants WHERE id = $1;

-- name: FindTenantById :one
SELECT * FROM tenants WHERE id = $1 LIMIT 1;
