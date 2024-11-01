# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [UNRELEASED] - 2024-11-01
### Added
- Initial project setup with Go and Gin for backend framework.
- Config management with Viper for environment-specific settings (`config/config.go`).
- Database connection with GORM, including PostgreSQL configuration (`config/database.go`).
- User model with password hashing and constraints (`models/user.go`).
- Auth service with JWT generation and password verification (`services/auth.go`).
- Middleware for JWT authentication (`middleware/auth.go`).
- Basic error handling utility (`utils/errors.go`).
- Initial routing structure, including:
    - `auth` routes for `/register` and `/login`.
    - Protected `/user/profile` route using JWT middleware.
- Controllers for user registration and login (`controllers/user.go`).
- README with project overview, structure, and setup instructions.
