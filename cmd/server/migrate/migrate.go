package migrate

import (
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
	//"github.com/scuba13/AmacoonServices/internal/breed"
	//"github.com/scuba13/AmacoonServices/internal/color"
//	"github.com/scuba13/AmacoonServices/internal/country"
	//"github.com/scuba13/AmacoonServices/internal/title"

)

func MigrateData(db *gorm.DB, dbOld *gorm.DB, logger *logrus.Logger) {
	logger.Info("Inicio Migração")

	//breed.MigrateBreeds(dbOld,db, logger)
	//color.MigrateColors(dbOld,db, logger)
	//country.MigrateCountries(dbOld,db, logger)
	//title.InsertTitles(db, logger)
	//owner.MigrateOwners(dbOld,db, logger)
	//federation.MigrateFederations(dbOld,db, logger)
	//cattery.MigrateCattery(dbOld,db, logger, 0.9)
	//cat.MigrateCats(dbOld, db)
	//cat.UpdateCatParents(db

	logger.Info("Fim Migração")
}
