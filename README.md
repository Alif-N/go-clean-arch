# Go Clean Architecture REST API

A RESTful API built with Go (Golang) following **Clean Architecture** principles for user management.

## Features

- Get all users
- Create new user
- MySQL database integration
- Environment-based configuration
- Clean Architecture pattern (Handler → Service → Repository)

## Project Structure

```
go-clean-arch/
├── internal/
│   ├── config/
│   │   └── db.go              # Database connection
│   ├── handler/
│   │   └── user_handler.go    # HTTP request handlers (Controller layer)
│   ├── models/
│   │   └── user.go            # User model definition
│   ├── repository/
│   │   ├── user_repository.go         # Repository interface
│   │   └── user_repository_impl.go    # Repository implementation
│   └── service/
│       └── user_service.go    # Business logic layer
├── main.go                    # Application entry point
├── .env                       # Environment variables
├── go.mod                     # Go module dependencies
└── README.md                  # Project documentation
```

## Architecture Layers

1. **Handler (Controller)** - Handles HTTP requests and responses
2. **Service** - Contains business logic and validation
3. **Repository** - Database operations and data access
4. **Models** - Data structures

## Prerequisites

- Go 1.26 or higher
- MySQL server running locally
- Git (optional)

## Installation

1. Clone or navigate to the project directory:
   ```bash
   cd go-clean-arch
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the project root:
   ```env
   DB_USERNAME=root
   DB_PASSWORD=your_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=your_database_name
   ```

4. Create the database and users table in MySQL:
   ```sql
   CREATE DATABASE your_database_name;
   USE your_database_name;
   
   CREATE TABLE users (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255) NOT NULL
   );
   ```

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint     | Description         | Request Body                    |
|--------|--------------|---------------------|--------------------------------|
| GET    | /users       | Get all users       | -                              |
| POST   | /users       | Create a new user   | `{"name": "John Doe"}`        |

### Example Requests

**Get all users:**
```bash
curl http://localhost:8080/users
```

**Response:**
```json
[
  {"id": 1, "name": "Alice"},
  {"id": 2, "name": "Bob"}
]
```

**Create a user:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe"}'
```

**Response:**
```json
{"id": 3, "name": "John Doe"}
```

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [MySQL](https://www.mysql.com/) - Database
- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## License

MIT License
