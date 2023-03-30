package handlers

import (
	"amacoonservices/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OwnerHandler struct {
	OwnerRepo repositories.OwnerRepository
}

func (h *OwnerHandler) GetOwnerByID(c echo.Context) error {
	id := c.Param("id")
	ownerID, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid owner ID")
	}

	owner, err := h.OwnerRepo.GetOwnerByExhibitorID(uint(ownerID))
	if err != nil {
		return c.String(http.StatusNotFound, "owner not found")
	}

	return c.JSON(http.StatusOK, owner)
}
