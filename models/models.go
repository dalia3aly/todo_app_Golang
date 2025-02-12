package models

import (
	"time"

	"gorm.io/gorm"
)

// Base model for common fields (matches BaseEntity.java)
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

// Category model (matches Category.java)
type Category struct {
	Model
	Name  string `gorm:"unique;not null" json:"name"`
	Todos []Todo `gorm:"foreignKey:CategoryID" json:"todos,omitempty"` // One-to-many relationship with Todo
}

// Todo model (matches Todo.java)
type Todo struct {
	Model
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CategoryID  uint      `gorm:"not null" json:"categoryId"` // Foreign key to Category
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// BeforeCreate hook to set default values
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

// BeforeUpdate hook to update the UpdatedAt timestamp
func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

// TableName for Category
func (Category) TableName() string {
	return "categories"
}

// TableName for Todo
func (Todo) TableName() string {
	return "todos"
}
