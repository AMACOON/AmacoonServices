package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/catshowcomplete"
	"github.com/sirupsen/logrus"
)

type CatShowCompleteHandler struct {
	CatShowCompleteService *catshowcomplete.CatShowCompleteService
	Logger                 *logrus.Logger
}

func NewCatShowCompleteHandler(catShowCompleteService *catshowcomplete.CatShowCompleteService, logger *logrus.Logger) *CatShowCompleteHandler {
	return &CatShowCompleteHandler{
		CatShowCompleteService: catShowCompleteService,
		Logger:                 logger,
	}
}

func (h *CatShowCompleteHandler) GetCatShowCompleteByID(c echo.Context) error {
	h.Logger.Info("Handler GetCatShowCompleteByRegistrationID")

	registrationID, err := strconv.ParseUint(c.Param("registrationID"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid RegistrationID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid RegistrationID format")
	}

	catShowComplete, err := h.CatShowCompleteService.GetCatShowCompleteByID(uint(registrationID))
	if err != nil {
		h.Logger.Errorf("Failed to get CatShowComplete by RegistrationID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowComplete by RegistrationID")
	}

	h.Logger.Info("Handler GetCatShowCompleteByRegistrationID OK")
	return c.JSON(http.StatusOK, catShowComplete)
}

func (h *CatShowCompleteHandler) GetCatShowCompleteByCatID(c echo.Context) error {
	h.Logger.Info("Handling GetCatShowCompleteByCatID request")

	// Extrai o CatID do caminho da URL como um parâmetro.
	catID, err := strconv.ParseUint(c.Param("catID"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid CatID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatID format")
	}

	// Chama o serviço CatShowCompleteService para buscar as informações completas baseadas no CatID.
	catShowCompletes, err := h.CatShowCompleteService.GetCatShowCompleteByCatID(uint(catID))
	if err != nil {
		h.Logger.Errorf("Failed to get CatShowComplete by CatID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowComplete by CatID")
	}

	// Retorna as informações completas encontradas com o status HTTP 200 OK.
	h.Logger.Info("GetCatShowCompleteByCatID request handled successfully")
	return c.JSON(http.StatusOK, catShowCompletes)
}

func (h *CatShowCompleteHandler) GetCatShowCompleteByCatShowIDs(c echo.Context) error {
	h.Logger.Info("Handler GetCatShowCompleteByCatShowIDs")

	catShowID, err := strconv.ParseUint(c.Param("catShowID"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid CatShowID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatShowID format")
	}

	// catShowSubID é opcional
	catShowSubIDParam := c.Param("catShowSubID")
	var catShowSubID *uint
	if catShowSubIDParam != "" {
		subID, err := strconv.ParseUint(catShowSubIDParam, 10, 64)
		if err != nil {
			h.Logger.Errorf("Invalid CatShowSubID format: %v", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatShowSubID format")
		}
		subIDUint := uint(subID)
		catShowSubID = &subIDUint
	}

	catShowCompletes, err := h.CatShowCompleteService.GetCatShowCompleteByCatShowIDs(uint(catShowID), catShowSubID)
	if err != nil {
		h.Logger.Errorf("Failed to get CatShowComplete by CatShowIDs: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowComplete by CatShowIDs")
	}

	h.Logger.Info("Handler GetCatShowCompleteByCatShowIDs OK")
	return c.JSON(http.StatusOK, catShowCompletes)
}

