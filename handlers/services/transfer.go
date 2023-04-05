package handlers

import (
	"net/http"
	"strconv"

	"amacoonservices/handlers/services/converter"
	"amacoonservices/models/services"
	"amacoonservices/repositories/services"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TransferHandler struct {
	TransferRepo *repositories.TransferRepository
	Logger     *logrus.Logger
}

func NewTransferHandler(transferRepo *repositories.TransferRepository, logger *logrus.Logger) *TransferHandler {
	return &TransferHandler{
		TransferRepo: transferRepo,
		Logger:     logger,
	}
}

func (h *TransferHandler) CreateCatTransferOwnership(c echo.Context) error {
	h.Logger.Info("Handler CreateCatTransferOwnership")
	var catTransferOwnership models.Transfer
	if err := c.Bind(&catTransferOwnership); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	catTransferOwnershipDB:= converter.CatTransferOwnershipToCatTransferOwnerShipDB(catTransferOwnership)
	transferID, protocolNumber, err := h.TransferRepo.CreateCatTransferOwnership(&catTransferOwnershipDB);
	if err != nil {
		h.Logger.Error("Failed to create Transfer:", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the LitterID as a response
	h.Logger.Info("Handler CreateCatTransferOwnership OK")
	return c.JSON(http.StatusOK, map[string]string{
		"transfer_id": strconv.Itoa(int(transferID)),
		"protocol":  protocolNumber,
	})
}

func (h *TransferHandler) GetCatTransferOwnershipByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	catTransferOwnership, err := h.TransferRepo.GetCatTransferOwnershipByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, catTransferOwnership)
}

func (h *TransferHandler) UpdateCatTransferOwnership(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var catTransferOwnership models.Transfer
	if err := c.Bind(&catTransferOwnership); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	catTransferOwnershipDB:= converter.CatTransferOwnershipToCatTransferOwnerShipDB(catTransferOwnership)
	if err := h.TransferRepo.UpdateCatTransferOwnership(uint(id), &catTransferOwnershipDB); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

// Lida com a solicitação para excluir uma transferência de gato
func (h *TransferHandler) DeleteCatTransferOwnership(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        h.Logger.Error("Failed to parse CatTransferOwnership ID:", err)
        return c.String(http.StatusBadRequest, "Invalid CatTransferOwnership ID")
    }

    // Chama o repositório para excluir a transferência de gato
    err = h.TransferRepo.DeleteCatTransferOwnership(uint(id))
    if err != nil {
        h.Logger.Error("Failed to delete CatTransferOwnership:", err)
        return c.String(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusNoContent)
}
