package judge

import (
	"errors"

	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JudgesS struct {
	gorm.Model
	IDJuizes int    `gorm:"primary_key;AUTO_INCREMENT;column:id_juizes"`
	Nome     string `gorm:"type:varchar(60);column:nome"`
	Pais     string `gorm:"type:varchar(2);column:pais"`
	Email    string `gorm:"type:varchar(50);column:email"`
	Cat1     string `gorm:"type:varchar(1);column:cat1"`
	Cat2     string `gorm:"type:varchar(1);column:cat2"`
	Cat3     string `gorm:"type:varchar(1);column:cat3"`
	Cat4     string `gorm:"type:varchar(1);column:cat4"`
}

func (JudgesS) TableName() string {
	return "juizes"
}

func MigrateJudges(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
	logger.Infof("Migrating judges...")

	var judgesS []JudgesS
	if err := dbOld.Unscoped().Find(&judgesS).Error; err != nil {
		return err
	}

	for _, judge := range judgesS {
		var countryModel country.Country
		if err := dbNew.Where("code = ?", judge.Pais).First(&countryModel).Error; err != nil {
			return err
		}

		var judgeModel Judge
		if err := dbNew.Where("email = ?", judge.Email).First(&judgeModel).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			judgeModel = Judge{
				Name:      judge.Nome,
				Email:     judge.Email,
		
				CountryID: uintPtr(countryModel.ID),
				Category1A: judge.Cat1 == "s",
				Category1B: judge.Cat1 == "s",
				Category2: judge.Cat2 == "s",
				Category3: judge.Cat3 == "s",
				Category4C: judge.Cat4 == "s",
				Category4D: judge.Cat4 == "s",
			}

			if err := dbNew.Create(&judgeModel).Error; err != nil {
				return err
			}
			logger.Infof("Judge %s migrated", judge.Email)
		} else {
			logger.Infof("Judge %s already exists in destination database", judge.Email)
		}
	}

	logger.Infof("Judges migration completed successfully")
	return nil
}

func uintPtr(n uint) *uint {
	return &n
}
