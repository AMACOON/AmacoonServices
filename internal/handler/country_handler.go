package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/country"
)

type CountryHandler struct {
	CountryService *country.CountryService
	Logger         *logrus.Logger
}

func NewCountryHandler(countryService *country.CountryService, logger *logrus.Logger) *CountryHandler {
	return &CountryHandler{
		CountryService: countryService,
		Logger:         logger,
	}
}

func (h *CountryHandler) GetAllCountry(c echo.Context) error {
	h.Logger.Infof("Handler GetAllCountry")
	countries, err := h.CountryService.GetAllCountries()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all countries")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Retrieved %d countries", len(countries))
	h.Logger.Infof("Handler GetAllCountry OK")
	return c.JSON(http.StatusOK, countries)
}

