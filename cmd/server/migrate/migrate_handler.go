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

func (h *MigrateHandler) MigrateData(c echo.Context) error {
	
	go h.MigrateService.MigrateData(h.MigrateService.DB, h.MigrateService.DBOld, h.Logger)

	return c.String(http.StatusOK, "Data migration started")
}


func SetupRouter(service *MigrateService, logger *logrus.Logger, e *echo.Echo) {
	
	handler := NewMigrateHandler(service, logger)
	
	e.GET("/migrate/data1", handler.MigrateData)
	
}
