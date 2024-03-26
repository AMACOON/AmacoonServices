package migrate

import (
	// "time"

	// "github.com/scuba13/AmacoonServices/internal/breed"
	// "github.com/scuba13/AmacoonServices/internal/cat"
	// "github.com/scuba13/AmacoonServices/internal/catshow"
	// "github.com/scuba13/AmacoonServices/internal/catshowclass"
	// "github.com/scuba13/AmacoonServices/internal/catshowregistration"

	// "github.com/scuba13/AmacoonServices/internal/cattery"
	// "github.com/scuba13/AmacoonServices/internal/club"
	// "github.com/scuba13/AmacoonServices/internal/color"
	// "github.com/scuba13/AmacoonServices/internal/country"
	// "github.com/scuba13/AmacoonServices/internal/federation"
	// "github.com/scuba13/AmacoonServices/internal/judge"
	// "github.com/scuba13/AmacoonServices/internal/owner"
	// "github.com/scuba13/AmacoonServices/internal/title"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/internal/catshowresult"
)

type MigrateService struct {
	DB     *gorm.DB
	DBOld  *gorm.DB
	Logger *logrus.Logger
}

func NewMigrateService(db *gorm.DB, dbOld *gorm.DB, logger *logrus.Logger) *MigrateService {
	return &MigrateService{
		DB:     db,
		DBOld:  dbOld,
		Logger: logger,
	}
}

func (s *MigrateService) MigrateData(db *gorm.DB, dbOld *gorm.DB, logger *logrus.Logger) {
	logger.Info("Inicio Migração")

	// logger.Info("Inicio Migração Breed")
	// breed.MigrateBreeds(dbOld, db, logger)
	// logger.Info("Fim Migração Breed")

	// logger.Info("Inicio Migração Color")
	// color.MigrateColors(dbOld, db, logger)
	// logger.Info("Fim Migração Color")

	// logger.Info("Inicio Migração Country")
	// country.MigrateCountries(dbOld, db, logger)
	// logger.Info("Fim Migração Country")

	// logger.Info("Inicio Migração Title")
	// title.InsertTitles(db, logger)
	// logger.Info("Fim Migração Title")

	// logger.Info("Inicio Migração Club")
	// club.MigrateClubs(dbOld, db, logger)
	// logger.Info("Fim Migração Club")

	// logger.Info("Inicio Migração Judges")
	// judge.MigrateJudges(dbOld, db, logger)
	// logger.Info("Fim Migração Judges")

	// logger.Info("Inicio Migração Owner")
	// owner.MigrateOwners(dbOld, db, logger)
	// logger.Info("Fim Migração Owner")

	// logger.Info("Inicio Migração Federation")
	// federation.MigrateFederations(dbOld, db, logger)
	// logger.Info("Fim Migração Federation")

	// logger.Info("Inicio Migração Owner Club")
	// owner.MigrateOwnersClubs(dbOld, db, logger)
	// logger.Info("Fim Migração Owner Club")

	// logger.Info("Inicio Migração Cattery")
	// cattery.MigrateCattery(dbOld, db, logger, 0.9)
	// logger.Info("Fim Migração Cattery")

	// logger.Info("Inicio Migração Cat")
	// cat.MigrateCats(dbOld, db)
	// logger.Info("Fim Migração Cat")

	// logger.Info("Aguardando 10 segundos...")
	// time.Sleep(10 * time.Second)

	// logger.Info("Inicio Migração Cat Parents")
	// cat.UpdateCatParents(db)
	// logger.Info("Fim Migração Cat Parents")

	// logger.Info("Inicio Migração Cat Show")
	// catshow.MigrateCatShows(dbOld, db, &catshow.CatShowService{})
	// logger.Info("Fim Migração Cat Show")

	// logger.Info("Inicio Migração Class")
	// catshowclass.MigrateClasses(dbOld, db, logger)
	// logger.Info("Fim Migração Class")

	// logger.Info("Inicio Migração Cat Show Registration")
	// catshowregistration.MigrateInscricoes(dbOld, db)
	// logger.Info("Fim Migração Cat Show Registration")

	// logger.Info("Inicio Migração Cat Show Registration Updated")
	// catshowregistration.MigrateInscricoesUpdate(dbOld, db)
	// logger.Info("Fim Migração Cat Show Registration Updated")

	// logger.Info("Inicio Migração ExposicoesRankingMatrix")
	// catshowresult.MigrateExposicoesRankingMatrix(dbOld, db)
	// logger.Info("Fim Migração ExposicoesRankingMatrix")

	logger.Info("Inicio Migração ExposicoesRanking")
	catshowresult.MigrateExposicoesRanking(dbOld, db)
	logger.Info("Fim Migração ExposicoesRanking")

	logger.Info("Fim Migração")
}
