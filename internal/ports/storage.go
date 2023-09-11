package ports

import (
	"context"
	"errors"
	"time"

	"github.com/sanzhh/todo/internal/models"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type InsertDTO struct {
	Name        string
	Description string
	Date        *string
}

type UpdateDTO struct {
	ID          string
	Name        *string
	Description *string
	Date        *string
	Done        *bool
}

type Storage interface {
	Insert(ctx context.Context, data InsertDTO) error
	Update(ctx context.Context, data UpdateDTO) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, offset uint64, limit uint64) ([]models.Todo, error)
	GetByDate(ctx context.Context, date time.Time) ([]models.Todo, error)
	GetByID(ctx context.Context, id string) (*models.Todo, error)
}
