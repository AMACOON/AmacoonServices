package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"
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
	var transferReq transfer.TransferRequest
	err := c.Bind(&transferReq)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	createdTransfer, err := h.TransferService.CreateTransfer(transferReq)
	if err != nil {
		h.Logger.WithError(err).Error("failed to create transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create transfer")
	}
	return c.JSON(http.StatusCreated, createdTransfer)
}

func (h *TransferHandler) GetTransferByID(c echo.Context) error {
	h.Logger.Infof("Handler GetTransferByID")
	id := c.Param("id")

	foundTransfer, err := h.TransferService.GetTransferByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get transfer")
	}
	h.Logger.Infof("Handler GetTransferByID OK")
	return c.JSON(http.StatusOK, foundTransfer)
}

func (h *TransferHandler) UpdateTransferStatus(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTransferStatus")
	id := c.Param("id")
	status := c.QueryParam("status")
	err := h.TransferService.UpdateTransferStatus(id, status)
	if err != nil {
		h.Logger.WithError(err).Error("failed to update transfer status")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update transfer status")
	}
	h.Logger.Infof("Handler UpdateTransferStatus OK")
	return c.NoContent(http.StatusOK)
}

func (h *TransferHandler) GetTransferFilesByID(c echo.Context) error {
	h.Logger.Infof("Handler GetTransferFilesByID")
	id := c.Param("id")
	files, err := h.TransferService.GetTransferFilesByID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get transfer files")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get transfer files")
	}
	h.Logger.Infof("Handler GetTransferFilesByID OK")
	return c.JSON(http.StatusOK, files)
}

func (h *TransferHandler) GetAllTransfersByRequesterID(c echo.Context) error {
	h.Logger.Infof("Handler GetAllTransfersByRequesterID")
	id := c.Param("requesterID")

	transfers, err := h.TransferService.GetAllTransfersByRequesterID(id)
	if err != nil {
		h.Logger.WithError(err).Error("failed to get transfers by Requester ID")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get transfers by Requester ID")
	}

	h.Logger.Infof("Handler GetAllTransfersByRequesterID OK")
	return c.JSON(http.StatusOK, transfers)
}

func (h *TransferHandler) UpdateTransfer(c echo.Context) error {
	h.Logger.Infof("Handler UpdateTransfer")
	id := c.Param("id")
	var transferObj transfer.TransferMongo
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

func (h *TransferHandler) AddTransferFiles(c echo.Context) error {
	h.Logger.Infof("Handler AddTransferFiles")
	id := c.Param("id")
	var files []utils.Files
	err := c.Bind(&files)
	if err != nil {
		h.Logger.Errorf("error binding request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	err = h.TransferService.AddTransferFiles(id, files)
	if err != nil {
		h.Logger.WithError(err).Error("failed to add files to transfer")
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to add files to transfer")
	}

	h.Logger.Infof("Handler AddTransferFiles OK")
	return c.NoContent(http.StatusOK)
}
