package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/title"
)

type TitleHandler struct {
	TitleService *title.TitleService
	Logger         *logrus.Logger
}

func NewTitleHandler(titleService *title.TitleService, logger *logrus.Logger) *TitleHandler {
	return &TitleHandler{
		TitleService: titleService,
		Logger:         logger,
	}
}

func (h *TitleHandler) GetAllTitles(c echo.Context) error {
	h.Logger.Infof("Handler GetAllTitles")
	titles, err := h.TitleService.GetAllTitles()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all countries")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Retrieved %d titles", len(titles))
	h.Logger.Infof("Handler GetAllTitles OK")
	return c.JSON(http.StatusOK, titles)
}

