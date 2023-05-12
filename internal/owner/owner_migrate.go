package owner

import (
	"github.com/scuba13/AmacoonServices/internal/country"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type OwnerS struct {
	gorm.Model
	OwnerID      uint   `gorm:"column:id_expositores;primaryKey"`
    Email        string `gorm:"column:email;unique"`
    PasswordHash string `gorm:"column:senha"`
    OwnerName    string `gorm:"column:nome"`
    Address      string `gorm:"column:endereco"`
    City         string `gorm:"column:cidade"`
    State        string `gorm:"column:estado"`
    ZipCode      string `gorm:"column:cep"`
    Country      string `gorm:"column:pais"`
    Phone        string `gorm:"column:telefone"`
    Valid        string `gorm:"column:valido"`
    ValidationID string `gorm:"column:id_validacao"`
    Observation  []byte `gorm:"column:observacao"`
    CreatedAt    time.Time `gorm:"column:datacadastro"`
    CPF          string `gorm:"column:cpf;default:0"`
}

func MigrateOwners(dbOld, dbNew *gorm.DB, logger *logrus.Logger) error {
    logger.Infof("Migrating owners...")

    // Ler os dados do modelo antigo do banco de dados SQL
    var owners []OwnerS
    if err := dbOld.Unscoped().Table("expositores").Unscoped().Find(&owners).Error; err != nil {
        logger.WithError(err).Error("Failed to get owners from old database")
        return err
    }

    // Converter os dados do modelo antigo para o novo modelo e salvar no banco de dados MySQL
    for _, owner := range owners {
        // Obter o ID do país do banco de dados MySQL para o país do modelo antigo
        var countryID uint
        var country country.Country
        if err := dbNew.Where("code = ?", owner.Country).First(&country).Error; err != nil {
            logger.WithError(err).Errorf("Failed to get country for owner %d", owner.OwnerID)
            return err
        }
        countryID = country.ID

        // Mapear os dados do proprietário do modelo antigo para o modelo novo
        o := Owner{
            Email:        owner.Email,
            PasswordHash: owner.PasswordHash,
            Name:         owner.OwnerName,
            CPF:          owner.CPF,
            Address:      owner.Address,
            City:         owner.City,
            State:        owner.State,
            ZipCode:      owner.ZipCode,
            CountryID:    uintPtr(countryID),
            Phone:        owner.Phone,
            Valid:        owner.Valid == "S",
            ValidId:      owner.ValidationID,
            Observation:  string(owner.Observation),
        }

        // Salvar o proprietário no banco de dados MySQL
        var count int64
        dbNew.Model(&Owner{}).Where("email = ?", o.Email).Count(&count)
        if count == 0 {
            if err := dbNew.Create(&o).Error; err != nil {
                logger.WithError(err).Errorf("Failed to migrate owner %s", o.Email)
                return err
            }
            logger.Infof("Owner %s migrated", o.Email)
        } else {
            logger.Infof("Owner %s already exists in destination database", o.Email)
        }
    }

    logger.Infof("Owners migration completed successfully")
    return nil
}


func uintPtr(n uint) *uint {
	return &n
}