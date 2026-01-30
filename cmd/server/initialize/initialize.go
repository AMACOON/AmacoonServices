package initialize

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/membership"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/scuba13/AmacoonServices/internal/catshowyear"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
    "github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"
	"github.com/scuba13/AmacoonServices/internal/catservice"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/handler"
	"github.com/scuba13/AmacoonServices/internal/litter"
	"github.com/scuba13/AmacoonServices/internal/login"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/scuba13/AmacoonServices/internal/titlerecognition"
	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"
	routes "github.com/scuba13/AmacoonServices/pkg/server"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitializeApp(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, s3Client *s3.S3) {
	// Initialize repositories
	logger.Info("Initialize Repositories")
	membershipRepo := membership.NewRepository(db, logger)
	catRepo := cat.NewCatRepository(db, logger)
	ownerRepo := owner.NewOwnerRepository(db, logger)
	colorRepo := color.NewColorRepository(db, logger)
	litterRepo := litter.NewLitterRepository(db, logger)
	litterFileRepo:= litter.NewFilesLitterRepository(db, logger)
	breedRepo := breed.NewBreedRepository(db, logger)
	countryRepo := country.NewCountryRepository(db, logger)
	transfeRepo := transfer.NewTransferRepository(db, logger)
	transferFileRepo := transfer.NewFilesTransferRepository(db, logger)
	catteryRepo := cattery.NewCatteryRepository(db, logger)
	catteryFileRepo:= cattery.NewFilesCatteryRepository(db, logger)
	federationRepo := federation.NewFederationRepository(db, logger)
	protocolRepo := utils.NewProtocolRepository(db, logger)
	titleRepo := title.NewTitleRepository(db, logger)
	titleRecognitionRepo := titlerecognition.NewTitleRecognitionRepository(db, logger)
	titleRecognitionFileRepo := titlerecognition.NewFilesTitleRecognitionRepository(db, logger)
	catServiceRepo := catservice.NewCatServiceRepository(db, logger)
	catFileRepo:= cat.NewFilesCatRepository(db, logger)
	loginRepo := login.NewLoginRepository(db, logger)
	catshowRepo := catshow.NewCatShowRepository(db, logger)
	catShowCatRepo:= catshowcat.NewCatShowCatRepository(db, logger)
	catShowCatFileRepo := catshowcat.NewFilesCatShowCatRepository(db, logger)
	catShowRegistrationRepo:= catshowregistration.NewCatShowRegistrationRepository(db, logger)
	catShowResultRepo := catshowresult.NewCatShowResultRepository(db, logger)
	catShowYearRepo := catshowyear.NewCatShowYearRepository(db, logger)
	
	logger.Info("Initialize Repositories OK")

	// Initialize services
	logger.Info("Initialize Services")
	membershipService := membership.NewService(membershipRepo, logger)
	filesService := utils.NewFilesService(s3Client, logger)
	protocolService := utils.NewProtocolService(protocolRepo, logger)
	smptService:= utils.NewSmtpService(logger)
	ownerEmailService := owner.NewOwnerEmailService(smptService, logger)
	catFileService:= cat.NewCatFileService(filesService, catFileRepo, logger)
	litterFileService:= litter.NewFilesLitterService(filesService, litterFileRepo, logger)
	titleRecognitionFileService:= titlerecognition.NewFilesTitleRecognitionService(filesService, titleRecognitionFileRepo, logger)
	transferFileService := transfer.NewFilesTransferService(filesService, transferFileRepo, logger)
	catteryFileService := cattery.NewCatteryFileService(filesService, catteryFileRepo, logger)
	
	litterService := litter.NewLitterService(litterRepo, litterFileService ,protocolService, logger)
	transferService := transfer.NewTransferService(transfeRepo, protocolService, transferFileService, logger)
	catService := cat.NewCatService(catRepo, catFileService, logger)
	breedService := breed.NewBreedService(breedRepo, logger)
	colorService := color.NewColorService(colorRepo, logger)
	countryService := country.NewCountryService(countryRepo, logger)
	ownerService := owner.NewOwnerService(ownerRepo, ownerEmailService,logger)
	catteryService := cattery.NewCatteryService(catteryRepo, catteryFileService, logger)
	federationService := federation.NewCatteryService(federationRepo, logger)
	titleService := title.NewTitleService(titleRepo, logger)
	titleRecognitionService := titlerecognition.NewTitleRecognitionService(titleRecognitionRepo, protocolService, titleRecognitionFileService, logger)
	catServiceService := catservice.NewCatServiceService(catServiceRepo, logger)
	loginService := login.NewLoginService(loginRepo, logger)
	catshowService := catshow.NewCatShowService(catshowRepo, logger)
	catShowCatFileService := catshowcat.NewCatShowCatFileService(filesService, catShowCatFileRepo, logger)
	catShowCatService := catshowcat.NewCCatShowtService(catShowCatRepo, catShowCatFileService,logger)
	catShowRegistrationService := catshowregistration.NewCatShowRegistrationService(logger, catShowCatService, catService, catShowRegistrationRepo)
	catShowResultService := catshowresult.NewCatShowResultService(logger,catShowResultRepo)
	catShowYearService := catshowyear.NewCatShowYearService(logger, catShowYearRepo)
	logger.Info("Initialize Services OK")

	// Initialize handlers
	logger.Info("Initialize Handlers")
	membershipHandler := membership.NewHandler(membershipService)
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
	loginHandler := handler.NewLoginHandler(loginService, logger)
	catshowHandler := handler.NewCatShowHandler(catshowService, logger)
	catShowRegistrationHandler := handler.NewCatShowRegistrationHandler(catShowRegistrationService, logger)
	catShowResultHandler := handler.NewCatShowResultHandler(catShowResultService, logger)
	catShowYearHandler := handler.NewCatShowYearHandler(catShowYearService, logger)
	logger.Info("Initialize Handlers OK")

	// Initialize router and routes
	logger.Info("Initialize Router and Routes")
	routes.NewRouter(
	catHandler, ownerHandler, colorHandler,
	litterHandler, breedHandler, countryHandler,
	transferHandler, catteryHandler, federationHandler,
	titleHandler, titleRecognitionHandler, catServiceHandler,
	filesHandler, loginHandler, catshowHandler, catShowRegistrationHandler,
	catShowResultHandler, catShowYearHandler,

	membershipHandler, // 👈 NOVO

	logger, e)

	logger.Info("Initialize Router and Routes OK")

}
