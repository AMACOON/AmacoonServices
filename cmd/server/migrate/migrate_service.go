package migrate

import (
	"sync"

	"time"

	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/club"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (s *MigrateService) MigrateData1(db *gorm.DB, dbOld *gorm.DB, logger *logrus.Logger) {
	logger.Info("Inicio Migração Breeds, Colors, Countries, Titles, Clubs")

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		breed.MigrateBreeds(dbOld, db, logger)
	}()

	go func() {
		defer wg.Done()
		color.MigrateColors(dbOld, db, logger)
	}()

	go func() {
		defer wg.Done()
		country.MigrateCountries(dbOld, db, logger)
	}()

	go func() {
		defer wg.Done()
		title.InsertTitles(db, logger)
	}()

	go func() {
		defer wg.Done()
		club.MigrateClubs(dbOld, db, logger)
	}()

	wg.Wait()

	logger.Info("Fim Migração Breeds, Colors, Countries, Titles")
}

func (s *MigrateService) MigrateData2(db *gorm.DB, dbOld *gorm.DB, logger *logrus.Logger) {
	logger.Info("Inicio Migração Owner, Owner Club Federation, Cattery")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		owner.MigrateOwners(dbOld, db, logger)
	}()

	go func() {
		defer wg.Done()
		federation.MigrateFederations(dbOld, db, logger)
	}()

	wg.Wait()
	owner.MigrateOwnersClubs(dbOld, db, logger)
	cattery.MigrateCattery(dbOld, db, logger, 0.9)
	logger.Info("Fim Migração Owner,Owner Club, Federation, Cattery")
}

func (s *MigrateService) MigrateData3(dbOld *gorm.DB, db *gorm.DB, logger *logrus.Logger) {
	logger.Info("Iniciando Migração de Cats")

	cat.MigrateCats(dbOld, db)

	logger.Info("Aguardando 10 segundos...")
	time.Sleep(10 * time.Second)

	logger.Info("Migração de Cats concluída. Iniciando atualização de Cat Parents")

	cat.UpdateCatParents(db)

	logger.Info("Atualização de Cat Parents concluída")
}
