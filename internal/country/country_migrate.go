package country

import (
    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type CountryS struct {
    *gorm.Model
    CountryCode string `gorm:"primaryKey;column:code"`
    CountryName string `gorm:"column:descricao"`
    Activate    string `gorm:"column:visivel"`
}

func (c *CountryS) TableName() string {
    return "country_codes"
}


func MigrateCountries(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
    logger.Infof("Migrating countries...")

    // Ler os dados do modelo antigo do banco de dados SQL
    var countries []CountryS
    if err := dbOld.Table("country_codes").Unscoped().Find(&countries).Error; err != nil {
        logger.WithError(err).Error("Failed to get countries from old database")
        return err
    }

    // Converter os dados do modelo antigo para o novo modelo e salvar no banco de dados MySQL
    for _, country := range countries {
        c := Country{
            Code:        country.CountryCode,
            Name:        country.CountryName,
            IsActivated: country.Activate == "S",
        }
        var count int64
        dbNew.Model(&Country{}).Where("code = ?", c.Code).Count(&count)
        if count == 0 {
            if err := dbNew.Create(&c).Error; err != nil {
                logger.WithError(err).Errorf("Failed to migrate country %s", c.Code)
                return err
            }
            logger.Infof("Country %s migrated", c.Code)
        } else {
            logger.Infof("Country %s already exists in destination database", c.Code)
        }
    }

    logger.Infof("Countries migration completed successfully")
    return nil
}
