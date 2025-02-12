package services

import (
	"errors"

	"github.com/dalia3aly/go-sqlite-backend/config"
	"github.com/dalia3aly/go-sqlite-backend/models"
)

type TodoService struct{}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := config.DB.Preload("Category").Find(&todos)
	return todos, result.Error
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	result := config.DB.Preload("Category").First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	// Verify category exists
	var category models.Category
	if err := config.DB.First(&category, todo.CategoryID).Error; err != nil {
		return errors.New("category not found")
	}

	return config.DB.Create(todo).Error
}

func (s *TodoService) UpdateTodo(id uint, todo *models.Todo) error {
	existingTodo, err := s.GetTodoByID(id)
	if err != nil {
		return err
	}

	// Update fields
	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description
	existingTodo.Completed = todo.Completed
	existingTodo.CategoryID = todo.CategoryID

	return config.DB.Save(existingTodo).Error
}

func (s *TodoService) DeleteTodo(id uint) error {
	result := config.DB.Delete(&models.Todo{}, id)
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return result.Error
}
