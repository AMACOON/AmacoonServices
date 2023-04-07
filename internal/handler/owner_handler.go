package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/owner"
)

type OwnerHandler struct {
	OwnerService *owner.OwnerService
	Logger       *logrus.Logger
}

func NewOwnerHandler(ownerService *owner.OwnerService, logger *logrus.Logger) *OwnerHandler {
	return &OwnerHandler{
		OwnerService: ownerService,
		Logger:       logger,
	}
}

func (h *OwnerHandler) GetOwnerByID(c echo.Context) error {
	id := c.Param("id")
	ownerID, err := strconv.Atoi(id)
	if err != nil {
		h.Logger.WithError(err).Warn("Failed to parse owner ID")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid owner ID")
	}

	h.Logger.WithField("ownerID", ownerID).Info("Getting owner by ID")

	owner, err := h.OwnerService.GetOwnerByID(ownerID)
	if err != nil {
		h.Logger.WithError(err).Warn("Failed to get owner from service")
		return echo.NewHTTPError(http.StatusInternalServerError, "owner not found")
	}
	if owner.OwnerName == "" {
		h.Logger.Info("Owner not found")
		 return c.JSON(http.StatusOK, "Owner not found")
	}

	h.Logger.WithField("ownerID", ownerID).Info("Successfully retrieved owner by ID")
	return c.JSON(http.StatusOK, owner)
}
