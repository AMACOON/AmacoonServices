package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/user"
)

type UserHandler struct {
	UserService *user.UserService
	Logger       *logrus.Logger
}

func NewUserHandler(userService *user.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Logger:       logger,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	h.Logger.Infof("Handler Login")
	loginRequest:= user.LoginRequest{}

	if err := c.Bind(&loginRequest); err != nil {
		h.Logger.WithError(err).Warn("Failed to bind request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Chame o serviço de login com os dados do request
	user, err := h.UserService.Login(loginRequest)
	
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"email": loginRequest.Email,
		}).Warn("User not found or password incorrect")
		return echo.NewHTTPError(http.StatusUnauthorized, "User not found or password incorrect")
	}

	h.Logger.Infof("Handler Login OK")
	return c.JSON(http.StatusOK, user)
}
