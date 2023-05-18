package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/scuba13/AmacoonServices/internal/login"
)

type LoginHandler struct {
	LoginService *login.LoginService
	Logger       *logrus.Logger
}

func NewLoginHandler(loginService *login.LoginService, logger *logrus.Logger) *LoginHandler {
	return &LoginHandler{
		LoginService: loginService,
		Logger:       logger,
	}
}

func (h *LoginHandler) Login(c echo.Context) error {
	h.Logger.Infof("Handler Login")
	loginRequest := login.LoginRequest{}

	if err := c.Bind(&loginRequest); err != nil {
		h.Logger.WithError(err).Warn("Failed to bind request body")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Chame o serviço de login com os dados do request
	loginResponse, err := h.LoginService.Login(loginRequest)

	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"email": loginRequest.Email,
		}).Warn("User not found or password incorrect")
		return echo.NewHTTPError(http.StatusUnauthorized, "User not found or password incorrect")
	}

	h.Logger.Infof("Handler Login OK")
	return c.JSON(http.StatusOK, loginResponse)
}
