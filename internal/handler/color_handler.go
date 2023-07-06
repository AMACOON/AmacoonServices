package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/color"
)

type ColorHandler struct {
	ColorService *color.ColorService
	Logger       *logrus.Logger
}

func NewColorHandler(colorService *color.ColorService, logger *logrus.Logger) *ColorHandler {
	return &ColorHandler{
		ColorService: colorService,
		Logger:       logger,
	}
}

func (h *ColorHandler) GetAllColorsByBreed(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetAllColorsByBreed")
	breedCode := c.Param("breedCode")
	
	h.Logger.WithFields(logrus.Fields{
		"breedCode": breedCode,
	}).Info("Getting Colors by Breed")
	
	colors, err := h.ColorService.GetAllColorsByBreed(breedCode)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get colors by breed")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	
	// Log de saída da função
	h.Logger.Infof("Handler GetAllColorsByBreed OK")
	return c.JSON(http.StatusOK, colors)
	}

func (h *ColorHandler) UpdateColor(c echo.Context) error {
	h.Logger.Infof("Handler UpdateColor")
	id := c.Param("id")
	
	var color color.Color
	if err := c.Bind(&color); err != nil {
		h.Logger.WithError(err).Error("failed to bind color")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.ColorService.UpdateColor(id, &color)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update color status")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler UpdateColor OK")
	return c.NoContent(http.StatusOK)
}

func (h *ColorHandler) GetColorByID(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetColorByID")
	id := c.Param("id")

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting Color by ID")

	color, err := h.ColorService.GetColorById(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get Color by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if color == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("Color not found by ID")
		return echo.NewHTTPError(http.StatusNotFound, "Color not found by ID")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetColorByID OK")
	return c.JSON(http.StatusOK, color)
}
	