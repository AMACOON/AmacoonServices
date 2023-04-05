package main

import (
	"amacoonservices/config"
	informationHandlers "amacoonservices/handlers/information"
	servicesHandlers "amacoonservices/handlers/services"

	informationRepo "amacoonservices/repositories/information"
	servicesRepo "amacoonservices/repositories/services"

	"amacoonservices/routes"

	servicesModel "amacoonservices/models/services"
	utilsModel "amacoonservices/models/utils"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo and Logger
	e := echo.New()
	logger := setupLogger()
	logger.Info("Initialize Echo and Logger")
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${remote_ip} ${method} ${uri} ${status} ${error}\n",
		Output: logrus.StandardLogger().Out,
	}))

	// Load configuration data
	logger.Info("Load configuration data")
	cfg := config.LoadConfig()

	// Connect to database
	db := setupDatabase(cfg, logger)

	// Initialize repositories, handlers, and routes
	initializeApp(e, db, logger)

	// Start server
	logger.Info("Starting Server")
	if err := e.Start(":" + "8080"); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.Info("Logger Initialized")
	return logger
}

/*
	 func setupLogger() *logrus.Logger {
	    logger := logrus.New()
	    logger.SetLevel(logrus.InfoLevel)
	    logger.SetFormatter(&logrus.TextFormatter{
	        FullTimestamp: true,
	    })
	    logger.SetOutput(os.Stdout)

	    logFile, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	    if err == nil {
	        logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
	    } else {
	        logger.Info("Failed to log to file, using default stderr")
	    }

	    return logger
	}
*/
func setupDatabase(cfg *config.Config, logger *logrus.Logger) *gorm.DB {
	logger.Info("Connecting DB")
	db, err := config.SetupDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to initialize DB connection: %v", err)
	}

	if err := db.AutoMigrate(
		&servicesModel.LitterDB{},
		&servicesModel.KittenDB{},
		&servicesModel.TransferDB{},
		&utilsModel.ProtocolDB{},
	); err != nil {
		logger.Fatalf("Failed to migrate database schema: %v", err)
	}
	logger.Info("Connected DB")
	return db
}

func initializeApp(e *echo.Echo, db *gorm.DB, logger *logrus.Logger) {
	// Initialize repositories
	logger.Info("Initialize Repositories")
	catRepo := informationRepo.NewCatRepository(db)
	ownerRepo := informationRepo.NewOwnerRepository(db)
	colorRepo := informationRepo.NewColorRepository(db)
	litterRepo := servicesRepo.NewLitterRepository(db, logger)
	breedRepo := informationRepo.NewBreedRepository(db)
	countryRepo := informationRepo.NewCountryRepository(db)
	transferepo := servicesRepo.NewTransferRepository(db)
	logger.Info("Initialize Repositories OK")

	// Initialize handlers
	logger.Info("Initialize Handlers")
	catHandler := informationHandlers.NewCatHandler(catRepo, logger)
	ownerHandler := informationHandlers.NewOwnerHandler(ownerRepo, logger)
	colorHandler := informationHandlers.NewColorHandler(colorRepo, logger)
	litterHandler := servicesHandlers.NewLitterHandler(litterRepo, logger)
	breedHandler := informationHandlers.NewBreedHandler(breedRepo, logger)
	countryHandler := informationHandlers.NewCountryHandler(countryRepo, logger)
	transferHandler := servicesHandlers.NewTransferHandler(transferepo, logger)
	logger.Info("Initialize Handlers OK")

	// Initialize router and routes
	logger.Info("Initialize Router and Routes")
	routes.NewRouter(catHandler, ownerHandler, colorHandler, litterHandler, breedHandler, countryHandler, transferHandler, logger, e)
	logger.Info("Initialize Router and Routes OK")
}
