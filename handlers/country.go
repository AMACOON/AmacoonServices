package handlers

import (
	"amacoonservices/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CountryHandler struct {
	CountryRepo *repositories.CountryRepository
}

func (h *CountryHandler) GetAllCountry(c echo.Context) error {
	country, err := h.CountryRepo.GetAllCountries()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, country)
}
