package handlers

import (
    "amacoonservices/repositories/information"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/sirupsen/logrus"
)

type OwnerHandler struct {
    OwnerRepo *repositories.OwnerRepository
    Logger    *logrus.Logger
}

func NewOwnerHandler(ownerRepo *repositories.OwnerRepository, logger *logrus.Logger) *OwnerHandler {
    return &OwnerHandler{
        OwnerRepo: ownerRepo,
        Logger:    logger,
    }
}

func (h *OwnerHandler) GetOwnerByID(c echo.Context) error {
    h.Logger.Info("Handler GetOwnerByID")
	id := c.Param("id")
    ownerID, err := strconv.Atoi(id)
    if err != nil {
        h.Logger.Error("Failed to parse owner ID: ", err)
        return c.String(http.StatusBadRequest, "invalid owner ID")
    }
	h.Logger.Info("Getting owner with ID:", ownerID)
    owner, err := h.OwnerRepo.GetOwnerByExhibitorID(uint(ownerID))
    if err != nil {
        h.Logger.Error("Failed to get owner from repository: ", err)
        return c.String(http.StatusNotFound, "owner not found")
    }

    h.Logger.Info("Successfully retrieved owner with ID ", ownerID)
	h.Logger.Info("Handler GetOwnerByID OK")
    return c.JSON(http.StatusOK, owner)
}
