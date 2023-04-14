package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/litter"
    "github.com/scuba13/AmacoonServices/internal/utils"
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
	var litter litter.Litter
	err := c.Bind(&litter)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	createdLitter, err := h.LitterService.CreateLitter(litter)
	if err != nil {
		h.Logger.WithError(err).Error("failed to create litter")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create litter")
	}
	return c.JSON(http.StatusCreated, createdLitter)
}

func (h *LitterHandler) GetLitterByID(c echo.Context) error {
	h.Logger.Infof("Handler GetLitterByID")
	id := c.Param("id")

	foundLitter, err := h.LitterService.GetLitterByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get litter")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get litter")
	}
	return c.JSON(http.StatusOK, foundLitter)
}

func (h *LitterHandler) UpdateLitterStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateLitterStatus")
	id := c.Param("id")
	status := c.QueryParam("status")
	err := h.LitterService.UpdateLitterStatus(id, status)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update litter status")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update litter status")
	}
	return c.NoContent(http.StatusOK)
}

func (h *LitterHandler) GetLitterFilesByID(c echo.Context) error {
	h.Logger.Infof("Handler GetLitterFilesByID")
	id := c.Param("id")
	files, err := h.LitterService.GetLitterFilesByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get litter files")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get litter files")
	}
    h.Logger.Infof("Handler GetLitterFilesByID OK")
	return c.JSON(http.StatusOK, files)
}

func (h *LitterHandler) GetAllLittersByOwner(c echo.Context) error {
	h.Logger.Infof("Handler GetAllLittersByOwner")
	id := c.Param("ownerId")

	litters, err := h.LitterService.GetAllLittersByOwner(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get litters by owner id")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get litters by owner id")
	}

	h.Logger.Infof("Handler GetAllLittersByOwner OK")
	return c.JSON(http.StatusOK, litters)
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

func (h *LitterHandler) AddTransferFiles(c echo.Context) error {
	h.Logger.Infof("Handler AddLitterFiles")
	id := c.Param("id")
	var files []utils.Files
	err := c.Bind(&files)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.LitterService.AddLitterFiles(id, files)
	if err != nil {
		h.Logger.WithError(err).Error("failed to add files to litter")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to add files to transfer")
	}

	h.Logger.Infof("Handler AddLitterFiles OK")
	return c.NoContent(http.StatusOK)
}
