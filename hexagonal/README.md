# Hexagonal Architecture Example (Ports & Adapters)

![Go Version](https://img.shields.io/badge/Go-1.23-blue?logo=go) ![Docker](https://img.shields.io/badge/Docker-Compose-blue?logo=docker) ![Hexagonal Architecture](https://img.shields.io/badge/Architecture-Hexagonal-green?logo=structure) ![License](https://img.shields.io/badge/License-MIT-green)

## Overview
This project demonstrates a **Hexagonal Architecture (Ports & Adapters)** in Go, as part of the [Golang Clean Architectures](../README.md) repository. It provides a clean separation between core business logic, application services, and external interfaces (HTTP, database), making the codebase highly maintainable and testable.

The application is a simple Post management API with full CRUD functionality, using Gin for HTTP, GORM for MySQL, and Docker for containerization.

## Features
- **Hexagonal (Ports & Adapters) Architecture**: Clear separation of domain, application, and infrastructure.
- **Post CRUD API**: Create, read, update, and delete posts.
- **Dockerized Environment**: Use Docker Compose for app, database, and phpMyAdmin.
- **Hot Reloading**: Enabled via `air`.

## Hexagonal Architecture
In Hexagonal Architecture, the application is organized around the core domain and use cases (the "inside"), with adapters (the "outside") for interacting with the world (e.g., HTTP, database). This approach:
- Promotes testability and maintainability.
- Allows easy swapping of external technologies.
- Keeps business logic independent from frameworks and drivers.

### Folder Structure
```
hexagonal/
├── cmd/
│   └── app/
│       └── main.go            # Application entry point
├── internal/
│   ├── adapters/
│   │   └── http/
│   │       ├── post_handler.go        # HTTP handlers for Post
│   │       ├── post_repository_gorm.go # GORM implementation of PostRepository
│   │       └── routes.go              # Route registration
│   ├── app/
│   │   ├── post_repository.go         # PostRepository interface (port)
│   │   └── post_service.go            # Application service (use case)
│   ├── config/
│   │   ├── config.go                  # Config loader
│   │   └── db.go                      # DB connection & migration
│   └── domain/
│       └── post.go                    # Post entity
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
```

### Key Concepts
- **Domain**: Core business entities (e.g., `Post`).
- **Application (App)**: Use cases and service logic, defined in terms of ports (interfaces).
- **Adapters**: Implementations of ports for external systems (e.g., HTTP handlers, GORM repository).

## Prerequisites
- Docker and Docker Compose
- Go 1.23+ (for local development)

## Getting Started
### Clone the Repository
```bash
git clone https://github.com/abdorrahmani/golang-clean-architectures.git
cd golang-clean-architectures/hexagonal
```

### Environment Variables
Copy the example environment file and adjust as needed:
```bash
cp .env.example .env
```

### Build and Run with Docker
```bash
docker-compose up --build
```
This will:
- Build and run the Go application
- Set up a MySQL database
- Start phpMyAdmin (http://localhost:8081)

### Access the Application
- API: `http://localhost:8080`
- phpMyAdmin: `http://localhost:8081`

## API Endpoints
All endpoints are under `/posts`:
- `POST   /posts`      - Create a new post
- `GET    /posts`      - Get all posts
- `GET    /posts/:id`  - Get a post by ID
- `PUT    /posts/:id`  - Update a post
- `DELETE /posts/:id`  - Delete a post

### Post Entity
```go
{
  "id": int,
  "title": string,
  "content": string,
  "createdAt": string, // RFC3339 timestamp
  "updatedAt": string  // RFC3339 timestamp
}
```

## Hot Reloading with Air
This project uses [cosmtrek/air](https://github.com/cosmtrek/air) for hot reloading during development. Any changes to Go files will automatically restart the server.

## Docker Compose Services
- **app**: Go application with Air for hot reload
- **db**: MySQL database
- **phpmyadmin**: Web UI for DB management
- **test_db**: MySQL for testing (not used by default)

## Contributing
Contributions are welcome! Fork, branch, and submit a pull request.

## License
MIT License. See [LICENSE](../LICENSE). 