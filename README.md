# Go REST API

A simple RESTful API built with Go (Golang) and Gin framework for user management.

## Features

- Get all users
- Create new user
- MySQL database integration
- Environment-based configuration

## Project Structure

```
go-rest-api/
├── config/
│   └── db.go          # Database connection configuration
├── controllers/
│   └── user-controller.go  # User request handlers
├── models/
│   └── user.go        # User model definition
├── main.go            # Application entry point
├── .env               # Environment variables (database config)
├── go.mod             # Go module dependencies
└── README.md          # Project documentation
```

## Prerequisites

- Go 1.20 or higher
- MySQL server running locally
- Git (optional)

## Installation

1. Clone or navigate to the project directory:
   ```bash
   cd go-rest-api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure the database by editing the `.env` file:
   ```env
   DB_USERNAME=root
   DB_PASSWORD=your_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=your_database_name
   ```

4. Create the database in MySQL:
   ```sql
   CREATE DATABASE your_database_name;
   ```

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description     |
|--------|----------|-----------------|
| GET    | /users   | Get all users   |
| POST   | /users   | Create a user   |

### Example Requests

**Get all users:**
```bash
curl http://localhost:8080/users
```

**Create a user:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [MySQL](https://www.mysql.com/) - Database
- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## License

MIT License
