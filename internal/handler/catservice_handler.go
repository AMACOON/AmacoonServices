package handler

import (
	"net/http"

	"github.com/scuba13/AmacoonServices/internal/catservice"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CatServiceHandler struct {
	CatServiceService *catservice.CatServiveService
	Logger     *logrus.Logger
}

func NewCatServiceHandler(catServiceService *catservice.CatServiveService, logger *logrus.Logger) *CatServiceHandler {
	return &CatServiceHandler{
		CatServiceService: catServiceService,
		Logger:     logger,
	}
}


func (h *CatServiceHandler) GetCatServiceByID(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetCatServiceByID")
	id := c.Param("id")

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting CatService by ID")

	cat, err := h.CatServiceService.GetCatServiceByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get CatService by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("CatService not found by ID")
		return echo.NewHTTPError(http.StatusNotFound, "CatService not found")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetCatServiceByID OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatServiceHandler) GetAllCatsServiceByOwnerAndGender(c echo.Context) error {
	h.Logger.Infof("Handler GetAllCatsServiceByOwnerAndGender")
	ownerId := c.QueryParam("ownerId")
	gender := c.QueryParam("gender")

	h.Logger.WithFields(logrus.Fields{
		"OwnerId": ownerId,
		"Gender": gender,
	}).Info("Getting cat by OwnerID and Gender")

	cat, err := h.CatServiceService.GetAllCatsServiceByOwnerAndGender(ownerId, gender)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by OwnerID and Gender")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"OwnerId": ownerId,
			"Gender": gender,
		}).Warn("Cat not found by OwnerID and Gender")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found by OwnerID and Gender")
	}
	h.Logger.Infof("Handler GetAllCatsServiceByOwnerAndGender OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatServiceHandler) GetAllCatsServiceByOwner(c echo.Context) error {
	h.Logger.Infof("Handler GetAllCatsServiceByOwner")
	ownerId := c.Param("ownerId")

	h.Logger.WithFields(logrus.Fields{
		"OwnerId": ownerId,
		
	}).Info("Getting cat by OwnerID")

	cat, err := h.CatServiceService.GetAllCatsServiceByOwner(ownerId)
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
	h.Logger.Infof("Handler GetAllCatsServiceByOwner OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatServiceHandler) GetCatServiceByRegistration(c echo.Context) error {
	h.Logger.Infof("Handler GetCatServiceByRegistration")
	resgistration := c.Param("registration")
	

	h.Logger.WithFields(logrus.Fields{
		"Resgistration": resgistration,
	}).Info("Getting cat by Registration")

	cat, err := h.CatServiceService.GetCatServiceByRegistration(resgistration)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by Registration")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"Resgistration": resgistration,
		
		}).Warn("Cat not found by Registration")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found by Registration")
	}
	h.Logger.Infof("Handler GetCatServiceByRegistration OK")
	return c.JSON(http.StatusOK, cat)
}


