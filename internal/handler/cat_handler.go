package handler

import (
	"net/http"

	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	//"github.com/golang-jwt/jwt/v4"
	"encoding/json"
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

func (h *CatHandler) CreateCat(c echo.Context) error {
	h.Logger.Infof("Handler CreateCat")

	// Get multipart form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Extract cat JSON
	catJson := form.Value["cat"][0]
	cat := &cat.Cat{}
	err = json.Unmarshal([]byte(catJson), cat)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse cat JSON")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate cat
	if err := utils.ValidateStruct(cat); err != nil {
		h.Logger.WithError(err).Error("Failed to validate cat")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Extract files
	files := form.File["file"]
	descriptions := form.Value["description"]

	var filesWithDesc []utils.FileWithDescription
	for i, file := range files {
		description := ""
		if i < len(descriptions) {
			description = descriptions[i]
		}
		filesWithDesc = append(filesWithDesc, utils.FileWithDescription{
			File:        file,
			Description: description,
		})
	}

	// Create cat
	cat, err = h.CatService.CreateCat(cat, filesWithDesc)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to create cat")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler CreateCat OK")
	return c.JSON(http.StatusCreated, cat)
}

func (h *CatHandler) GetCatsCompleteByID(c echo.Context) error {


	h.Logger.Infof("Handler GetCatsCompleteByID")
	id := c.Param("id")

	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting CatComplete by ID")

	cat, err := h.CatService.GetCatsCompleteByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get CatComplete by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("CatComplete not found by ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetCatsCompleteByID OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) GetCatsByOwner(c echo.Context) error {
	h.Logger.Infof("Handler GetCatCompleteByAllByOwner")
	ownerId := c.Param("ownerId")

	h.Logger.WithFields(logrus.Fields{
		"OwnerId": ownerId,
	}).Info("Getting cat by OwnerID")

	cat, err := h.CatService.GetCatsByOwner(ownerId)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cat by OwnerID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cat == nil {
		h.Logger.WithFields(logrus.Fields{
			"OwnerId": ownerId,
		}).Warn("Cat not found by OwnerID")
		return echo.NewHTTPError(http.StatusNotFound, "cat not found by OwnerID")
	}
	h.Logger.Infof("Handler GetCatCompleteByAllByOwner OK")
	return c.JSON(http.StatusOK, cat)
}

func (h *CatHandler) UpdateNeuteredStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateNeuteredStatus")

	// Get cat ID from path parameter
	catID := c.Param("id")

	// Get neutered status from query parameter
	neutered := c.QueryParam("neutered")

	// Update neutered status
	if err := h.CatService.UpdateNeuteredStatus(catID, neutered); err != nil {
		h.Logger.WithError(err).Error("Failed to update neutered status")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update neutered status")
	}

	h.Logger.Infof("Handler UpdateNeuteredStatus OK")
	return c.NoContent(http.StatusOK)
}

func (h *CatHandler) UpdateCat(c echo.Context) error {
	h.Logger.Infof("Handler UpdateCat")
	
	id := c.Param("id")
	
	var catObj cat.Cat
	err := c.Bind(&catObj)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.CatService.UpdateCat(id, &catObj)
	if err != nil {
		h.Logger.WithError(err).Errorf("failed to update cat with id %s", id)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler UpdateCat OK")
	return c.NoContent(http.StatusOK)
}

func (h *CatHandler) GetAllCats(c echo.Context) error {
	h.Logger.Infof("Handler GetAllCats")
	filter := c.QueryParam("filter")
	cats, err := h.CatService.GetAllCats(filter)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get cats")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler GetAllCats OK")
	return c.JSON(http.StatusOK, cats)
}
