package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/cattery"
)

type CatteryHandler struct {
	CatteryService *cattery.CatteryService
	Logger       *logrus.Logger
}

func NewCatteryHandler(catteryService *cattery.CatteryService, logger *logrus.Logger) *CatteryHandler {
	return &CatteryHandler{
		CatteryService: catteryService,
		Logger:       logger,
	}
}

func (h *CatteryHandler) GetAllCatteries(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetAllCatteries")

	catteries, err := h.CatteryService.GetAllCatteries()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all catteries")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get all catteries")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetAllCatteries OK")
	return c.JSON(http.StatusOK, catteries)
}

func (h *CatteryHandler) GetCatteryByID(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetCatteryByID")
	id := c.Param("id")
	

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting Cattery by ID")

	cattery, err := h.CatteryService.GetCatteryByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get Cattery by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	if cattery == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("Cattery not found by ID")
		return echo.NewHTTPError(http.StatusNotFound, "cattery not found")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetCatteryByID OK")
	return c.JSON(http.StatusOK, cattery)
}

