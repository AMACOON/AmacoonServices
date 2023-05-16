package migrate

import (
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
	//"github.com/scuba13/AmacoonServices/internal/breed"
	// "github.com/scuba13/AmacoonServices/internal/color"
	// "github.com/scuba13/AmacoonServices/internal/country"
	// "github.com/scuba13/AmacoonServices/internal/title"
	//"github.com/scuba13/AmacoonServices/internal/owner"
	//"github.com/scuba13/AmacoonServices/internal/cattery"
	//"github.com/scuba13/AmacoonServices/internal/federation"
	//"github.com/scuba13/AmacoonServices/internal/cat"
	//"github.com/scuba13/AmacoonServices/internal/user"
	//"time"

)

func MigrateData(db *gorm.DB, dbOld *gorm.DB, logger *logrus.Logger) {
	logger.Info("Inicio Migração")

	//breed.MigrateBreeds(dbOld,db, logger)
	//color.MigrateColors(dbOld,db, logger)
	//country.MigrateCountries(dbOld,db, logger)
	//title.InsertTitles(db, logger)
	//owner.MigrateOwners(dbOld,db, logger)
	//user.MigrateOwnersToUsers(db, logger)
	//federation.MigrateFederations(dbOld,db, logger)
	//cattery.MigrateCattery(dbOld,db, logger, 0.9)
	//cat.MigrateCats(dbOld, db)
	//time.Sleep(30 * time.Second)
	//cat.UpdateCatParents(db)

	logger.Info("Fim Migração")
}
