package get

import (
	"awesomeProject2/internal/core/entities"
	"awesomeProject2/internal/core/ports/repositories"
	"context"
)

type UseCase struct {
	repo repositories.ReadRepository
}

func NewGetByIDHandler(repo repositories.ReadRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (h *UseCase) Handle(ctx context.Context, command Command) (*entities.Profile, error) {
	profile, err := h.repo.GetByID(ctx, command.ID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
