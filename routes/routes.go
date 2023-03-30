package routes

import (
	"amacoonservices/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewRouter(catHandler *handlers.CatHandler,
			   ownerHandler *handlers.OwnerHandler,
			   colorHandler *handlers.ColorHandler,
			   litterHandler *handlers.LitterHandler,
			   breedHandler *handlers.BreedHandler,
			   countryHandler *handlers.CountryHandler,
			   ) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Timeout())
	e.Use(middleware.CORS())

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

	return e
}
