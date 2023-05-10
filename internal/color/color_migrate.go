package color

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	
)

type ColorS struct {
    *gorm.Model
    ColorID   int    `gorm:"column:id_cores;primaryKey"`
    BreedID   string `gorm:"column:id_raca"`
    EmsCode   string `gorm:"column:id_emscode"`
    ColorName string `gorm:"column:descricao"`
    Group     int    `gorm:"column:grupo"`
    SubGroup  int    `gorm:"column:sub_grupo"`
}

func (c *ColorS) TableName() string {
    return "cores"
}

func MigrateColors(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
    logger.Infof("Migrating colors...")

    batchSize := 1000 // Defina o tamanho do lote aqui
    var offset int
    for {
        // Ler os dados do modelo antigo do banco de dados SQL
        var colors []ColorS
        if err := dbOld.Table("cores").Offset(offset).Limit(batchSize).Unscoped().Find(&colors).Error; err != nil {
            logger.WithError(err).Error("Failed to get colors from old database")
            return err
        }

        // Sair do loop se não houver mais dados a serem lidos
        if len(colors) == 0 {
            break
        }

        // Converter os dados do modelo antigo para o novo modelo e salvar no banco de dados MySQL
        for _, color := range colors {
            c := Color{
                BreedCode: color.BreedID,
                EmsCode:   color.EmsCode,
                Name:      color.ColorName,
                Group:     color.Group,
                SubGroup:  color.SubGroup,
            }
            if err := dbNew.Create(&c).Error; err != nil {
                logger.WithError(err).Errorf("Failed to migrate color %d", color.ColorID)
                return err
            }
            logger.Infof("Color %d migrated", color.ColorID)
        }

        offset += batchSize
    }

    logger.Infof("Colors migration completed successfully")
    return nil
}
