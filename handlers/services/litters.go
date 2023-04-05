package handlers

import (
	"amacoonservices/handlers/services/converter"
	"amacoonservices/models/services"
	"amacoonservices/repositories/services"

	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LitterHandler struct {
	LitterRepo *repositories.LitterRepository
	Logger     *logrus.Logger
}

func NewLitterHandler(litterRepo *repositories.LitterRepository, logger *logrus.Logger) *LitterHandler {
	return &LitterHandler{
		LitterRepo: litterRepo,
		Logger:     logger,
	}
}

func (h *LitterHandler) GetAllLitters(c echo.Context) error {
	h.Logger.Info("Handler GetAllLitters")
	litters, err := h.LitterRepo.GetAllLitters()
	if err != nil {
		h.Logger.Errorf("Failed to get all litters: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var litterDatas []models.Litter

	// Transform each LitterDB and its KittensDB into a Litter struct
	for _, litter := range litters {
		kittens, err := h.LitterRepo.GetKittensByLitterID(litter.ID)
		if err != nil {
			h.Logger.Errorf("Failed to get kittens by litter ID %v: %v", litter.ID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		litterData := converter.LitterDBToLitter(&litter, kittens)
		litterDatas = append(litterDatas, *litterData)
	}
	h.Logger.Info("Handler GetAllLitters OK")
	return c.JSON(http.StatusOK, litterDatas)
}

func (h *LitterHandler) GetLitterByID(c echo.Context) error {
	litterIDStr := c.Param("id")
	h.Logger.Infof("Handler GetLitterByID - litter ID: %s", litterIDStr)

	litterID, err := strconv.ParseUint(litterIDStr, 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid litter ID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid litter ID")
	}

	// Call the repository to get the litter and its kittens
	litter, kittens, err := h.LitterRepo.GetLitterByID(uint(litterID))
	if err != nil {
		h.Logger.Errorf("Failed to get litter by ID %v: %v", litterID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Transform the models.Litter and []*models.Kitten into a models.LitterData struct
	litterData := converter.LitterDBToLitter(litter, kittens)

	// Return the LitterData as a response
	h.Logger.Info("Handler GetLitterByID OK")
	return c.JSON(http.StatusOK, litterData)
}

func (h *LitterHandler) CreateLitter(c echo.Context) error {

	h.Logger.Info("Handler CreateLitter")

	// Parse the request body into a LitterData struct
	var litterData models.Litter
	if err := c.Bind(&litterData); err != nil {
		h.Logger.Error("Failed to parse request body:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Transform LitterData into a models.Litter struct
	litter, kittens := converter.LitterToLitterDB(litterData)

	// Call the repository to create the litter and its kittens
	litterID, protocolNumber, err := h.LitterRepo.CreateLitter(&litter, kittens)
	if err != nil {
		h.Logger.Error("Failed to create litter:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the LitterID as a response
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
		h.Logger.Error("Failed to parse litter ID:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid litter ID")
	}
	// Parse the request body into a LitterData struct
	var litterData models.Litter
	if err := c.Bind(&litterData); err != nil {
		h.Logger.Error("Failed to parse request body:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Transform Litter into a models.LitterDB struct
	litter, kittens := converter.LitterToLitterDB(litterData)

	// Call the repository to update the litter and its kittens
	err = h.LitterRepo.UpdateLitter(uint(litterID), &litter, kittens)
	if err != nil {
		h.Logger.Error("Failed to update litter:", err)
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
		h.Logger.Error("Failed to parse litter ID:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid litter ID")
	}

	// Call the repository to delete the litter and its kittens
	err = h.LitterRepo.DeleteLitter(uint(litterID))
	if err != nil {
		h.Logger.Error("Failed to delete litter:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return success response
	h.Logger.Info("Handler DeleteLitter OK")
	return c.NoContent(http.StatusOK)
}
