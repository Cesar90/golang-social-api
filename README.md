# GopherSocial

Backend API for a social network built with **Go**, **PostgreSQL**, **Docker**, **Redis**, **Swagger**, and **golang-migrate**.

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
* [golang-migrate](https://github.com/golang-migrate/migrate) — required for database migrations
* [swag](https://github.com/swaggo/swag) — required only when regenerating Swagger documentation

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

Apply the database migrations:

```bash
make migrate-up
```

Run the API:

```bash
go run ./cmd/api
```

The API listens on:

```text
http://localhost:8080
```

## Database

PostgreSQL runs inside Docker and is exposed on port `5432`.

The default development connection string used by the application is:

```env
DB_ADDR=postgres://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable
```

The Docker Compose configuration creates:

| Setting  | Value           |
| -------- | --------------- |
| Database | `socialnetwork` |
| User     | `admin`         |
| Password | `adminpassword` |
| Port     | `5432`          |

Redis is exposed on port `6379`.

It is disabled by default in the application and can be enabled with:

```env
REDIS_ADDR=localhost:6379
REDIS_ENABLED=true
```

## Database Migrations

Database migrations are managed with [golang-migrate](https://github.com/golang-migrate/migrate) and stored in:

```text
./cmd/migrate/migrations
```

### Windows

If you are running the project on Windows, you can execute migrations directly from PowerShell.

Apply all pending migrations:

```powershell
migrate -path ./cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable" up
```

Check the current migration version:

```powershell
migrate -path ./cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable" version
```

Rollback the latest migration:

```powershell
migrate -path ./cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable" down 1
```

Create a new migration:

```powershell
migrate create -seq -ext sql -dir ./cmd/migrate/migrations <migration_name>
```

For example:

```powershell
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
```

This generates two files:

```text
cmd/migrate/migrations/
├── 000001_create_users.up.sql
└── 000001_create_users.down.sql
```

The `up.sql` file contains the changes required to apply the migration:

```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL
);
```

The `down.sql` file contains the changes required to undo it:

```sql
DROP TABLE users;
```

### Using Make

If `make` is available in your development environment, the same operations can be performed through the `Makefile`.

Run all pending migrations:

```bash
make migrate-up
```

Rollback one migration:

```bash
make migrate-down n=1
```

Create a new migration:

```bash
make migration name=<migration_name>
```

For example:

```bash
make migration name=create_users
```

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

Regenerate the Swagger documentation with:

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

On Windows without `make`:

```powershell
# Start PostgreSQL and Redis
docker compose up -d

# Apply migrations
migrate -path ./cmd/migrate/migrations -database "postgres://admin:adminpassword@localhost:5432/socialnetwork?sslmode=disable" up

# Start the API
go run ./cmd/api
```

Run the test suite with:

```bash
make test
```

## Project Structure

The project is organized as follows:

```text
.
├── cmd/
│   ├── api/                  # HTTP API entrypoint and handlers
│   └── migrate/
│       └── migrations/       # SQL database migrations
├── docs/                     # generated Swagger documentation
├── internal/                 # application packages
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
| PostgreSQL 16.3  | Relational database        |
| Redis 6.2        | Optional caching           |
| Swagger / swaggo | API documentation          |
| golang-migrate   | Database schema migrations |

## Useful Commands

```bash
# Start services
docker compose up -d

# Stop services
docker compose down

# Run API
go run ./cmd/api

# Apply migrations
make migrate-up

# Rollback migration
make migrate-down n=1

# Create migration
make migration name=create_users

# Seed database
make seed

# Generate Swagger documentation
make gen-docs

# Run tests
make test
```

For Windows users without `make`, the equivalent migration commands can be executed directly with the `migrate` CLI as shown in the **Database Migrations** section.
