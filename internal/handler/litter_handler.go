package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/litter"
	"strconv"
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

func (h *LitterHandler) GetAllLitters(c echo.Context) error {
	h.Logger.Info("Handler GetAllLitters")

	litters, err := h.LitterService.GetAllLitters()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all litters")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Info("Handler GetAllLitters OK")
	return c.JSON(http.StatusOK, litters)
}

func (h *LitterHandler) GetLitterByID(c echo.Context) error {
	litterIDStr := c.Param("id")
	h.Logger.Infof("Handler GetLitterByID - litter ID: %s", litterIDStr)

	litterID, err := strconv.ParseUint(litterIDStr, 10, 64)
	if err != nil {
		h.Logger.WithError(err).Error("Invalid litter ID")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid litter ID")
	}

	// Call the service to get the litter data
	litterData, err := h.LitterService.GetLitterByID(id string())
	if err != nil {
		h.Logger.WithError(err).Errorf("Failed to get litter by ID %v", litterID)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the LitterData as a response
	h.Logger.Info("Handler GetLitterByID OK")
	return c.JSON(http.StatusOK, litterData)
}


func (h *LitterHandler) CreateLitter(c echo.Context) error {
    h.Logger.Info("Handler CreateLitter")

    var litterData litter.Litter
    if err := c.Bind(&litterData); err != nil {
        h.Logger.WithError(err).Error("Failed to parse request body")
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    litterID, protocolNumber, err := h.LitterService.CreateLitter(litterData)
    if err != nil {
        h.Logger.WithError(err).Error("Failed to create litter")
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    h.Logger.Info("Handler CreateLitter OK")
    return c.JSON(http.StatusOK, map[string]string{
        "litter_id": strconv.Itoa(int(litterID)),
        "protocol":  protocolNumber,
    })
}

func (h *LitterHandler) UpdateLitter(c echo.Context) error {
    litterIDStr := c.Param("id")
    h.Logger.Infof("Handler UpdateLitter - litter ID: %s", litterIDStr)

    // Parse the LitterID from the request params
    litterID, err := strconv.ParseUint(litterIDStr, 10, 64)
    if err != nil {
        h.Logger.WithError(err).Error("Failed to parse litter ID")
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid litter ID")
    }

    // Parse the request body into a Litter struct
    var litter litter.Litter
    if err := c.Bind(&litter); err != nil {
        h.Logger.WithError(err).Error("Failed to parse request body")
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    // Call the service to update the litter and its kittens
    if err := h.LitterService.UpdateLitter(uint(litterID), litter); err != nil {
        h.Logger.WithError(err).Error("Failed to update litter")
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    // Return success response
    h.Logger.Info("Handler UpdateLitter OK")
    return c.NoContent(http.StatusOK)
}

func (h *LitterHandler) DeleteLitter(c echo.Context) error {
    litterIDStr := c.Param("id")
    h.Logger.Infof("Handler DeleteLitter - litter ID: %s", litterIDStr)

    // Parse the LitterID from the request params
    litterID, err := strconv.ParseUint(litterIDStr, 10, 64)
    if err != nil {
        h.Logger.WithError(err).Error("Failed to parse litter ID")
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid litter ID")
    }

    // Call the service to delete the litter and its kittens
    if err := h.LitterService.DeleteLitter(uint(litterID)); err != nil {
        h.Logger.WithError(err).Error("Failed to delete litter")
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    // Return success response
    h.Logger.Info("Handler DeleteLitter OK")
    return c.NoContent(http.StatusOK)
}


