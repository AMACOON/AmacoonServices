package routes

import (
    "net/http"

    echojwt "github.com/labstack/echo-jwt/v4"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/scuba13/AmacoonServices/internal/handler"
    "github.com/scuba13/AmacoonServices/internal/membership" // 👈 NOVO
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

func NewRouter(
    catHandler *handler.CatHandler,
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
    loginHandler *handler.LoginHandler,
    catShowHandler *handler.CatShowHandler,
    catShowRegistrationHandler *handler.CatShowRegistrationHandler,
    catShowResultHandler *handler.CatShowResultHandler,
    catShowYearHandler *handler.CatShowYearHandler,

    membershipHandler *membership.Handler, // 👈 NOVO

    logger *logrus.Logger,
    e *echo.Echo,
) {

	e.Use(middleware.Timeout())
	e.Use(middleware.CORS())
	e.Use(middleware.AddTrailingSlash())

	e.HTTPErrorHandler = customHTTPErrorHandler

	jwtConfig := getJWTConfig()
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
	setupTitlesRoutes(e, titleHandler, jwtConfig)
	setupTitleRecognitionRoutes(e, titleRecognitionHandler)
	setupCatServiceRoutes(e, catServiceHandler)
	setupFilesRoutes(e, filesHandler)
	setupLoginRoutes(e, loginHandler)
	setupCatShowRoutes(e, catShowHandler)
	setupCatShowRegistrationRoutes(e, catShowRegistrationHandler)
	setupCatShowResultRoutes(e, catShowResultHandler)
	setupCatShowYearRoutes(e, catShowYearHandler)
	setupMembershipRoutes(e, membershipHandler)

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
func getJWTConfig() echojwt.Config {
	secret := viper.GetString("AppJwtSecret")
	return echojwt.Config{
		SigningKey: []byte(secret),
	}
}

func setupHealthChecks(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Beta CatClubSystem")
	})
}

func setupCatRoutes(e *echo.Echo, catHandler *handler.CatHandler) {
	catGroup := e.Group("/api/cats")
	catGroup.GET("/:id", catHandler.GetCatsCompleteByID)
	catGroup.GET("/:ownerId/owner", catHandler.GetCatsByOwner)
	catGroup.POST("", catHandler.CreateCat)
	catGroup.PUT("/:id/neutered", catHandler.UpdateNeuteredStatus)
	catGroup.PUT("/:id", catHandler.UpdateCat)
	catGroup.GET("", catHandler.GetAllCats)
}

func setupCatServiceRoutes(e *echo.Echo, catServiceHandler *handler.CatServiceHandler) {
	catServiceGroup := e.Group("/api/catservice")
	catServiceGroup.GET("/:id", catServiceHandler.GetCatServiceByID)
	catServiceGroup.GET("/:registration/registration", catServiceHandler.GetCatServiceByRegistration)
	catServiceGroup.GET("", catServiceHandler.GetAllCatsServiceByOwnerAndGender)
	catServiceGroup.GET("/:ownerId/owner", catServiceHandler.GetAllCatsServiceByOwner)

}

func setupOwnerRoutes(e *echo.Echo, ownerHandler *handler.OwnerHandler) {
	ownerGroup := e.Group("/api/owners")
	ownerGroup.GET("/:id", ownerHandler.GetOwnerByID)
	ownerGroup.GET("", ownerHandler.GetAllOwners)
	ownerGroup.GET("/:cpf/cpf", ownerHandler.GetOwnerByCPF)
	ownerGroup.POST("", ownerHandler.CreateOwner)
	ownerGroup.PUT("/:id", ownerHandler.UpdateOwner)
	ownerGroup.DELETE("/:id", ownerHandler.DeleteOwnerByID)
	ownerGroup.GET("/:id/:validId/valid", ownerHandler.UpdateValidOwner)

}

func setupColorRoutes(e *echo.Echo, colorHandler *handler.ColorHandler) {
	colorGroup := e.Group("/api/colors")
	colorGroup.GET("/breed/:breedCode", colorHandler.GetAllColorsByBreed)
	colorGroup.GET("/:id", colorHandler.GetColorByID)
	colorGroup.PUT("/:id", colorHandler.UpdateColor)
}

func setupLitterRoutes(e *echo.Echo, litterHandler *handler.LitterHandler) {
	litterGroup := e.Group("/api/litters")
	litterGroup.POST("", litterHandler.CreateLitter)
	litterGroup.GET("/:id", litterHandler.GetLitterByID)
	litterGroup.PUT("/:id", litterHandler.UpdateLitter)
	litterGroup.PUT("/:id/status", litterHandler.UpdateLitterStatus)
	litterGroup.GET("/:requesterID/requesterID", litterHandler.GetAllLittersByRequesterID)
	litterGroup.DELETE("/:id", litterHandler.DeleteLitter)
}

func setupTransferRoutes(e *echo.Echo, transferHandler *handler.TransferHandler) {
	transferGroup := e.Group("/api/transfers")
	transferGroup.POST("", transferHandler.CreateTransfer)
	transferGroup.GET("/:id", transferHandler.GetTransferByID)
	transferGroup.PUT("/:id", transferHandler.UpdateTransfer)
	transferGroup.PUT("/:id/status", transferHandler.UpdateTransferStatus)
	transferGroup.GET("/:requesterID/requesterID", transferHandler.GetAllTransfersByRequesterID)
}

