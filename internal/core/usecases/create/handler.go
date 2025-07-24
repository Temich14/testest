package create

import (
	"awesomeProject2/internal/core/entities"
	"awesomeProject2/internal/core/ports/repositories"
	"context"
)

type UseCase struct {
	repo repositories.Repository
}

func NewCreateUserHandler(repo repositories.Repository) UseCase {
	return UseCase{repo: repo}
}
func (u *UseCase) Handle(ctx context.Context, command Command) error {
	profile := entities.Profile{
		ID:        command.ID,
		Email:     command.Email,
		FirstName: command.FirstName,
		LastName:  command.LastName,
	}
	err := u.repo.Add(ctx, profile)
	if err != nil {
		return err
	}
	return nil
}
