# ADR-004: Amazon Cognito OIDC and Server-Managed Session Architecture

## Status
**Proposed / Baseline Validation** (2026-11-17)

## Context
WorkSphere requires robust identity management, multi-factor authentication (MFA), password reset flows, and compliance with data security standards without developing and maintaining custom cryptographic password management systems.

## Decision
We select **Amazon Cognito User Pools** as our proposed Identity Provider (IdP) operating through OpenID Connect (OIDC) and OAuth 2.0 Authorization Code Flow with PKCE (S256):
- **Authentication**: Cognito holds user credentials, handles authentication challenges, MFA, and returns JWT tokens (`id_token`, `access_token`, `refresh_token`).
- **Next.js BFF (Backend-for-Frontend)**: Next.js Route Handlers act as the OAuth client. Tokens are kept securely on the server side and exchanged for an encrypted, signed HTTP-only session cookie.
- **Go API Independence**: The Go API receives verified bearer tokens or signed session assertions, validating Cognito's RS256 JWT public keys independently without calling Cognito on every request.
- **Identity Decoupling**: WorkSphere's database does not store user passwords. Instead, it maps Cognito's immutable `sub` claim to internal `user_id` records. All organization memberships, roles, and granular permissions are managed entirely within WorkSphere's PostgreSQL database.
- **Local Dev Resilience**: For local development without AWS, the authentication layer supports a mock/development auth provider returning synthetic claims.

## Consequences
- **Positive**: No credential liability, built-in MFA/security features, compliant session handling.
- **Validation Needed**: Latency and operational ergonomics of Cognito user pools in AWS regions closest to Sri Lanka (e.g. `ap-south-1` Mumbai / `ap-southeast-1` Singapore) must be verified during staging rollout.