func setupTitleRecognitionRoutes(e *echo.Echo, titleRecognitionHandler *handler.TitleRecognitionHandler) {
	titleRecognitionGroup := e.Group("/api/titles-recognition")
	titleRecognitionGroup.POST("", titleRecognitionHandler.CreateTitleRecognition)
	titleRecognitionGroup.GET("/:id", titleRecognitionHandler.GetTitleRecognitionByID)
	titleRecognitionGroup.PUT("/:id", titleRecognitionHandler.UpdateTitlesRecognition)
	titleRecognitionGroup.PUT("/:id/status", titleRecognitionHandler.UpdateTitleRecognitionStatus)
	titleRecognitionGroup.GET("/:requesterID/requesterID", titleRecognitionHandler.GetAllTitleRecognitionsByRequesterID)
	titleRecognitionGroup.DELETE("/:id", titleRecognitionHandler.DeleteTitleRecognition)
}

func setupBreedRoutes(e *echo.Echo, breedHandler *handler.BreedHandler) {
	breedGroup := e.Group("/api/breeds")
	breedGroup.GET("", breedHandler.GetAllBreeds)
	breedGroup.GET("/:id", breedHandler.GetBreedByID)
}

func setupCatteryRoutes(e *echo.Echo, catteryHandler *handler.CatteryHandler) {
	catteryGroup := e.Group("/api/catteries")
	catteryGroup.GET("", catteryHandler.GetAllCatteries)
	catteryGroup.GET("/:id", catteryHandler.GetCatteryByID)
	catteryGroup.POST("", catteryHandler.CreateCattery)
	catteryGroup.PUT("/:id", catteryHandler.UpdateCattery)
}

func setupFederationRoutes(e *echo.Echo, federationHandler *handler.FederationHandler) {
	federationGroup := e.Group("/api/federations")
	federationGroup.GET("", federationHandler.GetAllFederations)
	federationGroup.GET("/:id", federationHandler.GetFederationByID)
}

func setupCountryRoutes(e *echo.Echo, countryHandler *handler.CountryHandler) {
	countryGroup := e.Group("/api/countries")
	countryGroup.GET("", countryHandler.GetAllCountry)
}

func setupTitlesRoutes(e *echo.Echo, titleHandler *handler.TitleHandler, jwtConfig echojwt.Config) {
	titleGroup := e.Group("/api/titles")
	titleGroup.Use(echojwt.WithConfig(jwtConfig))
	titleGroup.GET("", titleHandler.GetAllTitles)
}

func setupFilesRoutes(e *echo.Echo, filesHandler *handler.FilesHandler) {
	fileGroup := e.Group("/api/files")
	fileGroup.POST("", filesHandler.SaveFiles)
}

func setupLoginRoutes(e *echo.Echo, loginHandler *handler.LoginHandler) {
	loginGroup := e.Group("/api/login")
	loginGroup.POST("/authenticate", loginHandler.Login)

}

func setupCatShowRoutes(e *echo.Echo, catShowHandler *handler.CatShowHandler) {
	catShowGroup := e.Group("/api/catshows")
	catShowGroup.POST("", catShowHandler.CreateCatShow)
	catShowGroup.GET("/:id", catShowHandler.GetCatShowByID)
	catShowGroup.PUT("/:id", catShowHandler.UpdateCatShow)
}

func setupCatShowRegistrationRoutes(e *echo.Echo, catShowRegistrationHandler *handler.CatShowRegistrationHandler) {
	catShowRegistrationGroup := e.Group("/api/catshowregistrations")
	catShowRegistrationGroup.POST("", catShowRegistrationHandler.CreateCatShowRegistration)
	//catShowRegistrationGroup.GET("/:id", catShowRegistrationHandler.GetCatShowRegistrationByID)
	//catShowRegistrationGroup.PUT("/:id", catShowRegistrationHandler.UpdateCatShowRegistration)
}

func setupCatShowResultRoutes(e *echo.Echo, catShowResultHandler *handler.CatShowResultHandler) {
	catShowResultGroup := e.Group("/api/catshowresults")
	catShowResultGroup.POST("", catShowResultHandler.CreateCatShowResult)
	catShowResultGroup.GET("/:id", catShowResultHandler.GetCatShowResultByID)
	catShowResultGroup.GET("/registration/:registrationID", catShowResultHandler.GetCatShowResultByRegistrationID)
	catShowResultGroup.GET("/cat/:catID", catShowResultHandler.GetCatShowResultByCatID)
	catShowResultGroup.PUT("/:id", catShowResultHandler.UpdateCatShowResult)
	catShowResultGroup.DELETE("/:id", catShowResultHandler.DeleteCatShowResult)
}

func setupCatShowYearRoutes(e *echo.Echo, catShowYearHandler *handler.CatShowYearHandler) {
	catShowCompleteGroup := e.Group("/api/catshowyears")
	catShowCompleteGroup.GET("/year/:catID", catShowYearHandler.GetCatShowCompleteByYear)
}

func setupMembershipRoutes(e *echo.Echo, membershipHandler *membership.Handler) {
	group := e.Group("/api/membership-requests")
	group.POST("", membershipHandler.Create)
}


