package repository

import (
	"github.com/Sojisoyoye/todo/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepositoryInterface interface {
	Create(todo *models.Todo) error
	FindAll() ([]models.Todo, error)
	FindByID(id uuid.UUID) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id uuid.UUID) error
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) FindByID(id uuid.UUID) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Todo{}, id).Error
}