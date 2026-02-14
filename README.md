# Handbook Backend

Backend server for the Handbook education platform, built with Go and MongoDB.

## Features

- **User Authentication**: Secure registration and login using JWT and bcrypt.
- **Role-Based Access Control**: Support for `user` and `admin` roles.
- **Course Management**: Create, read, and delete educational courses.
- **Progress Tracking**: Track user progress through course sections with "step-by-step" access.
- **Rating System**: Users can rate courses (1-5 stars), with automatic average calculation.
- **Certificate Generation**: Automated PDF certificate generation upon course completion.
- **API Documentation**: Interactive Swagger UI for API exploration.

## Tech Stack

- **Language**: Go (Golang)
- **Database**: MongoDB
- **Auth**: JWT (JSON Web Tokens)
- **PDF Generation**: `gofpdf`
- **Documentation**: `swaggo` (Swagger 2.0)

## Prerequisites

- Go 1.22 or higher
- MongoDB instance (Local or Atlas)

## Getting Started

1. **Clone the repository**
2. **Set up environment variables**:
   Create a `.env` file in the root directory:
   ```env
   MONGO_URI=your_mongodb_uri
   JWT_SECRET=your_jwt_secret
   PORT=8080
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```
4. **Run the server**:
   ```bash
   go run main.go
   ```

## API Documentation

Once the server is running, you can access the interactive Swagger documentation at:
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Maintenance and Migrations

### Migrate User Roles
If you have old user data without roles, run the migration script to assign the default "user" role:
```bash
go run scripts/migrate_roles.go
```

## Project Structure

- `handlers/`: API endpoint controllers.
- `models/`: Database schemas and Go structs.
- `config/`: Database connection and configuration logic.
- `docs/`: Auto-generated Swagger documentation.
- `frontend/`: Static frontend files (HTML/JS).
- `scripts/`: Maintenance and migration scripts.

## License

This project is licensed under the Apache 2.0 License.
