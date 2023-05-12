package cattery

import (
	"github.com/agext/levenshtein"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/internal/country"
	
	"github.com/scuba13/AmacoonServices/internal/owner"
)

type CatteryS struct {
	ID             string `gorm:"column:id_gatis;primaryKey"`
	Name           string `gorm:"column:nome_gatil"`
	BreederName    string `gorm:"column:criador_gatil"`
	BreederCountry string `gorm:"column:pais_gatil"`
}

func (b *CatteryS) TableName() string {
	return "gatis"
}

func MigrateCattery(dbOld, dbNew *gorm.DB, logger *logrus.Logger, minSimilarity float64) error {
	logger.Infof("Migrating catteries...")

	var catteriesSQL []CatteryS
	if err := dbOld.Unscoped().Table("gatis").Scan(&catteriesSQL).Error; err != nil {
		return err
	}

	var owners []owner.Owner
	if err := dbNew.Find(&owners).Error; err != nil {
		return err
	}

	var countries []country.Country
	if err := dbNew.Find(&countries).Error; err != nil {
		return err
	}

	// Create country code to ID map for faster lookups
	countryCodeToID := make(map[string]uint)
	for _, country := range countries {
		countryCodeToID[country.Code] = country.ID
	}

	// Create owner name to ID map for faster lookups
	ownerNameToID := make(map[string]uint)
	for _, owner := range owners {
		ownerNameToID[owner.Name] = owner.ID
	}

	for _, catterySQL := range catteriesSQL {
		var ownerID uint
		var maxSimilarity float64 = 0.0

		for _, owner := range owners {
			ownerName := owner.Name
			breederName := catterySQL.BreederName
			distance := levenshtein.Distance(ownerName, breederName, nil)
			similarity := 1 - float64(distance)/float64(Max(len(ownerName), len(breederName)))
			if similarity > maxSimilarity && similarity >= minSimilarity {
				maxSimilarity = similarity
				ownerID = owner.ID
			}
		}

		countryID, found := countryCodeToID[catterySQL.BreederCountry]
		if !found {
			logger.Warnf("country not found for cattery with breeder ID %s and country code %s", catterySQL.ID, catterySQL.BreederCountry)
		}

		catteryModel := &Cattery{
			Name:        catterySQL.Name,
			BreederName: catterySQL.BreederName,
			OwnerID:      uintPtr(ownerID),
			CountryID:    uintPtr(countryID),
		}

		if err := dbNew.Create(&catteryModel).Error; err != nil {
			return err
		}
	}

	logger.Infof("Catteries migration completed successfully")
	return nil
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func uintPtr(n uint) *uint {
	return &n
}
