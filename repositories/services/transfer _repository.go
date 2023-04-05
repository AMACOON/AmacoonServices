package repositories

import (
	"amacoonservices/models/services"
	"amacoonservices/utils"

	"gorm.io/gorm"
)
// passar log
type TransferRepository struct {
	DB *gorm.DB
	ProtocolService *util.ProtocolService
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	protocolService := util.NewProtocolService(db)
	return &TransferRepository{
		DB: db,
		ProtocolService: protocolService,
	}
}

func (r *TransferRepository) CreateCatTransferOwnership(catTransferOwnershipDB *models.TransferDB) (uint, string, error) {
	
	// Gere o número do protocolo
	protocolNumber, err := r.ProtocolService.GenerateProtocolNumber()
    if err != nil {
        return 0, "", err
    }

    // Associe o número do protocolo à ninhada
    catTransferOwnershipDB.ProtocolNumber = protocolNumber
	
	// Define o status da ninhada como "Pending"
	catTransferOwnershipDB.Status = "Pending"
	
	
	
	if err := r.DB.Create(catTransferOwnershipDB).Error; err != nil {
		return 0, "", err
	}
	 return catTransferOwnershipDB.ID, protocolNumber,nil
}

func (r *TransferRepository) GetCatTransferOwnershipByID(id uint) (*models.TransferDB, error) {
	var catTransferOwnership models.TransferDB
	if err := r.DB.First(&catTransferOwnership, id).Error; err != nil {
		return nil, err
	}
	return &catTransferOwnership, nil
}

func (r *TransferRepository) UpdateCatTransferOwnership(id uint, catTransferOwnershipDB *models.TransferDB) error {
	if err := r.DB.Where("id = ?", id).Updates(catTransferOwnershipDB).Error; err != nil {
		return err
	}
	return nil
}


func (r *TransferRepository) DeleteCatTransferOwnership(id uint) error {
    if err := r.DB.Where("id = ?", id).Delete(&models.TransferDB{}).Error; err != nil {
		return err
	}
    return nil
}

