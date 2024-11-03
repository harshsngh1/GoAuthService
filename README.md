# GoAuthService
This is a simple authentication and authorization service implemented in Golang, providing basic functionality for token-based user authentication, token refresh, and revocation.

## Introduction
GoAuthService is an API service for managing user sessions using JSON Web Tokens (JWT). It provides signup, login, token refresh, and token revocation functionalities with role-based access control. The application is designed with modularity in mind and follows clean code principles, making it easy to expand and maintain.

## Technologies Used
- **Language:** Go
- **Token Management:** JWT (JSON Web Tokens)
- **Database:** In-memory storage using maps (future-ready for integrating a database)
- **Project Structure:** Organized into distinct directories for handlers, middleware, routes, and storage for improved scalability and maintainability

## Project Structure
```
GoAuthService/
│
├── cmd/
│   └── main.go
│
├── internals/
│   ├── handlers/
│   │   ├── get_tokens_handler.go   
│   │   ├── get_users_handler.go
│   │   ├── login_handler.go
│   │   ├── refresh_token_handler.go
│   │   ├── revoke_handler.go
│   │   ├── signup_handler.go
│   │   └── welcome_app_handler.go
│   │
│   ├── middleware/
│   │   └── auth_middleware.go
│   │
│   ├── models/
│   │   ├── custom_claims.go
│   │   ├── revoke_token.go
│   │   ├── tokens.go
│   │   ├── user_request.go
│   │   └── users.go
│   │
│   ├── routes/
│   │   └── routes.go
│   │
│   └── storage/
│       ├── token_storage.go
│       └── user_storage.go
│
├── utils/
│   └── utils.go
│
├── Dockerfile
├── go.mod
└── go.sum
```
## Installation
- Clone the repository:
```
git clone https://github.com/harshsngh1/GoAuthService.git
```
- Navigate to the project directory:
```
cd GoAuthService
```
- Run the application:
```
go run cmd/main.go
```
The server will start on localhost:8080.

## Running via Docker
To run the application using Docker:
- Build the Docker image:
```
docker build -t goauthservice .
```
- Run the Docker container:
```
docker run -p 8080:8080 goauthservice
```
The application will be accessible at http://localhost:8080.

## cURL commands for testing

### List all users
```
curl -X GET http://localhost:8080/api/v1/get-users
```
### List all tokens
```
curl -X GET http://localhost:8080/api/v1/get-tokens
```
### Register a new user
```
curl -X POST http://localhost:8080/api/v1/signup -H "Content-Type: application/json" -d '{
    "email": "testuser@example.com",
    "password": "yourpassword"
}'
```
### Login a user and generate token
```
curl -X POST http://localhost:8080/api/v1/login \
-H "Content-Type: application/json" \
-d '{
    "email": "testuser@example.com",
    "password": "yourpassword"
}'
```
**Note :**`Save the token generated from previous cUrl`

### Authentication of user via token
```
curl -X GET http://localhost:8080/api/v1/welcome \
-H "Authorization: Bearer <Your token>"
```
### Refresh Token
```
curl -X POST http://localhost:8080/api/v1/refresh-token \
-H "Authorization: Bearer <current_token>" \
-H "Content-Type: application/json"
```

### Revoke Token
```
curl -X POST http://localhost:8080/api/v1/revoke \
-H "Content-Type: application/json" \
-d '{
  "user_id": "<UserId from the get user table>",
  "token": "<Token from get tokens table>",
  "reason": "Token revocation reason"
}'
```