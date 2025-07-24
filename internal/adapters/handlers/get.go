package handlers

import (
	"awesomeProject2/internal/core/usecases/get"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	command := get.Command{ID: id}
	profile, err := h.getUseCase.Handle(c.Request.Context(), command)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, profile)
}
