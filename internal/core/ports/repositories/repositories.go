package repositories

import (
	"awesomeProject2/internal/core/entities"
	"context"
)

type Repository interface {
	Add(ctx context.Context, profile entities.Profile) error
}
type ReadRepository interface {
	GetByID(ctx context.Context, id string) (*entities.Profile, error)
}
