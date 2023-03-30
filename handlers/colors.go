package handlers

import (
	"net/http"
	"amacoonservices/repositories"

	"github.com/labstack/echo/v4"
)

type ColorHandler struct {
	ColorRepo repositories.ColorRepository
}

func (h *ColorHandler) GetAllColorsByBreed(c echo.Context) error {
    breedID := c.QueryParam("breedID") 

    colors, err := h.ColorRepo.GetAllColorsByBreed(breedID)
    if err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, colors)
}
