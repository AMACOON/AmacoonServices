package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/sirupsen/logrus"
)

type CatShowHandler struct {
	CatShowService *catshow.CatShowService
	Logger         *logrus.Logger
}

func NewCatShowHandler(catShowService *catshow.CatShowService, logger *logrus.Logger) *CatShowHandler {
	return &CatShowHandler{
		CatShowService: catShowService,
		Logger:         logger,
	}
}

func (h *CatShowHandler) CreateCatShow(c echo.Context) error {
	h.Logger.Infof("Handler CreateCatShow")

	catShow := new(catshow.CatShow)
	if err := c.Bind(catShow); err != nil {
		h.Logger.Errorf("Failed to bind CatShow: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatShow data")
	}

	createdCatShow, err := h.CatShowService.CreateCatShow(catShow)
	if err != nil {
		h.Logger.Errorf("Failed to create CatShow: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create CatShow")
	}

	h.Logger.Infof("Handler CreateCatShow OK")
	return c.JSON(http.StatusCreated, createdCatShow)
}

func (h *CatShowHandler) GetCatShowByID(c echo.Context) error {
	h.Logger.Infof("Handler GetCatShowByID")

	id := c.Param("id")

	catShow, err := h.CatShowService.GetCatShowByID(id)
	if err != nil {
		h.Logger.Errorf("Failed to get CatShow by ID: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShow")
	}

	h.Logger.Infof("Handler GetCatShowByID OK")
	return c.JSON(http.StatusOK, catShow)
}

func (h *CatShowHandler) UpdateCatShow(c echo.Context) error {
	h.Logger.Infof("Handler UpdateCatShow")

	id := c.Param("id")

	catShow := new(catshow.CatShow)
	if err := c.Bind(catShow); err != nil {
		h.Logger.Errorf("Failed to bind CatShow: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatShow data")
	}

	if err := h.CatShowService.UpdateCatShow(id, catShow); err != nil {
		h.Logger.Errorf("Failed to update CatShow: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update CatShow")
	}

	h.Logger.Infof("Handler UpdateCatShow OK")
	return c.NoContent(http.StatusOK)
}
