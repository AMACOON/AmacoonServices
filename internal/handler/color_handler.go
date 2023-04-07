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
	h.Logger.Infof("Handler GetAllColorsByBreed")

	breedID := c.QueryParam("breedID")
	h.Logger.WithField("breedID", breedID).Info("GetAllColorsByBreed called")

	colors, err := h.ColorService.GetAllColorsByBreed(breedID)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get colors by breed")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler GetAllColorsByBreed OK")
	return c.JSON(http.StatusOK, colors)
}
