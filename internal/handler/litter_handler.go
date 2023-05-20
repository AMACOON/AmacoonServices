package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/litter"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"encoding/json"
)

type LitterHandler struct {
	LitterService *litter.LitterService
	Logger        *logrus.Logger
}

func NewLitterHandler(litterService *litter.LitterService, logger *logrus.Logger) *LitterHandler {
	return &LitterHandler{
		LitterService: litterService,
		Logger:        logger,
	}
}

func (h *LitterHandler) CreateLitter(c echo.Context) error {
	h.Logger.Infof("Handler CreateLitter")

	// Get multipart form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Extract litter JSON
	litterJson := form.Value["litter"][0]
	litter := &litter.Litter{}
	err = json.Unmarshal([]byte(litterJson), litter)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse litter JSON")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate litter
	if err := utils.ValidateStruct(litter); err != nil {
		h.Logger.WithError(err).Error("Failed to validate litter")
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

	// Create litter
	litter, err = h.LitterService.CreateLitter(*litter, filesWithDesc)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to create litter")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler CreateLitter OK")
	return c.JSON(http.StatusCreated, litter)
}


func (h *LitterHandler) GetLitterByID(c echo.Context) error {
	h.Logger.Infof("Handler GetLitterByID")
	id := c.Param("id")

	foundLitter, err := h.LitterService.GetLitterByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get litter")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler GetLitterByID OK")
	return c.JSON(http.StatusOK, foundLitter)
}

func (h *LitterHandler) UpdateLitter(c echo.Context) error {
	h.Logger.Infof("Handler UpdateLitter")
	id := c.Param("id")
	var litter litter.Litter
	err := c.Bind(&litter)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.LitterService.UpdateLitter(id, litter)
	if err != nil {
		h.Logger.WithError(err).Errorf("failed to update litter with id %s", id)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler UpdateLitter OK")
	return c.NoContent(http.StatusOK)
}

func (h *LitterHandler) UpdateLitterStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateLitterStatus")
	id := c.Param("id")
	status := c.QueryParam("status")
	err := h.LitterService.UpdateLitterStatus(id, status)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update litter status")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler UpdateLitterStatus OK")
	return c.NoContent(http.StatusOK)
}

func (h *LitterHandler) GetAllLittersByRequesterID(c echo.Context) error {
	h.Logger.Infof("Handler GetAllLittersByRequesterID")
	requesterID := c.Param("requesterID")

	litters, err := h.LitterService.GetAllLittersByRequesterID(requesterID)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get litters by Requester ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler GetAllLittersByRequesterID OK")
	return c.JSON(http.StatusOK, litters)
}

func (h *LitterHandler) DeleteLitter(c echo.Context) error {
	h.Logger.Infof("Handler DeleteLitter")
	id := c.Param("id")
	
	err := h.LitterService.DeleteLitter(id)
	if err != nil {
		h.Logger.WithError(err).Errorf("failed to delete litter with id %s", id)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler DeleteLitter OK")
	return c.NoContent(http.StatusOK)
}



