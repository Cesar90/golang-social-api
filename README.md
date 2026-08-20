# GopherSocial

Backend API for a social network built with Go, PostgreSQL, Docker, Swagger, and golang-migrate.

## Tech Stack

* **Go 1.27** — application/backend
* **Docker** — containerized local services
* **PostgreSQL 16.3** — primary database running in Docker
* **Redis 6.2** — optional cache running in Docker
* **Swagger / swaggo** — API documentation
* **golang-migrate** — database migrations

## Requirements

Make sure you have the following installed:

* [Go 1.27+](https://go.dev/)
* [Docker](https://www.docker.com/)
* [golang-migrate](https://github.com/golang-migrate/migrate) — required for migrations
* [swag](https://github.com/swaggo/swag) — required only when regenerating Swagger docs

## Getting Started

Clone the repository:

```bash
git clone https://github.com/Cesar90/golang-social-api.git
cd golang-social-api
```

Start PostgreSQL and Redis:

```bash
docker compose up -d
```

Run the API:

```bash
go run ./cmd/api
```

The API listens on `http://localhost:8080` by default.

## Database

PostgreSQL runs inside Docker on port `5432`.

The default development connection string used by the application is:

```env
DB_ADDR=postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable
```

The Docker Compose configuration creates:

* database: `socialnetwork`
* user: `admin`
* password: `adminpassword`

Redis is exposed on port `6379`. It is disabled by default in the application and can be enabled with:

```env
REDIS_ADDR=localhost:6379
REDIS_ENABLED=true
```

## Database Migrations

Database migrations are managed with **golang-migrate** and stored in:

```text
./cmd/migrate/migrations
```

Run all pending migrations:

```bash
make migrate-up
```

Rollback migrations:

```bash
make migrate-down 1
```

Create a new migration:

```bash
make migration <migration_name>
```

This creates migration files under `cmd/migrate/migrations/`.

Seed the database:

```bash
make seed
```

## API Documentation

Swagger documentation is generated with **swaggo** and served by the API.

Once the application is running, Swagger UI is available at:

```text
http://localhost:8080/v1/swagger/index.html
```

Regenerate the Swagger docs with:

```bash
make gen-docs
```

## Development

A typical local-development workflow is:

```bash
# Start PostgreSQL and Redis
docker compose up -d

# Apply migrations
make migrate-up

# Start the API
go run ./cmd/api
```

Run the test suite with:

```bash
make test
```

## Project Structure

The reference project is organized like this:

```text
.
├── cmd/
│   ├── api/              # HTTP API entrypoint and handlers
│   └── migrate/          # migrations and seed command
├── docs/                 # generated Swagger documentation
├── internal/             # application packages
├── scripts/
├── web/
├── .air.toml
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## Tech Overview

| Technology       | Purpose                    |
| ---------------- | -------------------------- |
| Go 1.27          | Backend development        |
| Docker           | Containerization           |
| PostgreSQL       | Relational database        |
| Redis            | Optional caching           |
| Swagger / swaggo | API documentation          |
| golang-migrate   | Database schema migrations |
