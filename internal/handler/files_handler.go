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

	// Log de entrada da função
	h.Logger.Infof("Handler SaveFiles")

	protocolNumber := c.Param("protocolNumber")
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid form data")
	}
	files := form.File["file"]

	savedFiles, err := h.FilesService.SaveFiles(protocolNumber, files)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to save files")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save files")
	}

	// Log de saída da função
	h.Logger.Infof("Handler SaveFiles OK")
	return c.JSON(http.StatusOK, savedFiles)
}
