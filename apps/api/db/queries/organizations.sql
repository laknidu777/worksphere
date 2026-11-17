-- name: GetOrganizationByID :one
SELECT id, name, slug, currency, timezone, locale, created_at, updated_at, deleted_at
FROM organizations
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetOrganizationBySlug :one
SELECT id, name, slug, currency, timezone, locale, created_at, updated_at, deleted_at
FROM organizations
WHERE slug = $1 AND deleted_at IS NULL;

-- name: CreateOrganization :one
INSERT INTO organizations (
    name, slug, currency, timezone, locale
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, name, slug, currency, timezone, locale, created_at, updated_at, deleted_at;

-- name: GetUserByExternalSub :one
SELECT id, external_sub, email, full_name, created_at, updated_at
FROM users
WHERE external_sub = $1;

-- name: CreateUser :one
INSERT INTO users (
    external_sub, email, full_name
) VALUES (
    $1, $2, $3
) RETURNING id, external_sub, email, full_name, created_at, updated_at;
