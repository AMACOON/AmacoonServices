package handlers

import (
	"amacoonservices/repositories/information"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CatHandler struct {
	CatRepo *repositories.CatRepository
	Logger  *logrus.Logger
}

func NewCatHandler(catRepo *repositories.CatRepository, logger *logrus.Logger) *CatHandler {
	return &CatHandler{
		CatRepo: catRepo,
		Logger:  logger,
	}
}

func getSexAsInt(sex string) (int, error) {
	switch sex {
	case "M":
		return 1, nil
	case "F":
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid sex parameter")
	}
}

func (h *CatHandler) GetCatsByExhibitorAndSex(c echo.Context) error {
	
	h.Logger.Infof("Handler GetCatsByExhibitorAndSex")
	exhibitorID, err := strconv.Atoi(c.QueryParam("id_exhibitor"))
	if err != nil {
		h.Logger.Warn("Invalid id_exhibitor parameter")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id_exhibitor parameter")
	}

	sex := c.QueryParam("sex")
	if sex == "" {
		h.Logger.Warn("Missing sex parameter")
		return echo.NewHTTPError(http.StatusBadRequest, "missing sex parameter")
	}

	sexAsInt, err := getSexAsInt(sex)
	if err != nil {
		h.Logger.Warn("Invalid sex parameter")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	h.Logger.WithFields(logrus.Fields{
		"exhibitorID": exhibitorID,
		"sex":         sexAsInt,
	}).Info("Getting cats by exhibitor and sex")

	cats, err := h.CatRepo.GetCatsByExhibitorAndSex(exhibitorID, sexAsInt)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cats by exhibitor and sex")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler GetCatsByExhibitorAndSex OK")
	return c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetCatByRegistration(c echo.Context) error {
	h.Logger.Infof("Handler GetCatByRegistration")
	registration := c.Param("registration")

	h.Logger.WithFields(logrus.Fields{
		"registration": registration,
	}).Info("Getting cat by registration")

	cat, err := h.CatRepo.GetCatByRegistration(registration)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by registration")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"registration": registration,
		}).Warn("Cat not found by registration")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found")
	}
	h.Logger.Infof("Handler GetCatByRegistration OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) GetCatsByExhibitorAndSexService(c echo.Context) error {
	h.Logger.Infof("Handler GetCatsByExhibitorAndSexService")
	idExhibitor, err := strconv.Atoi(c.QueryParam("id_exhibitor"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id_exhibitor parameter")
	}

	sex := c.QueryParam("sex")
	if sex == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing sex parameter")
	}

	sexAsInt, err := getSexAsInt(sex)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cats, err := h.CatRepo.GetCatsByExhibitorAndSexService(idExhibitor, sexAsInt)
	if err != nil {
		h.Logger.Errorf("Failed to get cats from repository: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get cats from repository")
	}

	h.Logger.Infof("Got %d cats by exhibitor %d and sex %s from repository", len(cats), idExhibitor, sex)
	h.Logger.Infof("Handler GetCatsByExhibitorAndSexService OK")
	return c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetCatByRegistrationService(c echo.Context) error {
	h.Logger.Infof("Handler GetCatByRegistrationService")
	registration := c.Param("registration")

	cat, err := h.CatRepo.GetCatByRegistrationService(registration)
	if err != nil {
		h.Logger.Errorf("Failed to get cat from repository: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get cat from repository")
	}

	if cat == nil {
		h.Logger.Warnf("Cat with registration %s not found", registration)
		return echo.NewHTTPError(http.StatusNotFound, "cat not found")
	}

	h.Logger.Infof("Got cat with registration %s from repository", registration)
	h.Logger.Infof("Handler GetCatByRegistrationService OK")
	return c.JSON(http.StatusOK, cat)
}
