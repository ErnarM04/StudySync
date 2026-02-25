# StudySync - Student Task Management System

A robust Go-based backend web service for managing student tasks, deadlines, and study schedules.

## 🚀 Features

- **JWT Authentication & Authorization** with role-based access (Admin/User)
- **CRUD Operations** for Users, Subjects, Tasks, Deadlines, and Sprints
- **Email Notifications** via SMTP for upcoming deadlines
- **Redis Caching** for improved performance
- **Database Migrations** using golang-migrate
- **Structured Logging** with file rotation
- **Swagger API Documentation**
- **Integration & Unit Tests**
- **Docker Support** for easy deployment
- **Graceful Shutdown** and background workers

## 🏗️ Architecture

- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Primary database
- **Redis** - Caching layer
- **JWT** - Authentication tokens
- **Docker** - Containerization
- **golang-migrate** - Database migrations

## Data model and tables

The service maps Go structs in `internal/models` to PostgreSQL tables managed by GORM and `migrations/`. Below is a concise description of each entity and how it is stored—without repeating struct definitions line by line.

| Concept | Table | Role |
|--------|--------|------|
| **User** | `users` | Represents an account: human-readable name, unique email, stored password hash, access role, and when the record was created. |
| **Subject** | `subjects` | A course or topic bucket: name plus optional longer description, with a creation timestamp. |
| **Sprint** | `sprints` | A named interval of work with fixed start and end times and a lifecycle status (for example planned, active, or completed). |
| **Task** | `tasks` | A concrete assignment: title, free-text description, workflow status, an optional due moment, foreign keys to a subject and optionally a sprint, and creation time. The database indexes tasks by status, due time, and subject for common queries. |
| **Deadline** | `deadlines` | Links a user to a task with a specific due date; used for notifications and tracking. Removing a task cascades to its deadline rows; user removal also cascades from the schema’s point of view. |

**How they connect:** each task belongs to at most one subject and may sit in a sprint; sprints group many tasks. Deadlines sit between users and tasks so the same task can carry per-user due dates where the product needs them.

## 📋 Prerequisites

- Go 1.21+
- PostgreSQL 15+
- Redis 7+
- Docker & Docker Compose (optional)

## 🛠️ Installation

### Option 1: Using Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/ernarm04/studysync.git
cd studysync

# Start services using Docker Compose
docker-compose up -d

# Updating DB migrations
migrate -path migrations -database "postgres://postgres:password@localhost:5433/studysync?sslmode=disable" up

# Execute the program
go run cmd/main.go

# The API will be available at http://localhost:8080
# Swagger UI: http://localhost:8080/swagger/index.html
```

### Option 2: Using MakerFile

```bash
# Clone the repository
git clone https://github.com/ernarm04/studysync.git
cd studysync

# Start services using Docker Compose
make docker up

# Updating DB migrations
make migrate-up

# Execute the program
make run

# For any other commands
make help
```
