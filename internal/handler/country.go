package handler

import (
	"github.com/scuba13/AmacoonServices/internal/country"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CountryHandler struct {
	CountryRepo *country.CountryRepository
	Logger      *logrus.Logger
}

func NewCountryHandler(countryRepo *country.CountryRepository, logger *logrus.Logger) *CountryHandler {
	return &CountryHandler{
		CountryRepo: countryRepo,
		Logger:      logger,
	}
}

func (h *CountryHandler) GetAllCountry(c echo.Context) error {
	h.Logger.Infof("Handler GetAllCountry")
	countries, err := h.CountryRepo.GetAllCountries()
	if err != nil {
		h.Logger.Errorf("Failed to get all countries: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Retrieved %d countries", len(countries))
	h.Logger.Infof("Handler GetAllCountry OK")
	return c.JSON(http.StatusOK, countries)
}
