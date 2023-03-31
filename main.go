package main

import (
	"amacoonservices/config"
	"amacoonservices/handlers"
	"amacoonservices/repositories"
	"amacoonservices/routes"

	"amacoonservices/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	// Initialize Echo and Log
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Logger.Info("Echo And Log Inicialize")

	// Load configuration Data
	cfg := config.LoadConfig()
	e.Logger.Info("Load Configuration DB")

	// Connect to database
	e.Logger.Info("DB Connection Inicialize")
	db, err := config.SetupDB(cfg)
	if err != nil {
		panic(err.Error())
	}

	// Testa a conexão com o banco de dados
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		panic(err.Error())
	}
	e.Logger.Info("DB Connection OK")

	// Migrate database schema
	logger := echo.New().Logger
	logger.Info("Migrate DB Schema")
	if err := db.AutoMigrate(
		&models.Litter{},
		&models.Kitten{},
	); err != nil {
		e.Logger.Fatalf("failed to migrate database schema: %v", err)
	}

	// Initialize repositories
	e.Logger.Info("Initialize Repositories")
	catRepo := repositories.CatRepository{DB: db}
	ownerRepo := repositories.OwnerRepository{DB: db}
	colorRepo := repositories.ColorRepository{DB: db}
	litterRepo := repositories.LitterRepository{DB: db}
	breedRepo := &repositories.BreedRepository{DB: db}
	countryRepo := &repositories.CountryRepository{DB: db}
	e.Logger.Info("Initialize Repositories OK")

	// Initialize handlers
	e.Logger.Info("Initialize Handlers")
	catHandler := &handlers.CatHandler{CatRepo: catRepo}
	ownerHandler := &handlers.OwnerHandler{OwnerRepo: ownerRepo}
	colorHandler := &handlers.ColorHandler{ColorRepo: colorRepo}
	litterHandler := &handlers.LitterHandler{LitterRepo: litterRepo}
	breedHandler := &handlers.BreedHandler{BreedRepo: breedRepo}
	countryHandler := &handlers.CountryHandler{CountryRepo: countryRepo}
	e.Logger.Info("Initialize Handlers OK")

	// Initialize router and routes
	e.Logger.Info("Initialize Router and Routes")

	routes.NewRouter(catHandler, ownerHandler, colorHandler, litterHandler, breedHandler, countryHandler, logger, e)

	e.Logger.Info("Initialize Router and Routes OK")

	// Start server
	e.Logger.Info("Start Server")
	if err := e.Start(":" + "8080"); err != nil {
		e.Logger.Fatalf("failed to start server: %v", err)
	}
}
