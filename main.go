package main

import (
	"amacoonservices/config"
	"amacoonservices/handlers"
	"amacoonservices/repositories"
	"amacoonservices/routes"
	"fmt"

	//"amacoonservices/models"

	"log"
)

func main() {
	// Load configuration file
	cfg := config.LoadConfig()

	// Connect to database
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
	fmt.Println("DB Connected")
	// Migrate database schema
	/* if err := db.AutoMigrate(
		&models.Litter{},
		&models.Kitten{},
	); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	} */

	// Initialize repositories
	catRepo := repositories.CatRepository{DB: db}
	ownerRepo := repositories.OwnerRepository{DB: db}
	colorRepo := repositories.ColorRepository{DB: db}
	litterRepo := repositories.LitterRepository{DB: db}
	breedRepo := &repositories.BreedRepository{DB: db}
	countryRepo := &repositories.CountryRepository{DB: db}

	// Initialize handlers
	catHandler := &handlers.CatHandler{CatRepo: catRepo}
	ownerHandler := &handlers.OwnerHandler{OwnerRepo: ownerRepo}
	colorHandler := &handlers.ColorHandler{ColorRepo: colorRepo}
	litterHandler := &handlers.LitterHandler{LitterRepo: litterRepo}
	breedHandler := &handlers.BreedHandler{BreedRepo: breedRepo}
	countryHandler := &handlers.CountryHandler{CountryRepo: countryRepo}

	// Initialize router and routes
	echo := routes.NewRouter(catHandler, ownerHandler, colorHandler, litterHandler, breedHandler, countryHandler)

	// Start server
	if err := echo.Start(":" + "8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
