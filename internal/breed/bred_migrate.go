package breed

import (
    "github.com/sirupsen/logrus"
	"gorm.io/gorm"
    "errors"
)

func MigrateBreeds(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
    logger.Infof("Migrating breeds...")

    // Ler os dados do modelo antigo do banco de dados SQL
    var breeds []BreedS
    if err := dbOld.Table("racas").Unscoped().Find(&breeds).Error; err != nil {
        logger.WithError(err).Error("Failed to get breeds from old database")
        return err
    }

    var compatibilities []BreedCompatibilityS
    if err := dbOld.Table("racas_compat").Unscoped().Find(&compatibilities).Error; err != nil {
        logger.WithError(err).Error("Failed to get breed compatibilities from old database")
        return err
    }

    // Converter os dados do modelo antigo para o novo modelo e salvar no banco de dados MySQL
    for _, breed := range breeds {
        // Verificar se a raça já foi migrada anteriormente
        var existingBreed Breed
        if err := dbNew.First(&existingBreed, "breed_code = ?", breed.BreedID).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            logger.WithError(err).Errorf("Failed to check if breed %s already exists in new database", breed.BreedID)
            return err
        }
        if existingBreed.BreedCode != "" {
            logger.Infof("Breed %s already exists in new database, skipping migration", breed.BreedID)
            continue
        }
    
        // Se a raça não foi migrada anteriormente, inserir na tabela de destino
        b := Breed{
            BreedCode:     breed.BreedID,
            BreedName:     breed.BreedName,
            BreedCategory: breed.BreedCategory,
            BreedByGroup:  breed.BreedByGroup,
        }
        if err := dbNew.Create(&b).Error; err != nil {
            logger.WithError(err).Errorf("Failed to migrate breed %s", breed.BreedID)
            return err
        }
        logger.Infof("Breed %s migrated", breed.BreedID)
    }
    

    for _, compatibility := range compatibilities {
        bc := BreedCompatibility{
            BreedCode1: compatibility.IDRaca1,
            BreedCode2: compatibility.IDRaca2,
        }
        if err := dbNew.Create(&bc).Error; err != nil {
            logger.WithError(err).Errorf("Failed to migrate breed compatibility %s-%s", compatibility.IDRaca1, compatibility.IDRaca2)
            return err
        }
        logger.Infof("Breed compatibility %s-%s migrated", compatibility.IDRaca1, compatibility.IDRaca2)
    }

    logger.Infof("Breeds migration completed successfully")
    return nil
}

type BreedS struct {
	*gorm.Model
	BreedID  string `gorm:"column:id_racas;primaryKey"`
	BreedName     string `gorm:"column:nome"`
	BreedCategory int    `gorm:"column:categoria"`
	BreedByGroup  string `gorm:"column:por_grupo"`
}

func (b *BreedS) TableName() string {
	return "racas"
}

type BreedCompatibilityS struct {
	*gorm.Model
	IDRaca1 string `gorm:"primaryKey;column:id_racas1"`
	IDRaca2 string `gorm:"primaryKey;column:id_racas2"`
}

func (b *BreedCompatibilityS) TableName() string {
	return "racas_compat"
}
