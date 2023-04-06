package handler
import (
	"github.com/scuba13/AmacoonServices/internal/cat"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CatHandler struct {
	CatService *cat.CatService
	Logger     *logrus.Logger
}

func NewCatHandler(catService *cat.CatService, logger *logrus.Logger) *CatHandler {
	return &CatHandler{
		CatService: catService,
		Logger:     logger,
	}
}

func getSexAsInt(sex string) (int, error) {
	switch sex {
	case "M":
		return 1, nil
	case "F":
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid sex parameter")
	}
}

func (h *CatHandler) GetCatsByExhibitorAndSexTable(c echo.Context) error {
	
	h.Logger.Infof("Handler GetCatsByExhibitorAndSex")
	exhibitorID, err := strconv.Atoi(c.QueryParam("id_exhibitor"))
	if err != nil {
		h.Logger.Warn("Invalid id_exhibitor parameter")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id_exhibitor parameter")
	}

	sex := c.QueryParam("sex")
	if sex == "" {
		h.Logger.Warn("Missing sex parameter")
		return echo.NewHTTPError(http.StatusBadRequest, "missing sex parameter")
	}

	sexAsInt, err := getSexAsInt(sex)
	if err != nil {
		h.Logger.Warn("Invalid sex parameter")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	h.Logger.WithFields(logrus.Fields{
		"exhibitorID": exhibitorID,
		"sex":         sexAsInt,
	}).Info("Getting cats by exhibitor and sex")

	cats, err := h.CatService.GetCatsByExhibitorAndSexTable(exhibitorID, sexAsInt)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cats by exhibitor and sex")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler GetCatsByExhibitorAndSex OK")
	return c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetCatByRegistrationTable(c echo.Context) error {
	h.Logger.Infof("Handler GetCatByRegistration")
	registration := c.Param("registration")

	h.Logger.WithFields(logrus.Fields{
		"registration": registration,
	}).Info("Getting cat by registration")

	cat, err := h.CatService.GetCatByRegistrationTable(registration)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by registration")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"registration": registration,
		}).Warn("Cat not found by registration")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found")
	}
	h.Logger.Infof("Handler GetCatByRegistration OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) GetCatsByExhibitorAndSex(c echo.Context) error {
	
	h.Logger.Infof("Handler GetCatsByExhibitorAndSex")
	exhibitorID, err := strconv.Atoi(c.QueryParam("id_exhibitor"))
	if err != nil {
		h.Logger.Warn("Invalid id_exhibitor parameter")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id_exhibitor parameter")
	}

	sex := c.QueryParam("sex")
	if sex == "" {
		h.Logger.Warn("Missing sex parameter")
		return echo.NewHTTPError(http.StatusBadRequest, "missing sex parameter")
	}

	sexAsInt, err := getSexAsInt(sex)
	if err != nil {
		h.Logger.Warn("Invalid sex parameter")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	h.Logger.WithFields(logrus.Fields{
		"exhibitorID": exhibitorID,
		"sex":         sexAsInt,
	}).Info("Getting cats by exhibitor and sex")

	cats, err := h.CatService.GetCatsByExhibitorAndSex(exhibitorID, sexAsInt)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cats by exhibitor and sex")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler GetCatsByExhibitorAndSex OK")
	return c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetCatByRegistration(c echo.Context) error {
	h.Logger.Infof("Handler GetCatByRegistration")
	registration := c.Param("registration")

	h.Logger.WithFields(logrus.Fields{
		"registration": registration,
	}).Info("Getting cat by registration")

	cat, err := h.CatService.GetCatByRegistration(registration)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by registration")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"registration": registration,
		}).Warn("Cat not found by registration")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found")
	}
	h.Logger.Infof("Handler GetCatByRegistration OK")
	return c.JSON(http.StatusOK, cat)
}


