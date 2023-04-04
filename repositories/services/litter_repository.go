package repositories

import (
	"amacoonservices/models/services"
	"amacoonservices/utils"

	"gorm.io/gorm"
)

type LitterRepository struct {
    DB             *gorm.DB
    ProtocolService *util.ProtocolService
}

func NewLitterRepository(db *gorm.DB) *LitterRepository {
    protocolService := util.NewProtocolService(db)
    return &LitterRepository{
        DB:              db,
        ProtocolService: protocolService,
    }
}


func (r *LitterRepository) GetAllLitters() ([]models.LitterDB, error) {
	var litters []models.LitterDB
	if err := r.DB.Find(&litters).Error; err != nil {
		return nil, err
	}
	return litters, nil
}

func (r *LitterRepository) GetLitterByID(id uint) (*models.LitterDB, []*models.KittenDB, error) {
	var litter models.LitterDB
	if err := r.DB.First(&litter, id).Error; err != nil {
		return nil, nil, err
	}

	var kittens []*models.KittenDB
	if err := r.DB.Where("id_ninhadas = ?", id).Find(&kittens).Error; err != nil {
		return nil, nil, err
	}

	return &litter, kittens, nil
}

func (r *LitterRepository) CreateLitter(litter *models.LitterDB, kittens []*models.KittenDB) (uint, string, error) {
	
	// Gere o número do protocolo
	protocolNumber, err := r.ProtocolService.GenerateProtocolNumber()
    if err != nil {
        return 0, "", err
    }

    // Associe o número do protocolo à ninhada
    litter.ProtocolNumber = protocolNumber
	
	// Define o status da ninhada como "Pending"
	litter.Status = "Pending"

	// Define o status dos filhotes como "Pending"
	for _, kitten := range kittens {
		kitten.Status = "Pending"
	}

	tx := r.DB.Begin()

	// Cria a ninhada
	if err := tx.Create(litter).Error; err != nil {
		tx.Rollback()
		return 0, "",err
	}

	// Define o LitterID
	for i, kitten := range kittens {
		kitten.LitterID = litter.ID
		kittens[i] = kitten
	}

	// Cria os gatos da ninhada
	if err := tx.Create(&kittens).Error; err != nil {
		tx.Rollback()
		return 0, "",err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, "",err
	}
	
	return litter.ID, protocolNumber,nil
}

func (r *LitterRepository) UpdateLitter(litterID uint, litter *models.LitterDB, kittens []*models.KittenDB) error {
    tx := r.DB.Begin()

    // Atualiza a ninhada
    if err := tx.Model(&models.LitterDB{}).Where("id = ?", litterID).Updates(litter).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Atualiza os filhotes
    for _, kitten := range kittens {
        if err := tx.Model(&models.KittenDB{}).Where("id = ? AND id_ninhadas = ?", kitten.ID, litterID).Updates(kitten).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return err
    }
    return nil
}



func (r *LitterRepository) DeleteLitter(litterID uint) error {
	// Exclui os filhotes associados à ninhada
	if err := r.DB.Where("id_ninhadas = ?", litterID).Delete(&models.KittenDB{}).Error; err != nil {
		return err
	}

	// Exclui a ninhada
	if err := r.DB.Where("id = ?", litterID).Delete(&models.LitterDB{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *LitterRepository) GetKittensByLitterID(litterID uint) ([]*models.KittenDB, error) {
	var kittens []*models.KittenDB
	if err := r.DB.Where("id_ninhadas = ?", litterID).Find(&kittens).Error; err != nil {
		return nil, err
	}
	return kittens, nil
}
