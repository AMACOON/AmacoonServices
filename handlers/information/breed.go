package handlers

import (
	"amacoonservices/repositories/information"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type BreedHandler struct {
	BreedRepo *repositories.BreedRepository
	Logger    *logrus.Logger
}

func NewBreedHandler(breedRepo *repositories.BreedRepository, logger *logrus.Logger) *BreedHandler {
	return &BreedHandler{
		BreedRepo: breedRepo,
		Logger:    logger,
	}
}

func (h *BreedHandler) GetAllBreeds(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetAllBreeds")

	breeds, err := h.BreedRepo.GetAllBreeds()
	if err != nil {
		h.Logger.Errorf("Erro ao obter todas as raças: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetAllBreeds OK")
	return c.JSON(http.StatusOK, breeds)
}

func (h *BreedHandler) GetCompatibleBreeds(c echo.Context, BreedID string) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetCompatibleBreeds")

	breeds, err := h.BreedRepo.GetCompatibleBreeds(BreedID)
	if err != nil {
		h.Logger.Errorf("Erro ao obter raças compatíveis: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(breeds) == 0 {
		h.Logger.Infof("A raça %s não tem raças compatíveis", BreedID)
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetCompatibleBreeds OK")
	return c.JSON(http.StatusOK, breeds)
}
