# Feature-Based Project with Docker

![Go Version](https://img.shields.io/badge/Go-1.23-blue?logo=go) ![Docker](https://img.shields.io/badge/Docker-Compose-blue?logo=docker) ![Feature-Based Architecture](https://img.shields.io/badge/Architecture-Feature--Based-green?logo=structure) ![License](https://img.shields.io/badge/License-MIT-green)

## Overview
This is part of the repository **Golang Clean Architectures** demonstrates a **Feature-Based architecture** implemented in Go (Golang), using Docker for containerization and MySQL as the database. The main objective of this project is to serve as a template for understanding and implementing Feature-Based architecture in modern software development.

The application includes authentication, user management, and post management features. It is built with principles of **Clean Code**, **SOLID**, and **modular design** for maintainability and scalability.

## Features
- **Feature-Based Folder Structure**: Organizes code by feature (e.g., `auth`, `post`), ensuring modularity.
- **Authentication**: JWT-based authentication with user registration and login.
- **Post Management**: Create, read, update, and delete posts.
- **Dockerized Environment**: Use of Docker Compose for app, database, and phpMyAdmin.
- **Hot Reloading**: Enabled via `cosmtrek/air`.

## Feature-Based Architecture
In **Feature-Based architecture**, the codebase is organized by **features**, with each feature having its own folder containing all relevant files. Unlike traditional layered architecture (e.g., separating by controller, service, and repository layers globally), this approach groups all files related to a specific feature together.

### Key Benefits
1. **Separation of Concerns**:
    - Each feature is self-contained and independent of others.
    - Reduces coupling and increases cohesion.
2. **Scalability**:
    - Adding new features or updating existing ones is simpler as each feature is modular.
    - Encourages clean and structured growth of the codebase.
3. **Team Collaboration**:
    - Developers can work independently on different features without stepping on each other's toes.
    - Code ownership becomes more explicit.

### Folder Structure
Here is a detailed breakdown of the folder structure used in this project:

```
Feature-Based/
├── cmd/
│   └── app/
│       └── main.go        # Main application entry point
├── configs/               # Configuration files
│   ├── config.go          # App configuration loader
├── internal/              # Application source code
│   ├── auth/              # Feature: Authentication
│   │   ├── handler/       # HTTP Handlers for Auth (e.g., Login, Register)
│   │   ├── service/       # Business Logic for Auth
│   │   ├── repository/    # Database Access Layer for Auth
│   │   ├── jwt.go         # JWT
│   │   └── user.go        # User Model
│   ├── post/              # Feature: Post Management
│   │   ├── handler/       # HTTP Handlers for Post (e.g., CRUD operations)
│   │   ├── service/       # Business Logic for Post
│   │   ├── repository/    # Database Access Layer for Post
│   │   └── post.go        # Post Model
├── pkg/                   # Shared utilities
├── go.mod                 # Go module definition
├── docker-compose.yml     # Docker Compose configuration (pre-committed)
├── Dockerfile             # Dockerfile for the application
├── .env.example           # Example environment file
└── .air.toml              # Air configuration for hot reloading
```

### Explanation of Key Folders
- **`internal/auth`**: This folder contains all files related to authentication, such as user registration, login, JWT handling, and database interactions for user data.
- **`internal/post`**: Manages post-related functionality, including creating, reading, updating, and deleting posts.
- **`pkg/`**: Contains shared utilities that can be used across multiple features (e.g., a logger).
- **`configs/`**: Manages application configuration and environment variables.

### How This Architecture Works
1. **Handlers**:
    - Responsible for receiving HTTP requests and sending HTTP responses.
    - Minimal logic; they delegate to services for business logic.
2. **Services**:
    - Contain the business logic for each feature.
    - Interact with repositories for data persistence.
3. **Repositories**:
    - Handle database interactions (e.g., CRUD operations).
    - Abstract the database layer from the rest of the application.

## Prerequisites
- Docker and Docker Compose installed.
- Golang 1.23+ installed (if running locally).

## Getting Started
### Clone the Repository
```bash
git clone https://github.com/abdorrahmani/golang-clean-architectures.git
cd golang-clean-architectures/Feature-Based
```

### Environment Variables
The `.env.example` file is pre-committed for your convenience. Copy it to `.env` and adjust the values as needed:
```bash
cp .env.example .env
```

### Build and Run with Docker
Start the application using Docker Compose:
```bash
docker-compose up --build
```
This will:
- Build and run the application.
- Set up a MySQL database.
- Start phpMyAdmin for database management (accessible at `http://localhost:8081`).

### Access the Application
- API: `http://localhost:8080`
- phpMyAdmin: `http://localhost:8081`

## Hot Reloading with Air
This project uses [cosmtrek/air](https://github.com/cosmtrek/air) for hot reloading during development. Any changes to the Go files will automatically restart the server.

## API Endpoints
### Public Endpoints
- `POST /api/v1/register`: Register a new user.
- `POST /api/v1/login`: User login.
- `GET /api/v1/posts`: Get all posts.
- `GET /api/v1/posts/:id`: Get a post by ID.
- `POST /api/v1/posts`: Create a new post.
- `PUT /api/v1/posts/:id`: Update a post.
- `DELETE /api/v1/posts/:id`: Delete a post.

### Protected Endpoints (JWT Required)
- `GET /api/v1/user`: Get authenticated user info.

## Services Overview
### Docker Compose Services
- **App**: Runs the Go application with Air for hot reloading.
- **MySQL**: Database service for storing user and post data.
- **phpMyAdmin**: Web-based UI for managing the database.

## Contributing
Feel free to fork this repository and contribute! Suggestions and pull requests are welcome.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

