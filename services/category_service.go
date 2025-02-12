package services

import (
	"errors"

	"github.com/dalia3aly/go-sqlite-backend/config"
	"github.com/dalia3aly/go-sqlite-backend/models"
)

type CategoryService struct{}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	return categories, result.Error
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	result := config.DB.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return config.DB.Create(category).Error
}

func (s *CategoryService) UpdateCategory(id uint, category *models.Category) error {
	existingCategory, err := s.GetCategoryByID(id)
	if err != nil {
		return err
	}

	existingCategory.Name = category.Name
	return config.DB.Save(existingCategory).Error
}

func (s *CategoryService) DeleteCategory(id uint) error {
	// Check if category has todos
	var count int64
	if err := config.DB.Model(&models.Todo{}).Where("category_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("cannot delete category with existing todos")
	}

	result := config.DB.Delete(&models.Category{}, id)
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return result.Error
}
