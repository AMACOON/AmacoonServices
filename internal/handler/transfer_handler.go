package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"encoding/json"
	
)

type TransferHandler struct {
	TransferService *transfer.TransferService
	Logger          *logrus.Logger
}

func NewTransferHandler(transferService *transfer.TransferService, logger *logrus.Logger) *TransferHandler {
	return &TransferHandler{
		TransferService: transferService,
		Logger:          logger,
	}
}

func (h *TransferHandler) CreateTransfer(c echo.Context) error {
	h.Logger.Infof("Handler CreateTransfer")
	// Get multipart form
	form, err := c.MultipartForm()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get multipart form")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Extract transfer JSON
	transferJson := form.Value["transfer"][0]
	transfer := &transfer.Transfer{}
	err = json.Unmarshal([]byte(transferJson), transfer)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to parse transfer JSON")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate transfer
	if err := utils.ValidateStruct(transfer); err != nil {
		h.Logger.WithError(err).Error("Failed to validate transfer")
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
	createdTransfer, err := h.TransferService.CreateTransfer(*transfer, filesWithDesc)
	if err != nil {
		h.Logger.WithError(err).Error("failed to create transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler CreateTransfer OK")
	return c.JSON(http.StatusCreated, createdTransfer)
}

func (h *TransferHandler) GetTransferByID(c echo.Context) error {
	h.Logger.Infof("Handler GetTransferByID")
	id := c.Param("id")

	foundTransfer, err := h.TransferService.GetTransferByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler GetTransferByID OK")
	return c.JSON(http.StatusOK, foundTransfer)
}

func (h *TransferHandler) UpdateTransfer(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTransfer")
	id := c.Param("id")
	var transferObj transfer.Transfer
	err := c.Bind(&transferObj)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.TransferService.UpdateTransfer(id, transferObj)
	if err != nil {
		h.Logger.WithError(err).Errorf("failed to update transfer with id %s", id)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler UpdateTransfer OK")
	return c.NoContent(http.StatusOK)
}

func (h *TransferHandler) UpdateTransferStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTransferStatus")
	id := c.Param("id")
	status := c.QueryParam("status")
	err := h.TransferService.UpdateTransferStatus(id, status)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update transfer status")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	h.Logger.Infof("Handler UpdateTransferStatus OK")
	return c.NoContent(http.StatusOK)
}

func (h *TransferHandler) GetAllTransfersByRequesterID(c echo.Context) error {
	h.Logger.Infof("Handler GetAllTransfersByRequesterID")
	id := c.Param("requesterID")

	transfers, err := h.TransferService.GetAllTransfersByRequesterID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get transfers by Requester ID")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Handler GetAllTransfersByRequesterID OK")
	return c.JSON(http.StatusOK, transfers)
}




