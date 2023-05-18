package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"strconv"
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
		h.Logger.WithError(err).Error("Failed to get all breeds")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetAllBreeds OK")
	return c.JSON(http.StatusOK, breeds)
}

func (h *BreedHandler) GetBreedByID(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetBreedByID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse Breed ID")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Breed ID")
	}

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting Breed by ID")

	breed, err := h.BreedService.GetBreedByID(uint(id))
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get Breed by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if breed == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("Breed not found by ID")
		return echo.NewHTTPError(http.StatusNotFound, "breed not found")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetBreedByID OK")
	return c.JSON(http.StatusOK, breed)
}
