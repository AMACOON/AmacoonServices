package handlers

import (
	"net/http"
	"amacoonservices/repositories/information"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ColorHandler struct {
	ColorRepo *repositories.ColorRepository
	Logger    *logrus.Logger
}

func NewColorHandler(colorRepo *repositories.ColorRepository, logger *logrus.Logger) *ColorHandler {
    return &ColorHandler{
        ColorRepo: colorRepo,
        Logger:    logger,
    }
}

func (h *ColorHandler) GetAllColorsByBreed(c echo.Context) error {
   
    h.Logger.Infof("Handler GetAllColorsByBreed")
    breedID := c.QueryParam("breedID") 

    h.Logger.WithFields(logrus.Fields{
        "breedID": breedID,
    }).Info("GetAllColorsByBreed called")

    colors, err := h.ColorRepo.GetAllColorsByBreed(breedID)
    if err != nil {
        h.Logger.WithError(err).Error("Failed to get colors by breed")
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    h.Logger.Infof("Handler GetAllColorsByBreed OK")
    return c.JSON(http.StatusOK, colors)
}
