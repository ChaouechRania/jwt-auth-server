# JWT with Gin Framework

This project demonstrates how to use the Gin framework for building a RESTful API with JWT authentication. It includes libraries for ORM, environment management, and password encryption.

## Development Setup

To set up the project, install the following dependencies:

### Dependencies
- **Gin Framework**: A lightweight HTTP web framework.
  ```bash
  go get -u github.com/gin-gonic/gin
  ```

- **GORM**: An ORM library for database interactions.
  ```bash
  go get -u github.com/jinzhu/gorm
  ```

- **JWT**: To authenticate and generate JSON Web Tokens.
  ```bash
  go get -u github.com/dgrijalva/jwt-go
  ```

- **Godotenv**: To manage environment variables.
  ```bash
  go get -u github.com/joho/godotenv
  ```

- **Crypto**: To encrypt user passwords.
  ```bash
  go get -u golang.org/x/crypto
  ```

---

## Endpoints

### Health Check

- **Endpoint**: `GET /health`
- **Description**: Simple endpoint to check if the service is running.

#### Example Response:
```json
{
   "message": "Health check!"
}
```