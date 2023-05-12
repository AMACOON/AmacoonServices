package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/scuba13/AmacoonServices/internal/handler"
	"github.com/sirupsen/logrus"
)

func NewRouter(catHandler *handler.CatHandler,
	ownerHandler *handler.OwnerHandler,
	colorHandler *handler.ColorHandler,
	litterHandler *handler.LitterHandler,
	breedHandler *handler.BreedHandler,
	countryHandler *handler.CountryHandler,
	transferHandler *handler.TransferHandler,
	catteryHandler *handler.CatteryHandler,
	federationHandler *handler.FederationHandler,
	titleHandler *handler.TitleHandler,
	titleRecognitionHandler *handler.TitleRecognitionHandler,
	catServiceHandler *handler.CatServiceHandler,
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
	setupCatteryRoutes(e, catteryHandler)
	setupFederationRoutes(e, federationHandler)
	setupTitlesRoutes(e, titleHandler)
	setupTitleRecognitionRoutes(e, titleRecognitionHandler)
	setupCatServiceRoutes(e, catServiceHandler)
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

func setupCatRoutes(e *echo.Echo, catHandler *handler.CatHandler) {
	e.GET("/cats/:id", catHandler.GetCatsCompleteByID)
	// e.GET("/cats/:registration/registration", catHandler.GetCatCompleteByRegistration)
	// e.GET("/cats", catHandler.GetCatsByOwnerAndGender)
	// e.GET("/cats/:ownerId/owner", catHandler.GetAllByOwner)
}

func setupCatServiceRoutes(e *echo.Echo, catServiceHandler *handler.CatServiceHandler) {
	e.GET("/catservice/:id", catServiceHandler.GetCatServiceByID)
	e.GET("/catservice/:registration/registration", catServiceHandler.GetCatServiceByRegistration)
	e.GET("/catservice", catServiceHandler.GetAllCatsServiceByOwnerAndGender)
	e.GET("/catservice/:ownerId/owner", catServiceHandler.GetAllCatsServiceByOwner)
}

func setupOwnerRoutes(e *echo.Echo, ownerHandler *handler.OwnerHandler) {
	e.GET("/owners/:id", ownerHandler.GetOwnerByID)
	e.GET("/owners", ownerHandler.GetAllOwners)
	e.GET("/owners/:cpf/cpf", ownerHandler.GetOwnerByCPF)
	e.POST("/owners", ownerHandler.CreateOwner)
}

func setupColorRoutes(e *echo.Echo, colorHandler *handler.ColorHandler) {
	e.GET("/colors/:breedCode", colorHandler.GetAllColorsByBreed)
}

func setupLitterRoutes(e *echo.Echo, litterHandler *handler.LitterHandler) {
	e.POST("/litters", litterHandler.CreateLitter)
	e.GET("/litters/:id", litterHandler.GetLitterByID)
	e.PUT("/litters/:id", litterHandler.UpdateLitter)
	e.PUT("/litters/:id/status", litterHandler.UpdateLitterStatus)
	e.GET("/litters/:requesterID/requesterID", litterHandler.GetAllLittersByRequesterID)

}

func setupTransferRoutes(e *echo.Echo, transferHandler *handler.TransferHandler) {
	e.POST("/transfers", transferHandler.CreateTransfer)
	e.GET("/transfers/:id", transferHandler.GetTransferByID)
	e.PUT("/transfers/:id", transferHandler.UpdateTransfer)
	e.PUT("/transfers/:id/status", transferHandler.UpdateTransferStatus)
	e.GET("/transfers/:requesterID/requesterID", transferHandler.GetAllTransfersByRequesterID)
}

func setupTitleRecognitionRoutes(e *echo.Echo, titleRecognitionHandler *handler.TitleRecognitionHandler) {
	e.POST("/titles-recognition", titleRecognitionHandler.CreateTitleRecognition)
	e.GET("/titles-recognition/:id", titleRecognitionHandler.GetTitleRecognitionByID)
	e.PUT("/titles-recognition/:id", titleRecognitionHandler.UpdateTitlesRecognition)
	e.PUT("/titles-recognition/:id/status", titleRecognitionHandler.UpdateTitleRecognitionStatus)
	e.GET("/titles-recognition/:requesterID/requesterID", titleRecognitionHandler.GetAllTitleRecognitionsByRequesterID)
}

func setupBreedRoutes(e *echo.Echo, breedHandler *handler.BreedHandler) {
	e.GET("/breeds", breedHandler.GetAllBreeds)
	e.GET("/breeds/:id", breedHandler.GetBreedByID)
}

func setupCatteryRoutes(e *echo.Echo, catteryHandler *handler.CatteryHandler) {
	e.GET("/catteries", catteryHandler.GetAllCatteries)
	e.GET("/catteries/:id", catteryHandler.GetCatteryByID)
}

func setupFederationRoutes(e *echo.Echo, federationHandler *handler.FederationHandler) {
	e.GET("/federations", federationHandler.GetAllFederations)
	e.GET("/federations/:id", federationHandler.GetFederationByID)
}

func setupCountryRoutes(e *echo.Echo, countryHandler *handler.CountryHandler) {
	e.GET("/countries", countryHandler.GetAllCountry)
}

func setupTitlesRoutes(e *echo.Echo, titleHandler *handler.TitleHandler) {
	e.GET("/titles", titleHandler.GetAllTitles)
}
