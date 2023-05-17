package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	
	"github.com/scuba13/AmacoonServices/internal/judge"
)

type JudgeHandler struct {
	JudgeService *judge.JudgeService
	Logger       *logrus.Logger
}

func NewJudgeHandler(judgeService *judge.JudgeService, logger *logrus.Logger) *JudgeHandler {
	return &JudgeHandler{
		JudgeService: judgeService,
		Logger:       logger,
	}
}

func (h *JudgeHandler) GetAllJudges(c echo.Context) error {
	h.Logger.Infof("Handler GetAllJudges")
	judges, err := h.JudgeService.GetAllJudges()
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get all judges")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Retrieved %d judges", len(judges))
	h.Logger.Infof("Handler GetAllJudges OK")
	return c.JSON(http.StatusOK, judges)
}

func (h *JudgeHandler) GetJudgeByID(c echo.Context) error {
	h.Logger.Infof("Handler GetJudgeByID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	judge, err := h.JudgeService.GetJudgeByID(uint(id))
	if err != nil {
		h.Logger.WithError(err).Error("Failed to get judge")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Retrieved judge ID %d", id)
	h.Logger.Infof("Handler GetJudgeByID OK")
	return c.JSON(http.StatusOK, judge)
}

func (h *JudgeHandler) UpdateJudge(c echo.Context) error {
	h.Logger.Infof("Handler UpdateJudge")

	// Parse ID from path
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	// Initialize an empty Judge object
	var judge judge.Judge

	// Bind the request body to the Judge object
	if err := c.Bind(&judge); err != nil {
		h.Logger.WithError(err).Error("Failed to parse request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Call the UpdateJudge service
	updatedJudge, err := h.JudgeService.UpdateJudge(uint(id), &judge)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to update judge")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Updated judge ID %d", id)
	h.Logger.Infof("Handler UpdateJudge OK")
	return c.JSON(http.StatusOK, updatedJudge)
}



func (h *JudgeHandler) DeleteJudge(c echo.Context) error {
	h.Logger.Infof("Handler DeleteJudge")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	if err := h.JudgeService.DeleteJudge(uint(id)); err != nil {
		h.Logger.WithError(err).Error("Failed to delete judge")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	h.Logger.Infof("Deleted judge ID %d", id)
	h.Logger.Infof("Handler DeleteJudge OK")
	return c.NoContent(http.StatusOK)
}
