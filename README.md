# FoodCourt Project

## Prerequisites

- Docker and Docker Compose installed on your system.
- Basic understanding of Docker and Go (Golang).

## Configuration

### Docker

Use `docker-compose.yml` to define services, including MariaDB for the database and the backend application.

#### Database Service

- Utilizes MariaDB.
- Mounts the initialization script from `./database/init.sql`.
- Configures the database with environment variables for security and access.

#### Backend Service

- Compiles the Go application as per the `Dockerfile`.
- Passes SMTP configuration for email functionalities through environment variables. Replace `${SMTP_HOST}`, `${SMTP_USER}`, and `${SMTP_PASS}` with your own SMTP server details.

#### Adminer

- Provides a web interface for database management, accessible at `localhost:1333`.

### Dockerfile

Outlines the construction of the Go application:

- Uses the official Go image.
- Copies `go.mod` and `go.sum`, then downloads dependencies.
- Compiles the application and designates the compiled app as the entry point.

### Environment Variables

Create a `.env` file in the project root with your SMTP server details:

```env
SMTP_HOST=smtp.example.com
SMTP_USER=user@example.com
SMTP_PASS=secret
```

## Running the Project

In the project directory, execute:

```bash
docker-compose up --build
```

This builds and starts the containers.

## Accessing the Application

With the containers up, access the backend via `http://localhost:8095`.

## Stopping the Project

To halt the containers, use:

```bash
docker-compose down
```
