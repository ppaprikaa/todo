package ports

import (
	"context"
	"time"

	"github.com/sanzhh/todo/internal/models"
)

type InsertDTO struct {
	Name        string
	Description string
	Date        *time.Time
}

type UpdateDTO struct {
	ID          string
	Name        *string
	Description *string
	Date        *time.Time
	Done        *bool
}

type Storage interface {
	Insert(ctx context.Context, data InsertDTO) (*models.Todo, error)
	Update(ctx context.Context, data UpdateDTO) (*models.Todo, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, offset int, limit int) ([]models.Todo, error)
	GetByDate(ctx context.Context, date time.Time) ([]models.Todo, error)
	GetByID(ctx context.Context, id string) (*models.Todo, error)
}
