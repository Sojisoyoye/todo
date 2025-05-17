package service

import (
	"github.com/Sojisoyoye/todo/internal/models"
	"github.com/Sojisoyoye/todo/internal/repository"
	"github.com/google/uuid"
)

type TodoServiceInterface interface {
	CreateTodo(todo *models.Todo) error
	GetAllTodos() ([]models.Todo, error)
	GetTodoByID(id uuid.UUID) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uuid.UUID) error
}

type TodoService struct {
	repo repository.TodoRepositoryInterface
}

func NewTodoService(repo repository.TodoRepositoryInterface) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.FindAll()
}

func (s *TodoService) GetTodoByID(id uuid.UUID) (*models.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *TodoService) UpdateTodo(todo *models.Todo) error {
    existing, err := s.repo.FindByID(todo.ID)
    if err != nil {
        return err
    }
  
    existing.Title = todo.Title
    existing.Description = todo.Description
    existing.Completed = todo.Completed
    return s.repo.Update(existing)
}

func (s *TodoService) DeleteTodo(id uuid.UUID) error {
	return s.repo.Delete(id)
}
