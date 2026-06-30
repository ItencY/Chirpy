# Chirpy 🐦

A RESTful API service for creating and managing "chirps" (short messages), built with Go and PostgreSQL.

## Overview

Chirpy is a backend service that provides functionality for users to create, read, and delete short messages called "chirps". The service includes user authentication with JWT tokens, token refresh/revocation capabilities, and webhook support for platform integrations.

## Features

- **User Management**: Create and authenticate users with secure password hashing (Argon2id)
- **Chirp Management**: Create, retrieve, and delete chirps with content filtering
- **Authentication**: JWT-based token authentication with token refresh and revocation
- **Webhooks**: Webhook support for external integrations (e.g., Polka platform)
- **Metrics**: Admin endpoints for monitoring fileserver hits and system reset
- **Health Check**: Simple health check endpoint for service monitoring

## Tech Stack

- **Language**: Go 1.26.3
- **Database**: PostgreSQL
- **Dependencies**:
  - `github.com/lib/pq` - PostgreSQL driver
  - `github.com/golang-jwt/jwt/v5` - JWT token handling
  - `github.com/alexedwards/argon2id` - Password hashing
  - `github.com/google/uuid` - UUID generation
  - `github.com/joho/godotenv` - Environment variable management

## API Endpoints

### Health Check
- `GET /api/healthz` - Service health check

### Users
- `POST /api/users` - Create a new user
- `PUT /api/users` - Update user information
- `POST /api/login` - Authenticate user and get JWT token
- `POST /api/refresh` - Refresh JWT token
- `POST /api/revoke` - Revoke JWT token

### Chirps
- `GET /api/chirps` - Get all chirps
- `GET /api/chirps/{chirpID}` - Get a specific chirp
- `POST /api/chirps` - Create a new chirp
- `DELETE /api/chirps/{chirpID}` - Delete a chirp

### Admin
- `GET /admin/metrics` - View metrics
- `POST /admin/reset` - Reset the system

### Webhooks
- `POST /api/polka/webhooks` - Handle Polka platform webhooks

### File Server
- `GET /app/*` - Serve static files with metrics tracking

## Getting Started

### Prerequisites

- Go 1.26.3 or higher
- PostgreSQL database
- Environment variables configured (see below)

### Environment Setup

Create a `.env` file with the following variables:

```env
DB_URL=postgres://user:password@localhost:5432/chirpy
JWT=your_jwt_secret_key
PLATFORM=production
POLKA_KEY=your_polka_webhook_key
```

### Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Project Structure

```
.
├── main.go                 # Main application entry point
├── go.mod                  # Go module definition
├── internal/
│   └── database/          # Database queries and operations
└── [other source files]
```

## Configuration

The application requires the following environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `DB_URL` | PostgreSQL connection string | Yes |
| `JWT` | JWT secret token for authentication | Yes |
| `PLATFORM` | Platform identifier | Yes |
| `POLKA_KEY` | API key for Polka webhooks | Yes |

## Development

### Building

```bash
go build -o chirpy
```

### Running Tests

```bash
go test ./...
```

## Architecture

Chirpy follows a clean architecture pattern with:
- **API Layer**: HTTP request handlers for all endpoints
- **Database Layer**: PostgreSQL queries abstracted in the `internal/database` package
- **Authentication**: JWT-based token management with refresh and revocation
- **Middleware**: Metrics tracking middleware for fileserver requests

## License

This project is open source. See LICENSE file for details.

## Author

Created by ItencY
