package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	
)

type FilesHandler struct {
	FilesService *utils.FilesService
	Logger       *logrus.Logger
}

func NewFilesHandler(filesService *utils.FilesService, logger *logrus.Logger) *FilesHandler {
	return &FilesHandler{
		FilesService: filesService,
		Logger:       logger,
	}
}

func (h *FilesHandler) SaveFiles(c echo.Context) error {
	h.Logger.Infof("Handler SaveFiles")

	// Get form values
	identifier := c.FormValue("identifier")
	domain := c.FormValue("domain")

	// Get files from form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

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

	savedFiles, err := h.FilesService.SaveFiles(identifier, domain, filesWithDesc)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to save files")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log out from function
	h.Logger.Infof("Handler SaveFiles OK")
	return c.JSON(http.StatusOK, savedFiles)
}
