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
		return echo.NewHTTPError(http.StatusNotFound, "Owner not found")
	}

	h.Logger.Infof("Handler GetOwnerByID OK")
	return c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) GetAllOwners(c echo.Context) error {
	h.Logger.Infof("Handler GetAllOwners")

	owners, err := h.OwnerService.GetAllOwners()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all owners")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get all owners")
	}

	h.Logger.Infof("Handler GetAllOwners OK")
	return c.JSON(http.StatusOK, owners)
}

func (h *OwnerHandler) GetOwnerByCPF(c echo.Context) error {
	h.Logger.Infof("Handler GetOwnerByCPF")

	cpf := c.Param("cpf")
	owner, err := h.OwnerService.GetOwnerByCPF(cpf)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get owner by CPF")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get owner by CPF")
	}

	h.Logger.Infof("Handler GetOwnerByCPF OK")
	return c.JSON(http.StatusOK, owner)
}

func (h *OwnerHandler) CreateOwner(c echo.Context) error {
	h.Logger.Infof("Handler CreateOwner")

	var owner owner.Owner
	if err := c.Bind(&owner); err != nil {
		h.Logger.WithError(err).Error("Failed to parse request body")
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse request body")
	}

	createdOwner, err := h.OwnerService.CreateOwner(&owner)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to create owner")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create owner")
	}

	h.Logger.Infof("Handler CreateOwner OK")
	return c.JSON(http.StatusCreated, createdOwner)
}

func (h *OwnerHandler) UpdateOwner(c echo.Context) error {
	h.Logger.Infof("Handler UpdateOwner")

	id := c.Param("id")
	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Updating Owner")

	owner := new(owner.Owner)
	if err := c.Bind(owner); err != nil {
		h.Logger.WithError(err).Error("Failed to parse request body")
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse request body")
	}

	err := h.OwnerService.UpdateOwner(id, owner)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to update owner")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update owner")
	}

	h.Logger.Infof("Handler UpdateOwner OK")
	return c.NoContent(http.StatusOK)
}

func (h *OwnerHandler) DeleteOwnerByID(c echo.Context) error {
	h.Logger.Infof("Handler DeleteOwnerByID")

	id := c.Param("id")
	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Deleting Owner")

	err := h.OwnerService.DeleteOwnerByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to delete owner")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete owner")
	}

	h.Logger.Infof("Handler DeleteOwnerByID OK")
	return c.NoContent(http.StatusOK)
}

func (h *OwnerHandler) UpdateValidOwner(c echo.Context) error {
	h.Logger.Infof("Handler UpdateValidOwner")

	id := c.Param("id")
	h.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("Updating Valid Owner")

	validID := c.Param("validId")
	
	err := h.OwnerService.UpdateValidOwner(id, validID)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to update owner validation")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update owner validation")
	}

	h.Logger.Infof("Handler UpdateValidOwner OK")
	return c.NoContent(http.StatusOK)
}







func (h *OwnerHandler) Login(c echo.Context) error {
	h.Logger.Infof("Handler Login")
	loginRequest:= owner.LoginRequest{}

	if err := c.Bind(&loginRequest); err != nil {
		h.Logger.WithError(err).Warn("Failed to bind request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Chame o serviço de login com os dados do request
	user, err := h.OwnerService.Login(loginRequest)
	
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"email": loginRequest.Email,
		}).Warn("User not found or password incorrect")
		return echo.NewHTTPError(http.StatusUnauthorized, "User not found or password incorrect")
	}

	h.Logger.Infof("Handler Login OK")
	return c.JSON(http.StatusOK, user)
}
