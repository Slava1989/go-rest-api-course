# Go REST API Course - Comment Service

This project is a minimal REST API for creating, retrieving, updating, and deleting
comments. It is organized into clear layers (HTTP transport, service, and database)
to show a typical backend flow.

## High-level flow

1. **Server startup** (`cmd/server/main.go`)
   - Connects to Postgres using environment variables.
   - Runs database migrations from `migrations/`.
   - Builds the comment service and HTTP handler.
   - Starts the HTTP server on `0.0.0.0:8080`.

2. **HTTP request lifecycle** (`cmd/internal/transport/http/`)
   - Gorilla Mux routes map each endpoint to a handler function.
   - Middleware runs for every request:
     - `JSONMiddleware` sets JSON response headers.
     - `LoggingMiddleware` logs method and path.
     - `TimeoutMiddleware` enforces a 15s request timeout.
   - JWT-protected routes require a `Bearer` token.

3. **Service layer** (`cmd/internal/comment/`)
   - `Service` provides application logic and delegates to a `Store`.
   - This keeps transport concerns (HTTP) separate from storage.

4. **Database layer** (`cmd/internal/db/`)
   - Implements the `Store` interface with Postgres queries.
   - Uses `sqlx` for DB access and UUIDs for IDs.

## API routes

- `GET /alive` — health check
- `POST /api/v1/comment` — create comment (JWT required)
- `GET /api/v1/comment/{id}` — fetch comment by ID
- `PUT /api/v1/comment/{id}` — update comment (JWT required)
- `DELETE /api/v1/comment/{id}` — delete comment (JWT required)

## Auth

JWT is validated in `cmd/internal/transport/http/auth.go`. The token must be passed
as `Authorization: Bearer <token>`. The signing key is hardcoded to
`missionimpossible` (for demo purposes).

## Database and migrations

Migrations live in `migrations/` and create a `comments` table with `id`, `slug`,
`author`, and `body` columns. They are applied automatically on startup via
`Database.MigrateDB()`.

## Environment variables

The app reads database settings from the environment:

- `DB_HOST`
- `DB_PORT`
- `DB_USERNAME`
- `DB_PASSWORD`
- `DB_DB`
- `SSL_MODE`

Docker Compose sets these for local development.

## Running locally

With Docker Compose:

```
docker-compose up --build
```

Or with Taskfile:

```
task run
```

## Tests

```
task test
```

Integration tests use the `integration` build tag and a running Postgres instance:

```
task integration-test
```
