package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"encoding/json"
)

type CatShowRegistrationHandler struct {
	CatShowRegistrationService *catshowregistration.CatShowRegistrationService
	Logger                     *logrus.Logger
}

func NewCatShowRegistrationHandler(catShowRegistrationService *catshowregistration.CatShowRegistrationService, logger *logrus.Logger) *CatShowRegistrationHandler {
	return &CatShowRegistrationHandler{
		CatShowRegistrationService: catShowRegistrationService,
		Logger:                     logger,
	}
}

func (h *CatShowRegistrationHandler) CreateCatShowRegistration(c echo.Context) error {
	h.Logger.Infof("Handler CreateCatShowRegistration")

	// Get multipart form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Extract registration JSON
	registrationJson := form.Value["registration"][0]
	registration := &catshowregistration.Registration{}
	err = json.Unmarshal([]byte(registrationJson), registration)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse registration JSON")
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

	createdRegistration, err := h.CatShowRegistrationService.CreateCatShowRegistration(registration, filesWithDesc)
	if err != nil {
		h.Logger.Errorf("Failed to create CatShowRegistration: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create CatShowRegistration")
	}

	h.Logger.Infof("Handler CreateCatShowRegistration OK")
	return c.JSON(http.StatusCreated, createdRegistration)
}
