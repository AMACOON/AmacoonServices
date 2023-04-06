package routes

import (
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/scuba13/AmacoonServices/internal/handler"
)

func NewRouter(catHandler *handler.CatHandler,
	ownerHandler *handler.OwnerHandler,
	colorHandler *handler.ColorHandler,
	litterHandler *handler.LitterHandler,
	breedHandler *handler.BreedHandler,
	countryHandler *handler.CountryHandler,
	transferHandler *handler.TransferHandler,
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

func setupCatRoutes(e *echo.Echo, catHandler *handler.CatHandler) {
	e.GET("/cats", catHandler.GetCatsByExhibitorAndSexTable)
	e.GET("/cats/:registration", catHandler.GetCatByRegistrationTable)
	e.GET("/catsservice", catHandler.GetCatsByExhibitorAndSex)
	e.GET("/catsservice/:registration", catHandler.GetCatByRegistration)
}

func setupOwnerRoutes(e *echo.Echo, ownerHandler *handler.OwnerHandler) {
	e.GET("/owners/:id", ownerHandler.GetOwnerByID)
}

func setupColorRoutes(e *echo.Echo, colorHandler *handler.ColorHandler) {
	e.GET("/colors", colorHandler.GetAllColorsByBreed)
}

func setupLitterRoutes(e *echo.Echo, litterHandler *handler.LitterHandler) {
	e.GET("/litters", litterHandler.GetAllLitters)
	e.GET("/litters/:id", litterHandler.GetLitterByID)
	e.POST("/litters", litterHandler.CreateLitter)
	e.PUT("/litters/:id", litterHandler.UpdateLitter)
	e.DELETE("/litters/:id", litterHandler.DeleteLitter)
}

func setupBreedRoutes(e *echo.Echo, breedHandler *handler.BreedHandler) {
	e.GET("/breeds", breedHandler.GetAllBreeds)
}

func setupCountryRoutes(e *echo.Echo, countryHandler *handler.CountryHandler) {
	e.GET("/countries", countryHandler.GetAllCountry)
}

func setupTransferRoutes(e *echo.Echo, transferHandler *handler.TransferHandler) {
    e.POST("/transfer", transferHandler.CreateTransfer)
    e.GET("/transfer/:id", transferHandler.GetTransferByID)
    e.PUT("/transfer/:id", transferHandler.UpdateTransfer)
    e.DELETE("/transfer/:id", transferHandler.DeleteTransfer)
}





