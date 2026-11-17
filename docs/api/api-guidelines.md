# WorkSphere API Guidelines & Contracts

## 1. REST API Conventions

All external and internal HTTP APIs exposed by `apps/api` adhere to standardized REST principles:

1. **Base URL Structure**:
   - Health and operational endpoints: `/health`, `/ready`, `/live`, `/metrics`
   - Public/Auth endpoints: `/api/v1/auth/*`
   - Tenant-scoped business endpoints: `/api/v1/orgs/{orgId}/*`
   - Platform super-admin endpoints: `/api/v1/admin/*`
2. **HTTP Verbs**:
   - `GET`: Safe, idempotent read operations.
   - `POST`: Non-idempotent resource creation or command execution.
   - `PUT`: Idempotent full replacement of an existing resource.
   - `PATCH`: Idempotent partial update of an existing resource.
   - `DELETE`: Resource deletion or soft-deletion.
3. **HTTP Status Codes**:
   - `200 OK`: Request succeeded with payload.
   - `201 Created`: Resource successfully created. Returns `Location` header or created entity.
   - `204 No Content`: Successful request with no response body.
   - `400 Bad Request`: Validation failure or malformed payload.
   - `401 Unauthorized`: Missing or invalid authentication token.
   - `403 Forbidden`: Authenticated user lacks permission or organization access.
   - `404 Not Found`: Resource does not exist or tenant is not authorized to know of its existence.
   - `409 Conflict`: Business state conflict (e.g. unique constraint, concurrent modification).
   - `422 Unprocessable Entity`: Semantic validation failure.
   - `500 Internal Server Error`: Unhandled server exception (sanitized in production).

---

## 2. Standardized Error Response Format

Errors adhere strictly to RFC 7807 (Problem Details for HTTP APIs):

```json
{
  "type": "https://api.worksphere.lk/errors/validation-failed",
  "title": "Validation Failed",
  "status": 422,
  "detail": "The payload contains one or more invalid fields.",
  "instance": "/api/v1/orgs/01933a5e-38c2-7472-8f92-ec64152e259b/invoices",
  "invalid_params": [
    {
      "name": "amount",
      "reason": "Amount must be a positive number formatted as LKR currency."
    }
  ]
}
```

---

## 3. OpenAPI 3.1 & Contract-First Development

API contracts are maintained in `packages/api-contracts/openapi`.
- Contracts define request schemas, response schemas, security definitions, and example payloads.
- Frontend TypeScript types and API query hooks are generated from the OpenAPI schemas using `openapi-typescript` / TanStack Query tools.
- Go backend handlers validate incoming JSON bodies against schemas using strict parsing.
