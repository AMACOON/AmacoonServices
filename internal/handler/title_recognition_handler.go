package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/titlerecognition"
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
	var titleRecognitionReq titlerecognition.TitleRecognitionRequest
	err := c.Bind(&titleRecognitionReq)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	createdTitleRecognition, err := h.TitleRecognitionService.CreateTitleRecognition(titleRecognitionReq)
	if err != nil {
		h.Logger.WithError(err).Error("failed to create title recognition")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create title recognition")
	}
	return c.JSON(http.StatusCreated, createdTitleRecognition)
}

func (h *TitleRecognitionHandler) GetTitleRecognitionByID(c echo.Context) error {
	h.Logger.Infof("Handler GetTitleRecognitionByID")
	id := c.Param("id")

	foundTitleRecognition, err := h.TitleRecognitionService.GetTitleRecognitionByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get title recognition")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get title recognition")
	}
	h.Logger.Infof("Handler GetTitleRecognitionByID OK")
	return c.JSON(http.StatusOK, foundTitleRecognition)
}

func (h *TitleRecognitionHandler) GetTitleRecognitionFilesByID(c echo.Context) error {
	h.Logger.Infof("Handler GetTitleRecognitionFilesByID")
	id := c.Param("id")
	files, err := h.TitleRecognitionService.GetTitleRecognitionFilesByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get title recognition files")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get title recognition files")
	}
	h.Logger.Infof("Handler GetTitleRecognitionFilesByID OK")
	return c.JSON(http.StatusOK, files)
}

func (h *TitleRecognitionHandler) GetAllTitleRecognitionsByRequesterID(c echo.Context) error {
	h.Logger.Infof("Handler GetAllTitleRecognitionsByRequesterID")
	id := c.Param("requesterID")

	titleRecognitions, err := h.TitleRecognitionService.GetAllTitleRecognitionsByRequesterID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get title recognitions by Requester ID")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get title recognitions by Requester ID")
	}

	h.Logger.Infof("Handler GetAllTitleRecognitionsByRequesterID OK")
	return c.JSON(http.StatusOK, titleRecognitions)
}

func (h *TitleRecognitionHandler) UpdateTitleRecognitionStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTitleRecognitionStatus")
	id := c.Param("id")
	status := c.QueryParam("status")
	err := h.TitleRecognitionService.UpdateTitleRecognitionStatus(id, status)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update title recognition status")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update title recognition status")
	}
	h.Logger.Infof("Handler UpdateTitleRecognitionStatus OK")
	return c.NoContent(http.StatusOK)
}

func (h *TitleRecognitionHandler) AddTitlesReconitionFiles(c echo.Context) error {
	h.Logger.Infof("Handler AddTitlesReconitionFiles")
	id := c.Param("id")
	var files []utils.Files
	err := c.Bind(&files)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.TitleRecognitionService.AddTitleRecognitionFiles(id, files)
	if err != nil {
		h.Logger.WithError(err).Error("failed to add files to title recognition")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to add files to title recognition")
	}

	h.Logger.Infof("Handler AddTitlesReconitionFiles OK")
	return c.NoContent(http.StatusOK)
}

func (h *TitleRecognitionHandler) UpdateTitlesRecognition(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTitlesRecognition")
	id := c.Param("id")
	var titleRecognitionObj titlerecognition.TitleRecognitionMongo
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

