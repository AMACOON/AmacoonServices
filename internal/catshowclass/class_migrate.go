package catshowclass

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)



// OldClass representa a estrutura da tabela `classes` conforme definido no banco de dados antigo.
type OldClass struct {
	IDClasses        string `gorm:"type:varchar(3);primaryKey"`
	Classe           string `gorm:"type:varchar(10);not null"`
	DescricaoClasse  string `gorm:"type:varchar(50);not null"`
	Ordem            int    `gorm:"not null"`
	OrdemNew         int    `gorm:"not null"`
}

func (OldClass) TableName() string {
	return "classes"
}


// MigrateClasses responsável pela migração dos dados de classes de um banco de dados antigo para um novo.
func MigrateClasses(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
	logger.Infof("Iniciando a migração de classes...")

	var oldClasses []OldClass
	if err := dbOld.Find(&oldClasses).Error; err != nil {
		return err
	}

	for _, oldClass := range oldClasses {
		newClass := Class{
			Code: 	  oldClass.IDClasses,
			Name:        oldClass.Classe,
			Description: oldClass.DescricaoClasse,
			Order:       oldClass.Ordem,
			NewOrder:    oldClass.OrdemNew,
		}

		// Verifica se a classe já existe no novo banco de dados para evitar duplicatas
		var existingClass Class
		if err := dbNew.Where("code = ?", newClass.Name).First(&existingClass).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Cria uma nova entrada se a classe não existir
				if err := dbNew.Create(&newClass).Error; err != nil {
					return err
				}
				logger.Infof("Classe %s migrada com sucesso.", oldClass.Classe)
			} else {
				return err
			}
		} else {
			logger.Infof("Classe %s já existe no novo banco de dados.", oldClass.Classe)
		}
	}

	logger.Infof("Migração de classes concluída com sucesso.")
	return nil
}
