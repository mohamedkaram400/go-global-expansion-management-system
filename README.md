# 🌍 Expansion Project

A **scalable backend system** built with **Golang (Gin)** to help companies **expand into new countries** through an easy, automated, and structured process.  
This project follows a **modular clean architecture** and supports **multiple services** like MySQL, MongoDB, and Redis — all orchestrated with Docker.

---

## 📑 Table of Contents
- [Tech Stack](#-tech-stack)
- [Requirements](#-requirements)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [Usage](#-usage)
- [Example API Request](#-example-api-request)
- [Features](#-features)
- [Endpoint Collection](#-endpoint-collection)
- [License](#-license)


## 🧰 Tech Stack

- 🐹 **Golang (Gin)** — REST API framework  
- 🐬 **MySQL** — Primary relational database  
- 🍃 **MongoDB** — Document storage for research and analytics  
- 🧠 **Redis** — Caching and session management  
- 🐳 **Docker & Docker Compose** — Containerization and environment orchestration

---

## 📋 Requirements

Before you begin, ensure you have the following installed:

- [Golang](https://go.dev/) `>= 1.21`
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- Git

---

### 🧠 Project Structure
```
.
├── cmd/                  # Application entry point
├── config/               # Configuration files
├── conn/                 # Database connections (MySQL, Mongo, Redis)
├── deploy/               # Docker & Nginx configs
├── internal/
│   ├── adapters/         # Repositories, notifiers, schedulers
│   ├── core/             # Business logic & services
│   ├── delivery/         # HTTP Handlers & Routes
├── db/
│   ├── migrations/       # Migrations
│   └── seeders/          # Initial data seeding
├── pkg/                  # Shared utilities & helpers
├── docker-compose.yml
├── Dockerfile
└── README.md
```
---

## 🚀 Getting Started

### 1. Clone the Repository
```bash
git clone git@github.com:mohamedkaram400/go-global-expansion-management-system.git
cd go-global-expansion-management-system
```

## ⚙️ Setup Environment
```bash
cp .env.example .env
```

## 🐳 Docker Setup
```bash
docker -f deploy/docker-compose.yml compose up --build
```
---

### Run the Application (Without Docker)

#### If you want to run locally without Docker:

```bash
go mod tidy
go run cmd/main.go
```
---

## 🧪 Usage
Once the application is running:

The server will be available at: http://localhost:9000

You can use tools like Postman or cURL to test the API.

---
## 🧪 Example API Request

### 📝 Create a New Project
```http
POST /api/v1/project/create-project
Content-Type: application/json
```
#### 📥 Request

```json
{
  "client_id": 1,
  "service_needed": ["Legal", "Recruitment"],
  "country": "Germany",
  "budget": 10000
}
```

#### 📤 Response

```json
{
    "message": "Project Created Successfully",
    "data": {
        "id": 22,
        "client_id": 6,
        "service_needed": [
            "Legal",
            "Recruitment"
        ],
        "country": "Germany",
        "budget": 10000,
        "status": "active"
    }
}
```

---
## ✨ Features:
- Modular clean architecture
- Multi-database support (MySQL + MongoDB + Redis)
- Project creation and management system
- Easy environment setup with Docker
- Centralized configuration management
- Ready for scaling and service extension

---
## 📚 Endpoint Collection
The full list of API endpoints is included in the project’s main directory for easy reference and testing.

---

## 🛡️ License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, and share this project with proper attribution.
