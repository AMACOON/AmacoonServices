package transfer

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransferRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewTransferRepository(db *gorm.DB, logger *logrus.Logger) *TransferRepository {
	return &TransferRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *TransferRepository) CreateTransfer(transfer Transfer) (Transfer, error) {
	r.Logger.Infof("Repository CreateTransfer")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the transfer record
	if err := tx.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return Transfer{}, err
	}

	// If everything goes well, commit the transaction
	tx.Commit()

	r.Logger.Infof("Repository CreateTransfer OK")
	return transfer, nil
}

func (r *TransferRepository) GetTransferByID(id uint) (Transfer, error) {
	r.Logger.Infof("Repository GetTransferByID")
	var transfer Transfer
	
	err := r.DB.First(&transfer, id).Error
	if err != nil {
		return Transfer{}, err
	}
	r.Logger.Infof("Repository GetTransferByID OK")
	return transfer, nil
}

func (r *TransferRepository) UpdateTransfer(id uint, transfer Transfer) error {
	r.Logger.Infof("Repository UpdateTransfer")

	result := r.DB.Model(&Transfer{}).Where("id = ?", id).Updates(&transfer)
	if result.Error != nil {
		return result.Error
	}

	r.Logger.Infof("Repository UpdateTransfer OK")
	return nil
}


func (r *TransferRepository) UpdateTransferStatus(id uint, status string) error {
    r.Logger.Infof("Repository UpdateTransferStatus")

    var transfer Transfer
    result := r.DB.Model(&transfer).Where("id = ?", id).Update("status", status)
    if result.Error != nil {
        return result.Error
    }

    r.Logger.Infof("Repository UpdateTransferStatus OK")
    return nil
}


func (r *TransferRepository) GetAllTransfersByRequesterID(requesterID uint) ([]Transfer, error) {
	r.Logger.Infof("Repository GetAllTransfersByRequesterID")


	var transfers []Transfer
	result := r.DB.Where("requester_id = ?", requesterID).Find(&transfers)

	if result.Error != nil {
		return nil, result.Error
	}

	r.Logger.Infof("Repository GetAllTransfersByRequesterID OK")
	return transfers, nil
}


