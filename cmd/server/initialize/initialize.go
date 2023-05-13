package initialize

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
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
	"github.com/scuba13/AmacoonServices/internal/user"
	"github.com/scuba13/AmacoonServices/internal/utils"
	routes "github.com/scuba13/AmacoonServices/pkg/server"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitializeApp(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, s3Client *s3.S3) {

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
	userRepo := user.NewUserRepository(db, logger)
	logger.Info("Initialize Repositories OK")

	// Initialize services
	logger.Info("Initialize Services")
	filesService := utils.NewFilesService(s3Client, logger)
	protocolService := utils.NewProtocolService(protocolRepo, logger)
	litterService := litter.NewLitterService(litterRepo, logger, protocolService, filesService)
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
	userService := user.NewUserService(userRepo, logger)
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
	filesHandler := handler.NewFilesHandler(filesService, logger)
	userHandler := handler.NewUserHandler(userService, logger)
	logger.Info("Initialize Handlers OK")

	// Initialize router and routes
	logger.Info("Initialize Router and Routes")
	routes.NewRouter(catHandler, ownerHandler, colorHandler,
		litterHandler, breedHandler, countryHandler,
		transferHandler, catteryHandler, federationHandler,
		titleHandler, titleRecognitionHandler, catServiceHandler,
		filesHandler, userHandler,
		logger, e)
	logger.Info("Initialize Router and Routes OK")

}
