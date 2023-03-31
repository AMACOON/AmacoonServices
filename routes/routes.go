package routes

import (
	"amacoonservices/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewRouter(catHandler *handlers.CatHandler,
    ownerHandler *handlers.OwnerHandler,
    colorHandler *handlers.ColorHandler,
    litterHandler *handlers.LitterHandler,
    breedHandler *handlers.BreedHandler,
    countryHandler *handlers.CountryHandler,
    logger echo.Logger,
    e *echo.Echo,
)  {

	e.Use(middleware.Timeout())
	e.Use(middleware.CORS())

	// Set response header
	e.Use(middleware.AddTrailingSlash())

	// Set up error handling
	e.HTTPErrorHandler = func(err error, c echo.Context) {
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

	// Cats endpoints
	e.GET("/cats", catHandler.GetCatsByExhibitorAndSex)
	e.GET("/cats/:registration", catHandler.GetCatByRegistration)
	e.GET("/catsservice", catHandler.GetCatsByExhibitorAndSexService)
	e.GET("/catsservice/:registration", catHandler.GetCatByRegistrationService)

	// Owners endpoints
	e.GET("/owners/:id", ownerHandler.GetOwnerByID)

	// Colors endpoints
	e.GET("/colors", colorHandler.GetAllColorsByBreed)

	// Litters endpoints
	e.GET("/litters", litterHandler.GetAllLitters)
	e.GET("/litters/:id", litterHandler.GetLitterByID)
	e.POST("/litters", litterHandler.CreateLitter)
	e.PUT("/litters/:id", litterHandler.UpdateLitter)
	e.DELETE("/litters/:id", litterHandler.DeleteLitter)

	// Breed endpoits
	e.GET("/breeds", breedHandler.GetAllBreeds)

	// Country endpoints
	e.GET("/countries", countryHandler.GetAllCountry)

	
}
