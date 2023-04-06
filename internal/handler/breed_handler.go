package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/breed"
)

type BreedHandler struct {
	BreedService *breed.BreedService
	Logger       *logrus.Logger
}

func NewBreedHandler(breedService *breed.BreedService, logger *logrus.Logger) *BreedHandler {
	return &BreedHandler{
		BreedService: breedService,
		Logger:       logger,
	}
}

func (h *BreedHandler) GetAllBreeds(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetAllBreeds")

	breeds, err := h.BreedService.GetAllBreeds()
	if err != nil {
		h.Logger.Errorf("Erro ao obter todas as raças: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetAllBreeds OK")
	return c.JSON(http.StatusOK, breeds)
}

func (h *BreedHandler) GetCompatibleBreeds(c echo.Context, breedID string) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetCompatibleBreeds")

	breeds, err := h.BreedService.GetCompatibleBreeds(breedID)
	if err != nil {
		h.Logger.Errorf("Erro ao obter raças compatíveis: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(breeds) == 0 {
		h.Logger.Infof("A raça %s não tem raças compatíveis", breedID)
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetCompatibleBreeds OK")
	return c.JSON(http.StatusOK, breeds)
}
