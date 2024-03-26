package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/catshowyear"
	"github.com/sirupsen/logrus"
)

type CatShowYearHandler struct {
	CatShowYearService *catshowyear.CatShowYearService
	Logger                 *logrus.Logger
}

func NewCatShowYearHandler(catShowYearService *catshowyear.CatShowYearService, logger *logrus.Logger) *CatShowYearHandler {
	return &CatShowYearHandler{
		CatShowYearService: catShowYearService,
		Logger:                 logger,
	}
}

func (h *CatShowYearHandler) GetCatShowCompleteByYear(c echo.Context) error {
    h.Logger.Info("Handling GetCatShowCompleteByYear request")

    // Extrai o CatID do caminho da URL como um parâmetro.
    catID, err := strconv.ParseUint(c.Param("catID"), 10, 64)
    if err != nil {
        h.Logger.Errorf("Invalid CatID format: %v", err)
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid CatID format")
    }

    // Chama o serviço CatShowCompleteService para buscar as informações completas baseadas no CatID, agrupadas por ano.
    catShowYearGroups, err := h.CatShowYearService.GetCatShowCompleteByYear(uint(catID))
    if err != nil {
        h.Logger.Errorf("Failed to get CatShowComplete by year: %v", err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CatShowComplete by year")
    }

    // Retorna as informações completas encontradas, agrupadas por ano, com o status HTTP 200 OK.
    h.Logger.Info("GetCatShowCompleteByYear request handled successfully")
    return c.JSON(http.StatusOK, catShowYearGroups)
}
