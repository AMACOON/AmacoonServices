package handler

import (
	"net/http"

	"github.com/scuba13/AmacoonServices/internal/cat"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CatHandler struct {
	CatService *cat.CatService
	Logger     *logrus.Logger
}

func NewCatHandler(catService *cat.CatService, logger *logrus.Logger) *CatHandler {
	return &CatHandler{
		CatService: catService,
		Logger:     logger,
	}
}


func (h *CatHandler) GetCatsCompleteByID(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetCatsCompleteByID")
	id := c.Param("id")

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting CatComplete by ID")

	cat, err := h.CatService.GetCatsCompleteByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get CatComplete by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("CatComplete not found by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetCatsCompleteByID OK")
	return c.JSON(http.StatusOK, cat)
}



func (h *CatHandler) GetCatCompleteByAllByOwner(c echo.Context) error {
	h.Logger.Infof("Handler GetCatCompleteByAllByOwner")
	ownerId := c.Param("ownerId")

	h.Logger.WithFields(logrus.Fields{
		"OwnerId": ownerId,
		
	}).Info("Getting cat by OwnerID")

	cat, err := h.CatService.GetCatCompleteByAllByOwner(ownerId)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by OwnerID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"OwnerId": ownerId,
		}).Warn("Cat not found by OwnerID")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found by OwnerID")
	}
	h.Logger.Infof("Handler GetCatCompleteByAllByOwner OK")
	return c.JSON(http.StatusOK, cat)
}
