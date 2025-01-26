# Vertical Slice Architecture Boilerplate

![Go Version](https://img.shields.io/badge/Go-1.23-blue?logo=go) ![Docker](https://img.shields.io/badge/Docker-Compose-blue?logo=docker) ![Vertical Slice Architecture](https://img.shields.io/badge/Architecture-Vertical--Slice-red?logo=structure) ![License](https://img.shields.io/badge/License-MIT-green)


## Overview

This project demonstrates a boilerplate implementation of the **Vertical Slice Architecture** using **Golang**. It is designed to showcase a clean, modular, and feature-oriented architecture that separates application layers by functionality rather than technical concerns. This repository is containerized with **Docker** and includes a database layer powered by MySQL.

---

## What is Vertical Slice Architecture?

The Vertical Slice Architecture focuses on splitting an application into distinct slices, each of which contains all the technical layers (e.g., controllers, services, models, and repositories) needed to implement a single feature or functionality. Each slice is self-contained and encapsulated, ensuring minimal coupling with other parts of the system.

### Key Characteristics:
1. **Feature-Oriented:** Each slice represents a single feature or use case, such as `CreateProduct`, `GetProductByID`, etc.
2. **Encapsulation:** All logic, including controllers, business rules, and data access, resides in the same slice.
3. **Independence:** Slices are designed to be self-contained, minimizing dependencies on other slices or shared resources.

### Folder Structure
Here is a high-level structure of the project:

```
internal
└───modules
    ├───products
    │   ├───commands
    │   ├───dtos
    │   ├───endpoints
    │   ├───models
    │   ├───queries
    │   └───repository
    └───shared
```

Each module (e.g., `products`) has:
- **Commands:** Handle write operations like create, update, or delete.
- **DTOs (Data Transfer Objects):** Represents the structure of data being transferred between the API and internal layers.
  - **Request DTOs:** Define the expected input structure from the user.
  - **Response DTOs:** Define the output structure returned to the user.
- **Queries:** Handle read operations like retrieving data.
- **Endpoints:** API endpoints for user interaction.
- **Models:** Represent the domain entities.
- **Repository:** Manage database access.

---

## Advantages of Vertical Slice Architecture

1. **Feature Focused:** Encourages developers to think in terms of features rather than technical layers.
2. **Independent Development:** Each slice can be developed and tested independently.
3. **Better Modularity:** Clear separation of concerns between features.
4. **Scalability:** Easy to add new features without affecting existing ones.
5. **Reduced Merge Conflicts:** Teams working on different slices rarely need to touch the same files.

---

## Disadvantages

1. **Learning Curve:** Developers unfamiliar with this architecture may require time to adapt.
2. **Duplication:** Common logic may need to be duplicated across slices to maintain independence.
3. **Boundary Management:** Mismanagement of boundaries between slices can lead to coupling.

---

## Managing the System
### Adding a New Feature
1. Create a new folder under `modules` with subfolders for `commands`, `queries`, etc.
2. Implement the feature's logic, starting with repository methods, followed by commands, handlers, and endpoints.

### Dependency Injection
Use the `shared` module to centralize and inject dependencies like database connections, logging, and configuration.

### Testing
Each slice is designed to be testable in isolation. Write unit and integration tests for the following:
- Commands (e.g., creating or updating entities).
- Queries (e.g., retrieving data).
- Endpoints (e.g., API integration tests).

---

## Running the Project

This project is containerized using Docker for consistent and portable development environments.

### Prerequisites
- Docker
- Docker Compose

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/abdorrahmani/golang-clean-architectures.git
   cd golang-clean-architectures/vertical-slice
   ```

2. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```

3. Start the containers:
   ```bash
   docker-compose up --build
   ```

4. Access the services:
    - **App:** [http://localhost:8080](http://localhost:8080)
    - **phpMyAdmin:** [http://localhost:8081](http://localhost:8081)

---

## Docker Configuration

### Dockerfile
The application runs in a lightweight Alpine-based Go container with live reload enabled by `air`.

### docker-compose.yml
The `docker-compose.yml` file orchestrates multiple services:

- **App:** Runs the Go application.
- **MySQL:** Provides the database backend.
- **phpMyAdmin:** Web interface for managing the database.

---

## Environment Variables

The project uses the following environment variables, defined in the `.env` file.

## Contribution Guidelines

1. Follow the folder structure strictly to maintain consistency.
2. Write unit and integration tests for every feature.
3. Use DTOs to manage data exchange between layers.
4. Avoid unnecessary dependencies between slices to keep the architecture clean and modular.
5. Document your code to help others understand it.

---

Feel free to fork this repository and contribute! Suggestions and pull requests are welcome.
