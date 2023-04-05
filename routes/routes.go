package routes

import (
	servicesHandlers"amacoonservices/handlers/services"
	informationHandlers"amacoonservices/handlers/information"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewRouter(catHandler *informationHandlers.CatHandler,
	ownerHandler *informationHandlers.OwnerHandler,
	colorHandler *informationHandlers.ColorHandler,
	litterHandler *servicesHandlers.LitterHandler,
	breedHandler *informationHandlers.BreedHandler,
	countryHandler *informationHandlers.CountryHandler,
	transferHandler *servicesHandlers.TransferHandler,
	logger *logrus.Logger,
	e *echo.Echo,
) {
	e.Use(middleware.Timeout())
	e.Use(middleware.CORS())
	e.Use(middleware.AddTrailingSlash())

	e.HTTPErrorHandler = customHTTPErrorHandler

	setupCatRoutes(e, catHandler)
	setupOwnerRoutes(e, ownerHandler)
	setupColorRoutes(e, colorHandler)
	setupLitterRoutes(e, litterHandler)
	setupBreedRoutes(e, breedHandler)
	setupCountryRoutes(e, countryHandler)
	setupTransferRoutes(e, transferHandler)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		code = http.StatusInternalServerError
		message = http.StatusText(code)
	}
	if _, ok := message.(string); ok {
		message = map[string]interface{}{"error": message}
	}
	if err := c.JSON(code, message); err != nil {
		c.Logger().Error(err)
	}
}

func setupCatRoutes(e *echo.Echo, catHandler *informationHandlers.CatHandler) {
	e.GET("/cats", catHandler.GetCatsByExhibitorAndSex)
	e.GET("/cats/:registration", catHandler.GetCatByRegistration)
	e.GET("/catsservice", catHandler.GetCatsByExhibitorAndSexService)
	e.GET("/catsservice/:registration", catHandler.GetCatByRegistrationService)
}

func setupOwnerRoutes(e *echo.Echo, ownerHandler *informationHandlers.OwnerHandler) {
	e.GET("/owners/:id", ownerHandler.GetOwnerByID)
}

func setupColorRoutes(e *echo.Echo, colorHandler *informationHandlers.ColorHandler) {
	e.GET("/colors", colorHandler.GetAllColorsByBreed)
}

func setupLitterRoutes(e *echo.Echo, litterHandler *servicesHandlers.LitterHandler) {
	e.GET("/litters", litterHandler.GetAllLitters)
	e.GET("/litters/:id", litterHandler.GetLitterByID)
	e.POST("/litters", litterHandler.CreateLitter)
	e.PUT("/litters/:id", litterHandler.UpdateLitter)
	e.DELETE("/litters/:id", litterHandler.DeleteLitter)
}

func setupBreedRoutes(e *echo.Echo, breedHandler *informationHandlers.BreedHandler) {
	e.GET("/breeds", breedHandler.GetAllBreeds)
}

func setupCountryRoutes(e *echo.Echo, countryHandler *informationHandlers.CountryHandler) {
	e.GET("/countries", countryHandler.GetAllCountry)
}

func setupTransferRoutes(e *echo.Echo, transferHandler *servicesHandlers.TransferHandler) {
    e.POST("/transfer", transferHandler.CreateCatTransferOwnership)
    e.GET("/transfer/:id", transferHandler.GetCatTransferOwnershipByID)
    e.PUT("/transfer/:id", transferHandler.UpdateCatTransferOwnership)
    e.DELETE("/transfer/:id", transferHandler.DeleteCatTransferOwnership)
}

