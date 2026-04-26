# TeamUp
TeamUp is a platform for finding partners and building teams for startups and pet projects.

## Status

Implemented:
- User registration and authentication
- Posts feed (search + tag filtering)
- Post details page
- Create post
- Respond to a post (message + Telegram contact)
- Post responses list for post owners
- "My posts" page for authors
- Profile page (from `/api/v1/profile/me`)

Planned:
- Projects and participants
- Chat
- Notifications

## Tech Stack

Frontend:
- React 19 + TypeScript 5.9 (Vite)
- Tailwind CSS 4.1
- Zustand 5
- Axios

Backend:
- Go 1.25 + Gin
- PostgreSQL 18 (pgx)
- sqlc

Infra:
- Docker + Docker Compose

## Getting Started (Local)

### 1) Clone
```bash
git clone https://github.com/IsMirakl/TeamUp.git
cd TeamUp
```

### 2) Configure environment
Backend loads env from `backend-go/cmd/.env`.

Required variables:
- `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `DB_HOST`, `DB_PORT`
- `SECRET_KEY`, `REFRESH_TOKEN_KEY`

### 3) Start Postgres
```bash
cd backend-go
docker compose up -d
```

If this is your first run, apply SQL migrations from `backend-go/internal/database/migration`.

If you use golang-migrate:
```bash
migrate -path internal/database/migration \
  -database "postgres://<user>:<pass>@localhost:5432/teamup?sslmode=disable" \
  up
```

### 4) Start Backend API
```bash
cd backend-go
go mod tidy
go run cmd/server.go
```

Backend runs on `http://localhost:8080`.

### 5) Start Frontend
```bash
cd frontend
npm install
npm run dev
```

Frontend runs on `http://localhost:5173` and calls API at `http://localhost:8080` (see `frontend/src/api/axiosConfig.ts`).

## API
Base path: `http://localhost:8080/api`

Common endpoints:
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `GET /api/v1/profile/me`

Posts:
- `GET /api/v1/posts/post?limit=50&offset=0`
- `POST /api/v1/posts/post`
- `GET /api/v1/posts/post/:id`
- `PUT /api/v1/posts/post/:id`
- `POST /api/v1/posts/post/:id/responses` (body: `message`, `telegram`)
- `GET /api/v1/posts/post/:id/responses` (owner-only UI)

## Repo Layout
- `backend-go/` Go API server
- `frontend/` React app

## Notes for Contributors
- sqlc: enum types created inside `DO $$ ... $$` blocks are not parsed by sqlc, so we keep minimal schema-only definitions in `backend-go/internal/database/schema.sql`.

## Author
Daniil Bodrov
