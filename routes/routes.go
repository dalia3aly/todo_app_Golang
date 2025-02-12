package routes

import (
	"todo-app-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	todoController := controllers.NewTodoController()
	categoryController := controllers.NewCategoryController()

	// Todo routes
	todoRoutes := router.Group("/api/todos")
	{
		todoRoutes.GET("/", todoController.GetAllTodos)
		todoRoutes.GET("/:id", todoController.GetTodoByID)
		todoRoutes.POST("/", todoController.CreateTodo)
		todoRoutes.PUT("/:id", todoController.UpdateTodo)
		todoRoutes.DELETE("/:id", todoController.DeleteTodo)
	}

	// Category routes
	categoryRoutes := router.Group("/api/categories")
	{
		categoryRoutes.GET("/", categoryController.GetAllCategories)
		categoryRoutes.GET("/:id", categoryController.GetCategoryByID)
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}
}
