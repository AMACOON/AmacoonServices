package handler

import (
	"net/http"
"strconv"
	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TransferHandler struct {
	TransferService    *transfer.TransferService
	FilesRepo          *utils.FilesRepository
	Logger             *logrus.Logger
	TransferConverter  *transfer.TransferConverter
}

func NewTransferHandler(transferService *transfer.TransferService, filesRepo *utils.FilesRepository, logger *logrus.Logger, transferConverter *transfer.TransferConverter) *TransferHandler {
	return &TransferHandler{
		TransferService:    transferService,
		FilesRepo:          filesRepo,
		Logger:             logger,
		TransferConverter:  transferConverter,
	}
}

func (h *TransferHandler) CreateTransfer(c echo.Context) error {
	h.Logger.Info("Handler CreateCatTransferOwnership")
	var transfer transfer.Transfer
	if err := c.Bind(&transfer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	transferID, protocolNumber, err := h.TransferService.CreateTransfer(transfer)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to create Transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the LitterID as a response
	h.Logger.Info("Handler CreateCatTransferOwnership OK")
	return c.JSON(http.StatusOK, map[string]string{
		"transfer_id": strconv.Itoa(int(transferID)),
		"protocol":    protocolNumber,
	})
}

func (h *TransferHandler) GetAlltransfers(c echo.Context) error {
	h.Logger.Info("Handler GetAlltransfers")
	transfers, err := h.TransferService.GetAllTransfers()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all Transfers")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Info("Handler GetAlltransfers OK")
	return c.JSON(http.StatusOK, transfers)
}


func (h *TransferHandler) GetTransferByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.Logger.WithField("id", id).Info("Handler GetTransferByID")

	transfer, err := h.TransferService.GetTransferByID(uint(id))
	if err != nil {
		h.Logger.WithError(err).Info("Failed to get transfer from service")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	h.Logger.Info("Handler GetTransferByID OK")
	return c.JSON(http.StatusOK, transfer)
}

func (h *TransferHandler) UpdateTransfer(c echo.Context) error {
	h.Logger.Info("Handler UpdateTransfer")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse transfer ID")
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transfer ID")
	}

	var transfer transfer.Transfer
	if err := c.Bind(&transfer); err != nil {
		h.Logger.WithError(err).Error("Failed to parse request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.TransferService.UpdateTransfer(uint(id), &transfer); err != nil {
		h.Logger.WithError(err).Error("Failed to update transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Info("Handler UpdateTransfer OK")
	return c.NoContent(http.StatusOK)
}

func (h *TransferHandler) DeleteTransfer(c echo.Context) error {
	h.Logger.Info("Handler DeleteTransfer")

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse transfer ID")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid transfer ID")
	}

	if err := h.TransferService.DeleteTransfer(uint(id)); err != nil {
		h.Logger.WithError(err).Error("Failed to delete transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Info("Handler DeleteTransfer OK")
	return c.NoContent(http.StatusNoContent)
}






