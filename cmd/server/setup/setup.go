package setup

import (
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/scuba13/AmacoonServices/config"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/litter"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/scuba13/AmacoonServices/internal/titlerecognition"
	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/scuba13/AmacoonServices/internal/user"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.Info("Logger Initialized")
	return logger
}

func SetupDatabase(cfg *config.Config, logger *logrus.Logger) *gorm.DB {
	logger.Info("Connecting DB")
	db, err := config.SetupDB(cfg, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize DB connection: %v", err)
	}

	//logger.Info("Connected DB")

	logger.Info("AutoMigrate DB")
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
		&user.User{},

	)
	logger.Info("AutoMigrate DB OK")
	logger.Info("Connected DB OK")
	return db
}

func SetupDatabaseOld(cfg *config.Config, logger *logrus.Logger) *gorm.DB {
	logger.Info("Connecting DB Old")
	db, err := config.SetupDBOld(cfg)
	if err != nil {
		logger.Fatalf("Failed to initialize DB Old connection: %v", err)
	}

	logger.Info("Connected DB Old")
	return db
}

func SetupS3(cfg *config.Config, logger *logrus.Logger) *s3.S3 {
	logger.Info("Connecting S3")
	db, err := config.SetupS3Session(cfg, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize s3 connection: %v", err)
	}
	logger.Info("Connected S3")
	return db
}
