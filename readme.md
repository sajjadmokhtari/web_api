🚗 Web API – Car Trading System
🧠 Overview Car Trading Web API is a modular and scalable backend system built with Go. It provides a clean architecture for managing car trading operations with full CRUD support (Create, Read, Update, Delete).

It consists of several main components:

🖥️ API Server: Handles car trading endpoints and business logic, backed by Postgres.

💾 Redis: Provides caching and session management.

📊 Monitoring Stack: Grafana + Prometheus for metrics, ELK Stack for logs.

📄 Swagger UI: Interactive API documentation and testing.

The goal is to deliver a high‑performance, production‑ready API for car trading platforms.

🧱 Architecture

Code
       ┌──────────────┐
       │   Client     │  ← Frontend / Consumers
       └──────┬───────┘
              │
    ┌─────────▼─────────┐
    │   API Server      │  ← Go (CRUD, Business Logic)
    └───────┬───────────┘
            │
 ┌──────────▼───────────┐
 │      Postgres        │  ← Main Database
 └──────────▲───────────┘
            │
 ┌──────────┴───────────┐
 │       Redis          │  ← Cache / Session
 └──────────▲───────────┘
            │
 ┌──────────┴───────────┐
 │ Monitoring & Logging │  ← Grafana, Prometheus, ELK
 └──────────────────────┘
🚀 Features

✅ API Server

Full CRUD operations for cars, users, and orders

Generic structure for easy extension

Swagger UI for documentation and testing

💾 Data Layer

Postgres as the main relational database

Redis for caching and session management

📊 Monitoring & Logging

Grafana dashboards for real‑time metrics

Prometheus for metric collection

ELK Stack (Elasticsearch, Logstash, Kibana) for log analysis

⚙️ Installation

Requirements: Go 1.18+, Postgres, Redis, Docker (for monitoring stack)

Clone the repository:

bash
git clone https://github.com/sajjadmokhtari/web_api.git
cd web_api
Install dependencies:

bash
go mod tidy
Run the project:

bash
go run src/main.go
🛠️ Usage

Run the API Server

bash
go run src/main.go
The server will start and expose endpoints for car trading operations.

Access Swagger UI

Code
http://localhost:8080/swagger/index.html
Example Endpoints

POST /cars → Add a new car

GET /cars/{id} → Get car details

PUT /cars/{id} → Update car info

DELETE /cars/{id} → Remove a car

POST /orders → Create a new order
