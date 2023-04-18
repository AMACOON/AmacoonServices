package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/owner"
)

type OwnerHandler struct {
	OwnerService *owner.OwnerService
	Logger       *logrus.Logger
}

func NewOwnerHandler(ownerService *owner.OwnerService, logger *logrus.Logger) *OwnerHandler {
	return &OwnerHandler{
		OwnerService: ownerService,
		Logger:       logger,
	}
}

func (h *OwnerHandler) GetOwnerByID(c echo.Context) error {
	h.Logger.Infof("Handler GetOwnerByID")

	id := c.Param("id")
	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Getting Owner by ID")

	owner, err := h.OwnerService.GetOwnerByID(id)
	if err != nil {
		h.Logger.WithError(err).Warn("Failed to get owner from service")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if owner == nil {
		h.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("Owner not found by ID")
		return echo.NewHTTPError(http.StatusNotFound, "cattery not found")
	}

	h.Logger.Infof("Handler GetOwnerByID OK")
	return c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) GetAllOwners(c echo.Context) error {

	// Log de entrada da função
	h.Logger.Infof("Handler GetAllOwners")

	owners, err := h.OwnerService.GetAllOwners()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all owners")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get all owners")
	}

	// Log de saída da função
	h.Logger.Infof("Handler GetAllOwners OK")
	return c.JSON(http.StatusOK, owners)
}

func (h *OwnerHandler) GetOwnerByCPF(c echo.Context) error {
    // Log de entrada da função
    h.Logger.Infof("Handler GetOwnerByCPF")

    cpf := c.Param("cpf")
    owner, err := h.OwnerService.GetOwnerByCPF(cpf)
    if err != nil {
        h.Logger.WithError(err).Error("Failed to get owner by CPF")
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get owner by CPF")
    }

    // Log de saída da função
    h.Logger.Infof("Handler GetOwnerByCPF OK")
    return c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) CreateOwner(c echo.Context) error {
    // Log de entrada da função
    h.Logger.Infof("Handler CreateOwner")

    var owner owner.OwnerMongo
    if err := c.Bind(&owner); err != nil {
        h.Logger.WithError(err).Error("Failed to parse request body")
        return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse request body")
    }

    createdOwner, err := h.OwnerService.CreateOwner(&owner)
    if err != nil {
        h.Logger.WithError(err).Error("Failed to create owner")
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create owner")
    }

    // Log de saída da função
    h.Logger.Infof("Handler CreateOwner OK")
    return c.JSON(http.StatusCreated, createdOwner)
}





