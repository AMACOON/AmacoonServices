package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"encoding/json"
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

func (h *CatteryHandler) CreateCattery(c echo.Context) error {
	
	// Log de entrada da função
	h.Logger.Infof("Handler CreateCattery")

	// Get multipart form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Extract cat JSON
	catteryJson := form.Value["cattery"][0]
	cattery := &cattery.Cattery{}
	err = json.Unmarshal([]byte(catteryJson), cattery)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse cattery JSON")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate cattery
	if err := utils.ValidateStruct(cattery); err != nil {
		h.Logger.WithError(err).Error("Failed to validate cattery")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Extract files
	files := form.File["file"]
	descriptions := form.Value["description"]

	var filesWithDesc []utils.FileWithDescription
	for i, file := range files {
		description := ""
		if i < len(descriptions) {
			description = descriptions[i]
		}
		filesWithDesc = append(filesWithDesc, utils.FileWithDescription{
			File:        file,
			Description: description,
		})
	}

	// Create cattery
	cattery, err = h.CatteryService.CreateCattery(cattery, filesWithDesc)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to create cattery")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler CreateCattery OK")
	return c.JSON(http.StatusCreated, cattery)
}

func (h *CatteryHandler) UpdateCattery(c echo.Context) error {
	
	// Log de entrada da função
	h.Logger.Infof("Handler UpdateCattery")

	id := c.Param("id")
	cattery := new(cattery.Cattery)
	if err := c.Bind(cattery); err != nil {
		h.Logger.WithError(err).Error("Failed to bind Cattery")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cattery, err := h.CatteryService.UpdateCattery(id, cattery)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to update Cattery")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler UpdateCattery OK")
	return c.JSON(http.StatusOK, cattery)
}

