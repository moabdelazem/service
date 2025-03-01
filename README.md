# service

## Description
A Go web application with Echo framework and modern project structure.

## Project Structure
- cmd/: Main applications
- internal/
  - handlers/: HTTP request handlers
  - middleware/: Custom middleware
  - models/: Data models
  - routes/: Route definitions
  - services/: Business logic
- pkg/: Library code
- docs/: Documentation
- test/: Tests

## Getting Started
1. Install dependencies:
   ~~~
   go mod tidy
   ~~~

2. Configure environment:
   ~~~
   cp .env.example .env
   ~~~

3. Run the server:
   ~~~
   go run ./cmd/main.go
   ~~~

## API Endpoints
- GET /: Welcome message
