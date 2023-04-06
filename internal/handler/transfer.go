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
	TransferRepo *transfer.TransferRepository
	FilesRepo    *utils.FilesRepository
	Logger       *logrus.Logger
	TransferConverter *transfer.TransferConverter
}

func NewTransferHandler(transferRepo *transfer.TransferRepository, filesRepo *utils.FilesRepository, logger *logrus.Logger, transferConverter *transfer.TransferConverter) *TransferHandler {
	return &TransferHandler{
		TransferRepo: transferRepo,
		FilesRepo:    filesRepo,
		Logger:       logger,
		TransferConverter: transferConverter,
	}
}

func (h *TransferHandler) CreateTransfer(c echo.Context) error {
	h.Logger.Info("Handler CreateCatTransferOwnership")
	var transfer transfer.Transfer
	if err := c.Bind(&transfer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	transferDB, filesDB := h.TransferConverter.TransferToTransferDB(transfer)
	transferID, protocolNumber, err := h.TransferRepo.CreateTransfer(&transferDB, filesDB)
	if err != nil {
		h.Logger.Error("Failed to create Transfer:", err)
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
	transfers, err := h.TransferRepo.GetAlltransfers()
	if err != nil {
		h.Logger.Errorf("Failed to get all Transfers: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var transferDatas []transfer.Transfer

	// Transform each TransferDB and FilesDB into a Transfer struct
	for _, transfer := range transfers {
		files, err := h.FilesRepo.GetFilesByServiceID(transfer.ID)
		if err != nil {
			h.Logger.Errorf("Failed to get files by litter ID %v: %v", transfer.ID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		transferData := h.TransferConverter.TransferDBToTransfer(&transfer, files)
		transferDatas = append(transferDatas, *transferData)
	}
	h.Logger.Info("Handler GetAlltransfers OK")
	return c.JSON(http.StatusOK, transferDatas)
}

func (h *TransferHandler) GetTransferByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transferDB, filesDB, err := h.TransferRepo.GetTransferByID(uint(id))
	if err != nil {
		c.Logger().Info("Erro Handler:", err)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	transfer := h.TransferConverter.TransferDBToTransfer(transferDB, filesDB)
	return c.JSON(http.StatusOK, transfer)
}

func (h *TransferHandler) UpdateTransfer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var transfer transfer.Transfer
	if err := c.Bind(&transfer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	transferDB, filesDB := h.TransferConverter.TransferToTransferDB(transfer)
	if err := h.TransferRepo.UpdateTransfer(uint(id), &transferDB, filesDB); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Lida com a solicitação para excluir uma transferência de gato
func (h *TransferHandler) DeleteTransfer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("Failed to parse Transfer ID:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Transfer ID")
	}

	// Chama o repositório para excluir a transferência de gato
	err = h.TransferRepo.DeleteTransfer(uint(id))
	if err != nil {
		h.Logger.Error("Failed to delete Transfer:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
