# Event API

A RESTful API for event management built with Go, Gin, and GORM.

## Overview

This project is a REST API for managing events, users, and event registrations. It allows users to create accounts, create and manage events, and register for events created by other users.

## Features

- User Authentication
  - User signup and login with JWT authentication
  - Password hashing for security
- Event Management
  - Create, read, update, and delete events
  - List all events
  - Get event details
- Event Registration
  - Register for events
  - Cancel registrations
- Database
  - SQLite database for persistence
  - GORM ORM for database interactions
- API Documentation
  - OpenAPI/Swagger documentation available

## Tech Stack

- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM library
- **SQLite** - Database
- **JWT** - Authentication
- **zerolog** - Logging
- **godotenv** - Environment variable loading

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Git

### Installation

1. Clone the repository

```bash
git clone https://github.com/wfernandez/rest-api.git
cd rest-api
```

2. Install dependencies

```bash
go mod download
```

3. Create .env file in the project root:

```
PORT=8080
JWT_SECRET=your_jwt_secret_here
```

4. Run the application

```bash
go run main.go
```

The server will start on port 8080 (or the port specified in your .env file).

## API Endpoints

### Authentication

- **POST /signup** - Register a new user
- **POST /login** - Authenticate and get JWT token

### Events

- **GET /events** - Get all events
- **GET /events/:id** - Get a specific event
- **POST /events** - Create a new event (Requires authentication)
- **PUT /events/:id** - Update an event (Requires authentication, only for event creator)
- **DELETE /events/:id** - Delete an event (Requires authentication, only for event creator)

### Event Registration

- **POST /events/:id/register** - Register for an event (Requires authentication)
- **DELETE /events/:id/register** - Cancel registration for an event (Requires authentication)

## API Request/Response Examples

### User Signup

**Request:**

```json
POST /signup
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**

```json
{
  "user": {
    "ID": 1,
    "email": "user@example.com"
  }
}
```

### User Login

**Request:**

```json
POST /login
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**

```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Create Event

**Request:**

```json
POST /events
Authorization: Bearer <JWT_TOKEN>
{
  "name": "Tech Conference 2023",
  "description": "A conference about the latest technologies",
  "location": "San Francisco Convention Center",
  "dateTime": "2023-12-15T18:00:00Z"
}
```

**Response:**

```json
{
  "event": {
    "id": 1,
    "name": "Tech Conference 2023",
    "description": "A conference about the latest technologies",
    "location": "San Francisco Convention Center",
    "dateTime": "2023-12-15T18:00:00Z",
    "userId": 1
  }
}
```

### Register for Event

**Request:**

```json
POST /events/1/register
Authorization: Bearer <JWT_TOKEN>
```

**Response:**

```json
{
  "message": "Registration successful",
  "event": {
    "id": 1,
    "name": "Tech Conference 2023",
    "description": "A conference about the latest technologies",
    "location": "San Francisco Convention Center",
    "dateTime": "2023-12-15T18:00:00Z",
    "userId": 1
  }
}
```

## Project Structure

```
├── cmd/                 # Command line tools
├── data/                # Contains the SQLite database file
├── db/                  # Database connection and operations
│   ├── db-orm.go        # GORM implementation
│   └── db.go            # Database setup
├── docs/                # API documentation
│   └── openapi.yaml     # OpenAPI/Swagger spec
├── middlewares/         # HTTP middleware functions
│   └── auth.go          # Authentication middleware
├── models/              # Data models
│   ├── event.go         # Event model and operations
│   ├── modelRegistry.go # Model registration for migrations
│   ├── registration.go  # Registration model
│   └── user.go          # User model and operations
├── routes/              # API routes
│   ├── events.go        # Event endpoints
│   ├── register.go      # Registration endpoints
│   ├── routes.go        # Route setup
│   └── users.go         # User endpoints
├── swagger/             # Swagger UI setup
│   └── server.go
├── utils/               # Utility functions
│   ├── envVariables.go  # Environment variable utilities
│   ├── hash.go          # Password hashing
│   ├── jwt.go           # JWT token utilities
│   ├── logger.go        # Logging setup
│   ├── middleware.go    # Common middleware
│   └── response.go      # Response formatting
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksum
└── main.go              # Application entry point
```

## Development

### Adding a New Model

1. Create a new model file in the `models` directory
2. Register the model in the init function using `RegisterModel(&YourModel{})`
3. Implement necessary methods (Create, Read, Update, Delete)

### Environment Variables

- `PORT` - Server port (default: 8080)
- `JWT_SECRET` - Secret key for JWT signing

## License

MIT
