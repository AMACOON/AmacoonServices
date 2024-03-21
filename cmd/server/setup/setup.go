package setup

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/scuba13/AmacoonServices/config"

	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
	"github.com/scuba13/AmacoonServices/internal/catshowclass"
	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/club"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/judge"
	"github.com/scuba13/AmacoonServices/internal/litter"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/scuba13/AmacoonServices/internal/titlerecognition"
	"github.com/scuba13/AmacoonServices/internal/transfer"
	"github.com/scuba13/AmacoonServices/internal/utils"
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

type ResponseInterceptor struct {
	http.ResponseWriter
	Body       bytes.Buffer
	StatusCode int
}

// WriteHeader captura o status code da resposta
func (w *ResponseInterceptor) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// Write captura o corpo da resposta
func (w *ResponseInterceptor) Write(b []byte) (int, error) {
	w.Body.Write(b)                  // Armazena a resposta para logar depois
	return w.ResponseWriter.Write(b) // Escreve na resposta original
}

// Header retorna os headers da resposta original
func (w *ResponseInterceptor) Header() http.Header {
	return w.ResponseWriter.Header()
}

func LogMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			// Lendo e copiando o corpo da requisição
			reqBody, err := io.ReadAll(req.Body)
			if err != nil {
				return err
			}
			// Não esqueça de fechar o corpo original após a leitura
			defer req.Body.Close()
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			originalWriter := c.Response().Writer
			interceptor := &ResponseInterceptor{ResponseWriter: originalWriter}
			c.Response().Writer = interceptor

			err = next(c)

			logger.WithFields(logrus.Fields{
				"method":   req.Method,
				"uri":      req.RequestURI,
				"ip":       c.RealIP(),
				"agent":    req.UserAgent(),
				"referer":  req.Referer(),
				"status":   interceptor.StatusCode,
				"request":  string(reqBody),
				"response": interceptor.Body.String(),
			}).Info("request processed")

			return err
		}
	}
}

func SetupDatabase(logger *logrus.Logger) (*gorm.DB, error) {
	logger.Info("Connecting DB")
	db, err := config.SetupDB(logger)
	if err != nil {
		logger.Errorf("Failed to initialize DB connection: %v", err)
		return nil, fmt.Errorf("failed to initialize DB connection: %w", err)
	}

	logger.Info("AutoMigrate DB")
	err = db.AutoMigrate(
		&breed.Breed{},
		&breed.BreedCompatibility{},
		&color.Color{},
		&country.Country{},
		&owner.Owner{},
		&owner.OwnerClub{},
		&federation.Federation{},
		&cattery.Cattery{},
		&cattery.FilesCattery{},
		&title.Title{},
		&cat.Cat{},
		&cat.TitlesCat{},
		&cat.FilesCat{},
		&litter.Litter{},
		&litter.KittenLitter{},
		&litter.FilesLitter{},
		&transfer.Transfer{},
		&transfer.FilesTransfer{},
		&titlerecognition.TitleRecognition{},
		&titlerecognition.Title{},
		&titlerecognition.FilesTitleRecognition{},
		&utils.Protocol{},
		&club.Club{},
		&judge.Judge{},
		&catshow.CatShow{},
		&catshow.CatShowSub{},
		&catshow.CatShowJudge{},
		&catshowclass.Class{},
		&catshowregistration.Registration{},
		&catshowregistration.RegistrationUpdated{},
		&catshowcat.CatShowCat{},
		&catshowcat.FilesCatShowCat{},
		&catshowcat.TitlesCatShowCat{},
	)
	if err != nil {
		logger.Errorf("AutoMigrate failed: %v", err)
		return nil, fmt.Errorf("AutoMigrate failed: %w", err)
	}

	logger.Info("AutoMigrate DB OK")
	logger.Info("Connected DB OK")
	return db, nil
}

func SetupDatabaseOld(logger *logrus.Logger) *gorm.DB {
	logger.Info("Connecting DB Old")
	db, err := config.SetupDBOld()
	if err != nil {
		logger.Fatalf("Failed to initialize DB Old connection: %v", err)
	}

	logger.Info("Connected DB Old")
	return db
}

func SetupS3(logger *logrus.Logger) *s3.S3 {
	logger.Info("Connecting S3")
	db, err := config.SetupS3Session(logger)
	if err != nil {
		logger.Fatalf("Failed to initialize s3 connection: %v", err)
	}
	logger.Info("Connected S3")
	return db
}
