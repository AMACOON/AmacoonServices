package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"
	"github.com/sirupsen/logrus"
	"strconv"
)

type CatShowResultHandler struct {
	CatShowResultService *catshowresult.CatShowResultService
	Logger               *logrus.Logger
}

func NewCatShowResultHandler(catShowResultService *catshowresult.CatShowResultService, logger *logrus.Logger) *CatShowResultHandler {
	return &CatShowResultHandler{
		CatShowResultService: catShowResultService,
		Logger:               logger,
	}
}

func (h *CatShowResultHandler) CreateCatShowResult(c echo.Context) error {
	h.Logger.Info("Handler CreateCatShowResult")

	catShowResult := new(catshowresult.CatShowResult)
	if err := c.Bind(catShowResult); err != nil {
		h.Logger.Errorf("Failed to bind CatShowResult: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatShowResult data")
	}

	createdCatShowResult, err := h.CatShowResultService.CreateCatShowResult(catShowResult)
	if err != nil {
		h.Logger.Errorf("Failed to create CatShowResult: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create CatShowResult")
	}

	h.Logger.Info("Handler CreateCatShowResult OK")
	return c.JSON(http.StatusCreated, createdCatShowResult)
}

func (h *CatShowResultHandler) GetCatShowResultByID(c echo.Context) error {
	h.Logger.Info("Handler GetCatShowResultByID")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid ID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	catShowResult, err := h.CatShowResultService.GetCatShowResultByID(uint(id))
	if err != nil {
		h.Logger.Errorf("Failed to get CatShowResult by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowResult")
	}

	h.Logger.Info("Handler GetCatShowResultByID OK")
	return c.JSON(http.StatusOK, catShowResult)
}

func (h *CatShowResultHandler) GetCatShowResultByRegistrationID(c echo.Context) error {
	h.Logger.Info("Handler GetCatShowResultByRegistrationID")

	registrationID, err := strconv.ParseUint(c.Param("registrationID"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid RegistrationID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid RegistrationID format")
	}

	catShowResult, err := h.CatShowResultService.GetCatShowResultByRegistrationID(uint(registrationID))
	if err != nil {
		h.Logger.Errorf("Failed to get CatShowResult by RegistrationID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowResult by RegistrationID")
	}

	h.Logger.Info("Handler GetCatShowResultByRegistrationID OK")
	return c.JSON(http.StatusOK, catShowResult)
}

func (h *CatShowResultHandler) GetCatShowResultByCatID(c echo.Context) error {
	h.Logger.Info("Handler GetCatShowResultByCatID")

	catID, err := strconv.ParseUint(c.Param("catID"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid CatID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatID format")
	}

	catShowResult, err := h.CatShowResultService.GetCatShowResultByCatID(uint(catID))
	if err != nil {
		h.Logger.Errorf("Failed to get CatShowResult by CatID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowResult by CatID")
	}

	h.Logger.Info("Handler GetCatShowResultByCatID OK")
	return c.JSON(http.StatusOK, catShowResult)
}

func (h *CatShowResultHandler) UpdateCatShowResult(c echo.Context) error {
	h.Logger.Info("Handler UpdateCatShowResult")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid ID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	catShowResult := new(catshowresult.CatShowResult)
	if err := c.Bind(catShowResult); err != nil {
		h.Logger.Errorf("Failed to bind CatShowResult: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatShowResult data")
	}

	if err := h.CatShowResultService.UpdateCatShowResult(uint(id), catShowResult); err != nil {
		h.Logger.Errorf("Failed to update CatShowResult: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update CatShowResult")
	}

	h.Logger.Info("Handler UpdateCatShowResult OK")
	return c.NoContent(http.StatusOK)
}

func (h *CatShowResultHandler) DeleteCatShowResult(c echo.Context) error {
	h.Logger.Info("Handler DeleteCatShowResult")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		h.Logger.Errorf("Invalid ID format: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	if err := h.CatShowResultService.DeleteCatShowResult(uint(id)); err != nil {
		h.Logger.Errorf("Failed to delete CatShowResult: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete CatShowResult")
	}

	h.Logger.Info("Handler DeleteCatShowResult OK")
	return c.NoContent(http.StatusNoContent)
}

