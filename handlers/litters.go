package handlers

import (
	"amacoonservices/handlers/converter"
	"amacoonservices/models"
	"amacoonservices/repositories"

	"net/http"

	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LitterHandler struct {
	LitterRepo repositories.LitterRepository
}

func (h *LitterHandler) GetAllLitters(c echo.Context) error {
	litters, err := h.LitterRepo.GetAllLitters()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var litterDatas []models.LitterData

	// Transform each Litter and its Kittens into a LitterData struct
	for _, litter := range litters {
		kittens, err := h.LitterRepo.GetKittensByLitterID(litter.ID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		litterData := converter.TransformLitterAndKittensToLitterData(&litter, kittens)
		litterDatas = append(litterDatas, *litterData)
	}

	return c.JSON(http.StatusOK, litterDatas)
}


func (h *LitterHandler) GetLitterByID(c echo.Context) error {
	// Parse the LitterID from the request params
	litterID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid litter ID")
	}

	// Call the repository to get the litter and its kittens
	litter, kittens, err := h.LitterRepo.GetLitterByID(uint(litterID))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Transform the models.Litter and []*models.Kitten into a models.LitterData struct
	litterData := converter.TransformLitterAndKittensToLitterData(litter, kittens)

	// Return the LitterData as a response
	return c.JSON(http.StatusOK, litterData)
}

func (h *LitterHandler) CreateLitter(c echo.Context) error {
	fmt.Println("Handler Litter Create")

	// Parse the request body into a LitterData struct
	var litterData models.LitterData
	if err := c.Bind(&litterData); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Transform LitterData into a models.Litter struct
	litter, kittens := converter.TransformLitterDataToLitterAndKittens(litterData)

	// Call the repository to create the litter and its kittens
	litterID, err := h.LitterRepo.CreateLitter(&litter, kittens)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the LitterID as a response
	fmt.Println("Handler Litter Create - OK")
	return c.JSON(http.StatusOK, map[string]string{
		"litter_id": strconv.Itoa(int(litterID)),
	})

}

func (h *LitterHandler) UpdateLitter(c echo.Context) error {

	return c.JSON(http.StatusOK, "Testeupdate")
}

func (h *LitterHandler) DeleteLitter(c echo.Context) error {

	return c.NoContent(http.StatusOK)
}
