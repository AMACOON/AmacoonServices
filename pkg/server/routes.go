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
	filesHandler *handler.FilesHandler,
	logger *logrus.Logger,
	e *echo.Echo,
) {
	e.Use(middleware.Timeout())
	e.Use(middleware.CORS())
	e.Use(middleware.AddTrailingSlash())

	e.HTTPErrorHandler = customHTTPErrorHandler

	setupHealthChecks(e)
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
	setupFilesRoutes(e, filesHandler)
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

func setupHealthChecks(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func setupCatRoutes(e *echo.Echo, catHandler *handler.CatHandler) {
	catGroup := e.Group("/cats")
	catGroup.GET("/:id", catHandler.GetCatsCompleteByID)
	catGroup.GET("/:ownerId/owner", catHandler.GetCatCompleteByAllByOwner)
}

func setupCatServiceRoutes(e *echo.Echo, catServiceHandler *handler.CatServiceHandler) {
	catServiceGroup := e.Group("/catservice")
	catServiceGroup.GET("/:id", catServiceHandler.GetCatServiceByID)
	catServiceGroup.GET("/:registration/registration", catServiceHandler.GetCatServiceByRegistration)
	catServiceGroup.GET("", catServiceHandler.GetAllCatsServiceByOwnerAndGender)
	catServiceGroup.GET("/:ownerId/owner", catServiceHandler.GetAllCatsServiceByOwner)

}

func setupOwnerRoutes(e *echo.Echo, ownerHandler *handler.OwnerHandler) {
	ownerGroup := e.Group("/owners")
	ownerGroup.GET("/:id", ownerHandler.GetOwnerByID)
	ownerGroup.GET("", ownerHandler.GetAllOwners)
	ownerGroup.GET("/:cpf/cpf", ownerHandler.GetOwnerByCPF)
	ownerGroup.POST("", ownerHandler.CreateOwner)
	ownerGroup.PUT("/:id", ownerHandler.UpdateOwner)
	ownerGroup.DELETE("/:id", ownerHandler.DeleteOwnerByID)
	ownerGroup.GET("/:id/:validId/valid", ownerHandler.UpdateValidOwner)
	ownerGroup.POST("/login", ownerHandler.Login)
	
}

func setupColorRoutes(e *echo.Echo, colorHandler *handler.ColorHandler) {
	colorGroup := e.Group("/colors")
	colorGroup.GET("/:breedCode", colorHandler.GetAllColorsByBreed)
}

func setupLitterRoutes(e *echo.Echo, litterHandler *handler.LitterHandler) {
	litterGroup := e.Group("/litters")
	litterGroup.POST("", litterHandler.CreateLitter)
	litterGroup.GET("/:id", litterHandler.GetLitterByID)
	litterGroup.PUT("/:id", litterHandler.UpdateLitter)
	litterGroup.PUT("/:id/status", litterHandler.UpdateLitterStatus)
	litterGroup.GET("/:requesterID/requesterID", litterHandler.GetAllLittersByRequesterID)
	litterGroup.DELETE("/:id", litterHandler.DeleteLitter)
}

func setupTransferRoutes(e *echo.Echo, transferHandler *handler.TransferHandler) {
	transferGroup := e.Group("/transfers")
	transferGroup.POST("", transferHandler.CreateTransfer)
	transferGroup.GET("/:id", transferHandler.GetTransferByID)
	transferGroup.PUT("/:id", transferHandler.UpdateTransfer)
	transferGroup.PUT("/:id/status", transferHandler.UpdateTransferStatus)
	transferGroup.GET("/:requesterID/requesterID", transferHandler.GetAllTransfersByRequesterID)
}

func setupTitleRecognitionRoutes(e *echo.Echo, titleRecognitionHandler *handler.TitleRecognitionHandler) {
	titleRecognitionGroup := e.Group("/titles-recognition")
	titleRecognitionGroup.POST("", titleRecognitionHandler.CreateTitleRecognition)
	titleRecognitionGroup.GET("/:id", titleRecognitionHandler.GetTitleRecognitionByID)
	titleRecognitionGroup.PUT("/:id", titleRecognitionHandler.UpdateTitlesRecognition)
	titleRecognitionGroup.PUT("/:id/status", titleRecognitionHandler.UpdateTitleRecognitionStatus)
	titleRecognitionGroup.GET("/:requesterID/requesterID", titleRecognitionHandler.GetAllTitleRecognitionsByRequesterID)
	titleRecognitionGroup.DELETE("/:id", titleRecognitionHandler.DeleteTitleRecognition)
}

func setupBreedRoutes(e *echo.Echo, breedHandler *handler.BreedHandler) {
	breedGroup := e.Group("/breeds")
	breedGroup.GET("", breedHandler.GetAllBreeds)
	breedGroup.GET("/:id", breedHandler.GetBreedByID)
}

func setupCatteryRoutes(e *echo.Echo, catteryHandler *handler.CatteryHandler) {
	catteryGroup := e.Group("/catteries")
	catteryGroup.GET("", catteryHandler.GetAllCatteries)
	catteryGroup.GET("/:id", catteryHandler.GetCatteryByID)
}

func setupFederationRoutes(e *echo.Echo, federationHandler *handler.FederationHandler) {
	federationGroup := e.Group("/federations")
	federationGroup.GET("", federationHandler.GetAllFederations)
	federationGroup.GET("/:id", federationHandler.GetFederationByID)
}

func setupCountryRoutes(e *echo.Echo, countryHandler *handler.CountryHandler) {
	countryGroup := e.Group("/countries")
	countryGroup.GET("", countryHandler.GetAllCountry)
}

func setupTitlesRoutes(e *echo.Echo, titleHandler *handler.TitleHandler) {
	titleGroup := e.Group("/titles")
	titleGroup.GET("", titleHandler.GetAllTitles)
}

func setupFilesRoutes(e *echo.Echo, filesHandler *handler.FilesHandler) {
	fileGroup := e.Group("/files")
	fileGroup.POST("/:protocolNumber", filesHandler.SaveFiles)

}
