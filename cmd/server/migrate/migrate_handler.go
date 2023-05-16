package migrate

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type MigrateHandler struct {
	MigrateService *MigrateService
	Logger         *logrus.Logger
}

func NewMigrateHandler(service *MigrateService, logger *logrus.Logger) *MigrateHandler {
	return &MigrateHandler{
		MigrateService: service,
		Logger:         logger,
	}
}

func (h *MigrateHandler) MigrateData1(c echo.Context) error {
	h.MigrateService.MigrateData1(h.MigrateService.DB, h.MigrateService.DBOld, h.Logger)

	return c.String(http.StatusOK, "Data 1 migrated successfully")
}

func (h *MigrateHandler) MigrateData2(c echo.Context) error {
	h.MigrateService.MigrateData2(h.MigrateService.DB, h.MigrateService.DBOld, h.Logger)

	return c.String(http.StatusOK, "Data 2 migrated successfully")
}

func (h *MigrateHandler) MigrateData3(c echo.Context) error {
	h.MigrateService.MigrateData3(h.MigrateService.DBOld, h.MigrateService.DB, h.Logger)

	return c.String(http.StatusOK, "Data 3 migrated successfully")
}

func SetupRouter(service *MigrateService, logger *logrus.Logger, e *echo.Echo) {
	handler := NewMigrateHandler(service, logger)

	

	e.GET("/migrate/data1", handler.MigrateData1)
	e.GET("/migrate/data2", handler.MigrateData2)
	e.GET("/migrate/data3", handler.MigrateData3)

	
}
