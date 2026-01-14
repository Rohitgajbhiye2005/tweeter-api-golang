🐦 Go Tweets API

A scalable backend REST API for a Twitter-like application built using Golang and Gin, following clean architecture principles (handler → service → repository).

This project focuses on real-world backend design, secure authentication, and efficient SQL usage.

🚀 Features
🔐 Authentication

User registration & login

JWT-based authentication

Refresh token mechanism

Protected routes using middleware

📝 Posts

Create, update, and delete posts

Get all posts (public)

Get post details by ID

Like / Unlike a post

💬 Comments

Create comments on posts

Get comments by post

Like / Unlike a comment

👍 Likes System

Toggle like / unlike

Prevent duplicate likes

Efficient like counts using SQL aggregation

🛠 Tech Stack
Backend

Golang

Gin Web Framework

MySQL

JWT (JSON Web Tokens)

Tools & Libraries

github.com/gin-gonic/gin

github.com/go-playground/validator

database/sql

docker-compose

📂 Project Structure
go-tweets/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/                 # App & DB configuration
│   ├── dto/                    # Request / response DTOs
│   ├── handlers/               # HTTP handlers
│   │   ├── user/
│   │   ├── post/
│   │   └── comment/
│   ├── middleware/             # Authentication middleware
│   ├── models/                 # Database models
│   ├── repository/             # DB access layer
│   └── service/                # Business logic
├── pkg/
│   ├── internalsql/             # DB connection
│   ├── jwt/                     # JWT utilities
│   └── refreshtoken/            # Refresh token logic
├── db/
│   ├── migrations/              # SQL migrations
│   └── schema.sql
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md

🧠 Architecture Overview
HTTP Request
   ↓
Handler (Gin)
   ↓
Service (Business Logic)
   ↓
Repository (Database)
   ↓
MySQL

Why this architecture?

Clear separation of concerns

Easy to maintain and test

Scales well for larger applications

🔑 Authentication Flow

User logs in → receives Access Token + Refresh Token

Access token is used for protected routes

Refresh token generates a new access token

Middleware validates JWT and injects userID into request context


⚙️ Environment Variables

APP_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=CHANGE_ME
DB_NAME=go_tweets
JWT_SECRET=CHANGE_ME

🐳 Run with Docker
docker-compose up --build

▶️ Run Locally
go mod tidy
go run cmd/main.go

👨‍💻 Author

Rohit Gajbhiye
Backend Developer | Golang | REST APIs
