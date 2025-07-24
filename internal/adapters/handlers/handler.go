package handlers

import "awesomeProject2/internal/core/usecases/get"

type GetHandler struct {
	getUseCase *get.UseCase
}

func NewGetHandler(getUseCase *get.UseCase) *GetHandler {
	return &GetHandler{getUseCase: getUseCase}
}
