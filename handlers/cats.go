package handlers

import (
	"net/http"

	"amacoonservices/repositories"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CatHandler struct {
	CatRepo repositories.CatRepository
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

func (h *CatHandler) GetCatsByExhibitorAndSex(c echo.Context) error {
	
	
	idExhibitor, err := strconv.Atoi(c.QueryParam("id_exhibitor"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id_exhibitor parameter")
	}

	sex := c.QueryParam("sex")
	if sex == "" {
		return c.String(http.StatusBadRequest, "missing sex parameter")
	}

	sexAsInt, err := getSexAsInt(sex)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	cats, err := h.CatRepo.GetCatsByExhibitorAndSex(idExhibitor, sexAsInt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetCatByRegistration(c echo.Context) error {
  
	registration := c.Param("registration")

    cat, err := h.CatRepo.GetCatByRegistration(registration)
    if err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    if cat == nil {
        return c.String(http.StatusNotFound, "cat not found")
    }

    return c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) GetCatsByExhibitorAndSexService(c echo.Context) error {
	
	
	idExhibitor, err := strconv.Atoi(c.QueryParam("id_exhibitor"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id_exhibitor parameter")
	}

	sex := c.QueryParam("sex")
	if sex == "" {
		return c.String(http.StatusBadRequest, "missing sex parameter")
	}

	sexAsInt, err := getSexAsInt(sex)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	cats, err := h.CatRepo.GetCatsByExhibitorAndSexService(idExhibitor, sexAsInt)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetCatByRegistrationService(c echo.Context) error {
    
	registration := c.Param("registration")

    cat, err := h.CatRepo.GetCatByRegistrationService(registration)
    if err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    if cat == nil {
        return c.String(http.StatusNotFound, "cat not found")
    }

    return c.JSON(http.StatusOK, cat)
}