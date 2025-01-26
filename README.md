# Golang Clean Architectures

![Go Version](https://img.shields.io/badge/Go-1.23-blue?style=flat-square&logo=go) ![Docker](https://img.shields.io/badge/Docker-Supported-blue?style=flat-square&logo=docker) ![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

## Overview
**Golang Clean Architectures** is a repository that showcases various architecture patterns commonly used in Golang projects. Each folder in this repository demonstrates an independent architecture style, enabling developers to explore, learn, and implement them in real-world projects.

### Current Architectures
- **Feature-Based Architecture**:
    - The first implementation in this repository.
    - Organizes the codebase by features to improve modularity, scalability, and maintainability.
    - Includes a fully functional project with JWT-based authentication, user registration, and post management.


- **Vertical Slice Architecture**:
  - Focuses on organizing the codebase by vertical features (or slices) rather than horizontal technical layers.
  - Each slice contains all necessary logic (e.g., endpoints, commands, queries, and database access) for a specific functionality.
  - Demonstrates modularity, scalability, and independence of features, making it easier to scale and maintain large applications.


### Future Architectures (Coming Soon)
- **Hexagonal Architecture** (Ports and Adapters)
- **Clean Architecture**
- **Layered Architecture**
- **Microservices Architecture**

Each implementation will include:
- A fully functional project.
- Documentation explaining the architecture.
- A practical use case.

## Getting Started
### Clone the Repository
```bash
git clone https://github.com/abdorrahmani/golang-clean-architectures.git
cd golang-clean-architectures
```

### Explore Architectures
Navigate to the folder of the architecture you want to explore. For example, to explore the Feature-Based Architecture:
```bash
cd Feature-Based
```
Follow the specific `README.md` in the folder for detailed setup instructions.

## Feature-Based Architecture
The **Feature-Based Architecture** is included in the `Feature-Based` folder. It demonstrates:
- Organizing code by feature.
- Using Docker Compose for containerization.
- Implementing modular design principles.

Learn more in the [Feature-Based README](Feature-Based/README.md).


---

## Vertical Slice Architecture

The **Vertical Slice Architecture** is a feature-oriented architecture that organizes the codebase into independent slices. Each slice is responsible for a specific feature, containing everything required to implement it, including API endpoints, business logic, data models, and database access.

### Benefits:
- Feature-focused development.
- Better modularity and reduced dependencies between features.
- Easier to maintain and scale in large codebases.
- Independent development and testing of features.

Learn more in the [Vertical Slice README](vertical-slice/README.md).

---

## Contribution Guidelines
We welcome contributions! Feel free to:
1. Add new architecture patterns.
2. Improve existing implementations.
3. Provide feedback and suggestions.

To contribute:
1. Fork the repository.
2. Create a new branch.
3. Submit a pull request with a clear description of your changes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
Stay tuned for more architecture patterns and examples to help you master Go development!

