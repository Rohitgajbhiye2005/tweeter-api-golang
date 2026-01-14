🐦 Go Tweets API

A scalable backend REST API for a Twitter-like application built using Golang and Gin, following clean architecture principles (handler → service → repository).

🚀 Features
🔐 Authentication

User registration & login

JWT-based authentication

Refresh token mechanism

Secure protected routes using middleware

📝 Posts

Create a post

Update a post

Delete a post

Get all posts (public)

Get post details by ID

Like / Unlike a post

💬 Comments

Create comment on a post

Get comments by post

Like / Unlike a comment

👍 Likes System

Toggle like / unlike

Prevent duplicate likes

Count likes efficiently using SQL aggregation

🛠 Tech Stack
Backend

Golang

Gin Web Framework

MySQL

JWT (JSON Web Tokens)

Database

MySQL

SQL migrations

Libraries / Tools

github.com/gin-gonic/gin

github.com/go-playground/validator

database/sql

docker-compose

JWT middleware

📂 Project Structure
go-tweets/
│
├── cmd/
│   └── main.go                 # Application entry point
│
├── internal/
│   ├── config/                 # App & DB configuration
│   ├── dto/                    # Request/response DTOs
│   ├── handlers/               # HTTP handlers (Gin)
│   │   ├── user/
│   │   ├── post/
│   │   └── comment/
│   │
│   ├── middleware/             # Auth & request middleware
│   ├── models/                 # DB models
│   ├── repository/             # DB access layer
│   │   ├── user/
│   │   ├── post/
│   │   └── comment/
│   │
│   └── service/                # Business logic layer
│       ├── user/
│       ├── post/
│       └── comment/
│
├── pkg/
│   ├── internalsql/             # DB connection setup
│   ├── jwt/                     # JWT utilities
│   └── refreshtoken/            # Refresh token logic
│
├── db/
│   ├── migrations/              # SQL migration files
│   └── schema.sql
│
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md

🧠 Architecture Overview

This project follows a clean layered architecture:

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

Easy to maintain

Easy to test

Clear separation of concerns

Scales well for large applications

🔑 Authentication Flow

User logs in → receives Access Token + Refresh Token

Access token used for protected routes

Refresh token used to generate a new access token

JWT middleware validates user and injects userID into context

📌 API Endpoints (Overview)
Auth
POST   /auth/register
POST   /auth/login
POST   /auth/refresh

Posts
POST   /tweets/                  (auth)
PUT    /tweets/:post_id/update   (auth)
DELETE /tweets/:post_id/delete   (auth)
POST   /tweets/action            (auth)   → like/unlike

GET    /tweets/                  (public)
GET    /tweets/:post_id/detail   (public)

Comments
POST   /comment/                 (auth)
POST   /comment/action           (auth)   → like/unlike

⚙️ Environment Variables

Create a .env file:

APP_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=go_tweets
JWT_SECRET=your_secret_key

🐳 Run with Docker
docker-compose up --build

▶️ Run Locally
go mod tidy
go run cmd/main.go

📈 Future Improvements

Pagination for posts & comments

Cursor-based pagination

Unit & integration tests

Rate limiting

Caching with Redis

Follow / Unfollow users

Notifications system

👨‍💻 Author

Rohit Gajbhiye
Backend Developer | Golang | REST APIs
