<div align="center">

# 🚀 Go REST API

**A production-ready RESTful API built with Go, Gin, and PostgreSQL.**

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](./LICENSE)

</div>

---

## Overview

A clean, well-structured REST API demonstrating production Go patterns — JWT authentication, structured middleware chains, and type-safe error handling. Designed as a reference architecture for building scalable Go services.

## Features

- **JWT Authentication** — Secure token-based auth with refresh token rotation
- **Middleware Pipeline** — Logging, CORS, rate limiting, and auth middleware
- **Structured Error Handling** — Consistent JSON error responses with proper HTTP status codes
- **PostgreSQL Integration** — Type-safe database queries with connection pooling
- **Input Validation** — Request body validation with detailed error messages
- **Environment Configuration** — 12-factor app compliant configuration management

## Tech Stack

| Component | Technology |
|:----------|:-----------|
| Language | Go 1.22 |
| Framework | Gin |
| Database | PostgreSQL 16 |
| Auth | JWT (RS256) |
| Config | Environment Variables |

## Project Structure

```
go-rest-api/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/       # Auth, logging, CORS, rate-limit
│   ├── models/          # Data models and DTOs
│   ├── repository/      # Database access layer
│   └── services/        # Business logic layer
├── pkg/
│   └── utils/           # Shared utilities
├── go.mod
├── go.sum
└── .env.example
```

## Quick Start

```bash
# Clone
git clone https://github.com/TanishqKatiyar/go-rest-api.git
cd go-rest-api

# Configure
cp .env.example .env
# Edit .env with your PostgreSQL credentials

# Run
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description | Auth |
|:-------|:---------|:------------|:-----|
| `POST` | `/api/auth/register` | Create new account | ❌ |
| `POST` | `/api/auth/login` | Login & get token | ❌ |
| `GET` | `/api/users/me` | Get current user | ✅ |
| `GET` | `/api/resources` | List resources | ✅ |
| `POST` | `/api/resources` | Create resource | ✅ |
| `PUT` | `/api/resources/:id` | Update resource | ✅ |
| `DELETE` | `/api/resources/:id` | Delete resource | ✅ |

## License

[MIT](./LICENSE)
