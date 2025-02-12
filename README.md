# Todo App Backend API

A modern, production-ready RESTful API built with Go, Gin, and GORM. This API serves as the backend for a Todo application, providing robust endpoints for managing todos and their categories.

## 🚀 Tech Stack

- **Go** - Core programming language
- **Gin** - High-performance web framework
- **GORM** - Feature-rich ORM for database operations
- **PostgreSQL** - Database (via Supabase)
- **Supabase** - Database hosting and management

## ✨ Features

- **RESTful API Design** - Following REST best practices
- **Clean Architecture** - Organized in models, services, and controllers
- **Database Integration** - PostgreSQL with GORM for reliable data persistence
- **Auto Migrations** - Automatic database schema updates
- **CORS Support** - Configured for frontend integration
- **Environment Based Config** - Different settings for development and production
- **Error Handling** - Comprehensive error handling and reporting
- **Input Validation** - Request validation using GORM and Gin
- **Relationship Handling** - One-to-many relationship between categories and todos

## 📚 API Documentation

### Todo Endpoints

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/api/todos` | Get all todos | - |
| GET | `/api/todos/:id` | Get todo by ID | - |
| POST | `/api/todos` | Create new todo | `{"title": "string", "description": "string", "categoryId": number}` |
| PUT | `/api/todos/:id` | Update todo | `{"title": "string", "description": "string", "completed": boolean, "categoryId": number}` |
| DELETE | `/api/todos/:id` | Delete todo | - |

### Category Endpoints

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/api/categories` | Get all categories | - |
| GET | `/api/categories/:id` | Get category by ID | - |
| POST | `/api/categories` | Create new category | `{"name": "string"}` |
| PUT | `/api/categories/:id` | Update category | `{"name": "string"}` |
| DELETE | `/api/categories/:id` | Delete category | - |

## 🛠️ Development Setup

1. **Prerequisites**
   - Go 1.21 or higher
   - PostgreSQL (or Supabase account)
   - Git

2. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd todo-app-backend
   ```

3. **Environment Setup**
   ```bash
   cp .env.example .env
   # Update .env with your database credentials
   ```

4. **Install Dependencies**
   ```bash
   go mod download
   ```

5. **Run the Server**
   ```bash
   go run main.go
   ```

   Server will start at `http://localhost:8080`

## 📁 Project Structure

```
todo-app-backend/
├── config/         # Configuration files
├── controllers/    # HTTP request handlers
├── models/         # Database models
├── routes/         # API route definitions
├── services/       # Business logic
├── .env           # Environment variables
├── .env.example   # Example environment file
├── main.go        # Application entry point
└── README.md      # Project documentation
```

## 🔄 Database Models

### Todo Model
```go
type Todo struct {
    ID          uint      
    Title       string    
    Description string    
    Completed   bool      
    CategoryID  uint      
    Category    Category  
    CreatedAt   time.Time 
    UpdatedAt   time.Time 
    DeletedAt   gorm.DeletedAt
}
```

### Category Model
```go
type Category struct {
    ID        uint      
    Name      string    
    Todos     []Todo    
    CreatedAt time.Time 
    UpdatedAt time.Time 
    DeletedAt gorm.DeletedAt
}
```

## 🚀 Deployment

The API is designed to be deployed to various cloud platforms:

1. **Railway** (used by me)
   - Easy deployment process
   - Good free tier
   - Automatic HTTPS
2. **Render**
3. **Fly.io**

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| DATABASE_URL | PostgreSQL connection string | Yes | - |
| PORT | Server port | No | 8080 |
| GO_ENV | Environment (development/production) | No | development |

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Support

For support, email your-email@example.com or create an issue in the repository.
