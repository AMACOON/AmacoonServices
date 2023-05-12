package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/scuba13/AmacoonServices/config"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/catservice"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/handler"
	"github.com/scuba13/AmacoonServices/internal/litter"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/scuba13/AmacoonServices/internal/titlerecognition"
	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"
	routes "github.com/scuba13/AmacoonServices/pkg/server"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Connect to Mongo
	mongo := setupMongo(cfg, logger)
	//s3 := setupS3(cfg, logger)

	db := setupDatabase(cfg, logger)

	db.AutoMigrate(&breed.Breed{},
		&breed.BreedCompatibility{},
		&color.Color{},
		&country.Country{},
		&owner.Owner{},
		&federation.Federation{},
		&cattery.Cattery{},
		&title.Title{},
		&cat.Cat{},
		&cat.TitlesCat{},
		&litter.Litter{},
		&litter.KittenLitter{},
		&transfer.Transfer{},
		&titlerecognition.TitleRecognition{},
		&titlerecognition.TitleData{},
		&utils.Protocol{},
	)

	//dbOld:= setupDatabaseOld(cfg, logger)

	logger.Info("Inicio Migração")

	//breed.MigrateBreeds(dbOld,db, logger)
	//color.MigrateColors(dbOld,db, logger)
	//country.MigrateCountries(dbOld,db, logger)
	//title.InsertTitles(db, logger)
	//owner.MigrateOwners(dbOld,db, logger)
	//federation.MigrateFederations(dbOld,db, logger)
	//cattery.MigrateCattery(dbOld,db, logger, 0.9)
	//cat.MigrateCats(dbOld, db)
	//cat.UpdateCatParents(db)
	//cat.PopulateCatAI(db, mongo)

	logger.Info("Fim Migração")

	// Initialize repositories, handlers, and routes
	initializeApp(e, logger, db, mongo)

	// Start server
	logger.Info("Starting Server")
	if err := e.Start(":" + cfg.ServerPort); err != nil {
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

func setupDatabase(cfg *config.Config, logger *logrus.Logger) *gorm.DB {
	logger.Info("Connecting DB")
	db, err := config.SetupDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to initialize DB connection: %v", err)
	}

	logger.Info("Connected DB")
	return db
}

func setupDatabaseOld(cfg *config.Config, logger *logrus.Logger) *gorm.DB {
	logger.Info("Connecting DB Old")
	db, err := config.SetupDBOld(cfg)
	if err != nil {
		logger.Fatalf("Failed to initialize DB Old connection: %v", err)
	}

	logger.Info("Connected DB Old")
	return db
}

func setupMongo(cfg *config.Config, logger *logrus.Logger) *mongo.Client {
	logger.Info("Connecting Mongo")
	db, err := config.SetupMongo(cfg)
	if err != nil {
		logger.Fatalf("Failed to initialize Mongo connection: %v", err)
	}
	logger.Info("Connected Mongo")
	return db
}

func setupS3(cfg *config.Config, logger *logrus.Logger) *s3.S3 {
	logger.Info("Connecting S3")
	db, err := config.SetupS3Session(cfg)
	if err != nil {
		logger.Fatalf("Failed to initialize s3 connection: %v", err)
	}
	logger.Info("Connected S3")
	return db
}

func initializeApp(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, mongo *mongo.Client) {

	// Initialize repositories
	logger.Info("Initialize Repositories")
	catRepo := cat.NewCatRepository(db, logger)
	ownerRepo := owner.NewOwnerRepository(db, logger)
	colorRepo := color.NewColorRepository(db, logger)
	litterRepo := litter.NewLitterRepository(db, logger)
	breedRepo := breed.NewBreedRepository(db, logger)
	countryRepo := country.NewCountryRepository(db, logger)
	transferepo := transfer.NewTransferRepository(db, logger)
	catteryRepo := cattery.NewCatteryRepository(db, logger)
	federationRepo := federation.NewFederationRepository(db, logger)
	protocolRepo := utils.NewProtocolRepository(db, logger)
	titleRepo := title.NewTitleRepository(db, logger)
	titleRecognitionRepo := titlerecognition.NewTitleRecognitionRepository(db, logger)
	catServiceRepo := catservice.NewCatServiceRepository(db, logger)
	logger.Info("Initialize Repositories OK")

	// Initialize services
	logger.Info("Initialize Services")
	//filesService := utils.NewFilesService(s3, logger)
	protocolService := utils.NewProtocolService(protocolRepo, logger)
	litterService := litter.NewLitterService(litterRepo, logger, protocolService)
	transferService := transfer.NewTransferService(transferepo, logger, protocolService)
	catService := cat.NewCatService(catRepo, logger)
	breedService := breed.NewBreedService(breedRepo, logger)
	colorService := color.NewColorService(colorRepo, logger)
	countryService := country.NewCountryService(countryRepo, logger)
	ownerService := owner.NewOwnerService(ownerRepo, logger)
	catteryService := cattery.NewCatteryService(catteryRepo, logger)
	federationService := federation.NewCatteryService(federationRepo, logger)
	titleService := title.NewTitleService(titleRepo, logger)
	titleRecognitionService := titlerecognition.NewTitleRecognitionService(titleRecognitionRepo, logger, protocolService)
	catServiceService := catservice.NewCatServiceService(catServiceRepo, logger)
	logger.Info("Initialize Services OK")

	// Initialize handlers
	logger.Info("Initialize Handlers")
	catHandler := handler.NewCatHandler(catService, logger)
	ownerHandler := handler.NewOwnerHandler(ownerService, logger)
	colorHandler := handler.NewColorHandler(colorService, logger)
	litterHandler := handler.NewLitterHandler(litterService, logger)
	breedHandler := handler.NewBreedHandler(breedService, logger)
	countryHandler := handler.NewCountryHandler(countryService, logger)
	transferHandler := handler.NewTransferHandler(transferService, logger)
	catteryHandler := handler.NewCatteryHandler(catteryService, logger)
	federationHandler := handler.NewFederationHandler(federationService, logger)
	titleHandler := handler.NewTitleHandler(titleService, logger)
	titleRecognitionHandler := handler.NewTitleRecognitionHandler(titleRecognitionService, logger)
	catServiceHandler := handler.NewCatServiceHandler(catServiceService, logger)
	logger.Info("Initialize Handlers OK")

	// Initialize router and routes
	logger.Info("Initialize Router and Routes")
	routes.NewRouter(catHandler, ownerHandler, colorHandler,
		litterHandler, breedHandler, countryHandler,
		transferHandler, catteryHandler, federationHandler,
		titleHandler, titleRecognitionHandler, catServiceHandler, logger, e)
	logger.Info("Initialize Router and Routes OK")

}
