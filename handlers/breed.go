package handlers

import (
	"amacoonservices/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BreedHandler struct {
	BreedRepo *repositories.BreedRepository
}

func (h *BreedHandler) GetAllBreeds(c echo.Context) error {
	breeds, err := h.BreedRepo.GetAllBreeds()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, breeds)
}
func (h *BreedHandler) GetCompatibleBreeds(c echo.Context, BreedID string) error {
	logger := c.Logger()
	breeds, err := h.BreedRepo.GetCompatibleBreeds(BreedID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if len(breeds) == 0 {
		logger.Info("raça", BreedID, "não tem raças compatíveis")
	}

	return c.JSON(http.StatusOK, breeds)
}
