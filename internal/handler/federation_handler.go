package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/sirupsen/logrus"
)

type FederationHandler struct {
	FederationService *federation.FederationService
	Logger            *logrus.Logger
}

func NewFederationHandler(federationService *federation.FederationService, logger *logrus.Logger) *FederationHandler {
	return &FederationHandler{
		FederationService: federationService,
		Logger:            logger,
	}
}

func (h *FederationHandler) GetAllFederations(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetAllFederations")

	federations, err := h.FederationService.GetAllFederations()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all federations")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetAllFederations OK")
	return c.JSON(http.StatusOK, federations)
}

func (h *FederationHandler) GetFederationByID(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetFederationByID")
	id := c.Param("id")

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting Cattery by ID")

	federation, err := h.FederationService.GetFederationByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get Federation by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if federation == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("Federation not found by ID")
		return echo.NewHTTPError(http.StatusNotFound, "federation not found")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetFederationByID OK")
	return c.JSON(http.StatusOK, federation)
}
