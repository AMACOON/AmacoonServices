package federation

import (
	"github.com/scuba13/AmacoonServices/internal/country"

	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FederationSQL struct {
	ID             string `gorm:"column:id_federacoes;primaryKey"`
	FederationCode string `gorm:"column:sigla_federacoes"`
	FederationName string `gorm:"column:descricao"`
	CountryCode    string `gorm:"column:country_code"`
}

func (f *FederationSQL) TableName() string {
	return "federacoes"
}

func MigrateFederations(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
	logger.Infof("Migrating federations...")

	var federationsSQL []FederationSQL
	if err := dbOld.Find(&federationsSQL).Error; err != nil {
		return err
	}

	for _, federationSQL := range federationsSQL {
		var countryModel country.Country
		if err := dbNew.Where("code = ?", federationSQL.CountryCode).First(&countryModel).Error; err != nil {
			return err
		}

		var federationModel Federation
		if err := dbNew.Where("federation_code = ?", federationSQL.FederationCode).First(&federationModel).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			federationModel = Federation{
				Name:           federationSQL.FederationName,
				FederationCode: federationSQL.FederationCode,
				CountryID:      uintPtr(countryModel.ID),
				
			}

			if err := dbNew.Create(&federationModel).Error; err != nil {
				return err
			}
			logger.Infof("Federation %s migrated", federationSQL.FederationCode)
		} else {
			logger.Infof("Federation %s already exists in destination database", federationSQL.FederationCode)
		}
	}

	logger.Infof("Federations migration completed successfully")
	return nil
}

func uintPtr(n uint) *uint {
	return &n
}