package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/titlerecognition"
	"encoding/json"
	"github.com/scuba13/AmacoonServices/internal/utils"
)

type TitleRecognitionHandler struct {
	TitleRecognitionService *titlerecognition.TitleRecognitionService
	Logger                  *logrus.Logger
}

func NewTitleRecognitionHandler(titleRecognitionService *titlerecognition.TitleRecognitionService, logger *logrus.Logger) *TitleRecognitionHandler {
	return &TitleRecognitionHandler{
		TitleRecognitionService: titleRecognitionService,
		Logger:                  logger,
	}
}

func (h *TitleRecognitionHandler) CreateTitleRecognition(c echo.Context) error {
	h.Logger.Infof("Handler CreateTitleRecognition")
	// Get multipart form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Extract titlerecognition JSON
	titleRecognitionJson := form.Value["titlerecognition"][0]
	titleRecognition := &titlerecognition.TitleRecognition{}
	err = json.Unmarshal([]byte(titleRecognitionJson), titleRecognition)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse titleRecognition JSON")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate titlerecognition
	if err := utils.ValidateStruct(titleRecognition); err != nil {
		h.Logger.WithError(err).Error("Failed to validate titleRecognition")
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

	// Create titleRecognition
	titleRecognition, err = h.TitleRecognitionService.CreateTitleRecognition(*titleRecognition, filesWithDesc)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to create titleRecognition")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler CreateTitleRecognition OK")
	return c.JSON(http.StatusCreated, titleRecognition)
}

func (h *TitleRecognitionHandler) GetTitleRecognitionByID(c echo.Context) error {
	h.Logger.Infof("Handler GetTitleRecognitionByID")
	id := c.Param("id")

	foundTitleRecognition, err := h.TitleRecognitionService.GetTitleRecognitionByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get title recognition")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler GetTitleRecognitionByID OK")
	return c.JSON(http.StatusOK, foundTitleRecognition)
}

func (h *TitleRecognitionHandler) UpdateTitlesRecognition(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTitlesRecognition")
	id := c.Param("id")
	var titleRecognitionObj titlerecognition.TitleRecognition
	err := c.Bind(&titleRecognitionObj)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.TitleRecognitionService.UpdateTitleRecognition(id, titleRecognitionObj)
	if err != nil {
		h.Logger.WithError(err).Errorf("failed to update title recognition with id %s", id)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler UpdateTitlesRecognition OK")
	return c.NoContent(http.StatusOK)
}

func (h *TitleRecognitionHandler) UpdateTitleRecognitionStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTitleRecognitionStatus")
	id := c.Param("id")
	status := c.QueryParam("status")
	err := h.TitleRecognitionService.UpdateTitleRecognitionStatus(id, status)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update title recognition status")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler UpdateTitleRecognitionStatus OK")
	return c.NoContent(http.StatusOK)
}

func (h *TitleRecognitionHandler) GetAllTitleRecognitionsByRequesterID(c echo.Context) error {
	h.Logger.Infof("Handler GetAllTitleRecognitionsByRequesterID")
	id := c.Param("requesterID")

	titleRecognitions, err := h.TitleRecognitionService.GetAllTitleRecognitionsByRequesterID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get title recognitions by Requester ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler GetAllTitleRecognitionsByRequesterID OK")
	return c.JSON(http.StatusOK, titleRecognitions)
}

func (h *TitleRecognitionHandler) DeleteTitleRecognition(c echo.Context) error {
	h.Logger.Infof("Handler DeleteTitleRecognition")
	
	id := c.Param("id")
	err := h.TitleRecognitionService.DeleteTitleRecognition(id)
	if err != nil {
		h.Logger.WithError(err).Errorf("failed to delete title recognition with id %s", id)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler DeleteTitleRecognition OK")
	return c.NoContent(http.StatusOK)
}



