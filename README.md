# 🛒 Ecommerce Microservice with Go & Fiber

A high-performance, scalable **Ecommerce Backend** built with **Go (Golang)** and the **Fiber** framework, following microservice architecture.

## 🚀 Features
- **User Authentication** (JWT-based)
- **Product Management** (CRUD operations)
- **Order Processing** (Checkout, Payments)
- **Inventory Management**
- **RESTful API** with Fiber
- **Dockerized** for easy deployment
- **PostgreSQL** database
- **Swagger** API documentation

## 📦 Project Structure

ecommerce-microservice/
├── cmd/ # Main application entrypoints
├── internal/ # Core business logic
│ ├── models/ # Database models (GORM)
│ ├── handlers/ # HTTP request handlers
│ ├── middleware/ # Auth & logging middleware
│ └── repository/ # Database operations
├── pkg/ # Shared utilities (config, errors, etc.)
├── api/ # API routes (Fiber)
├── docs/ # Swagger/OpenAPI docs
├── Dockerfile # Docker configuration
├── docker-compose.yml # Multi-container setup
├── go.mod # Go dependencies
└── README.md # Project documentation


## 🛠️ Prerequisites
- **Go 1.21+** ([Install](https://go.dev/dl/))
- **PostgreSQL** ([Install](https://www.postgresql.org/download/))
- **Docker** (Optional, for containerization) ([Install](https://docs.docker.com/get-docker/))

## ⚡ Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/nadduli/ShipSwift.git
cd ShiptSwift