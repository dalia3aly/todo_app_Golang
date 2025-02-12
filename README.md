# Todo App Backend

A Go-based REST API for managing todos and categories.

## Features

- RESTful API endpoints for todos and categories
- PostgreSQL database with GORM
- CORS enabled
- Environment-based configuration
- Automatic database migrations

## API Endpoints

### Todos
- `GET /api/todos` - Get all todos
- `GET /api/todos/:id` - Get a specific todo
- `POST /api/todos` - Create a new todo
- `PUT /api/todos/:id` - Update a todo
- `DELETE /api/todos/:id` - Delete a todo

### Categories
- `GET /api/categories` - Get all categories
- `GET /api/categories/:id` - Get a specific category
- `POST /api/categories` - Create a new category
- `PUT /api/categories/:id` - Update a category
- `DELETE /api/categories/:id` - Delete a category

## Environment Variables

- `DATABASE_URL` - PostgreSQL connection string
- `PORT` - Server port (default: 8080)
- `GO_ENV` - Environment (development/production)

## Development

1. Clone the repository
2. Copy `.env.example` to `.env` and update the values
3. Run `go mod download`
4. Run `go run main.go`

## Deployment

The API can be deployed to various platforms:

1. Railway.app (recommended)
2. Render
3. Fly.io

Make sure to set the environment variables in your deployment platform.
