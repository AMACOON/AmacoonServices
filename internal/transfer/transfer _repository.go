package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/utils"

	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
)

// passar log
type TransferRepository struct {
	DB              *gorm.DB
	ProtocolService *utils.ProtocolService
	Logger          *logrus.Logger
}

func NewTransferRepository(db *gorm.DB, logger *logrus.Logger) *TransferRepository {
	protocolService := utils.NewProtocolService(db)
	return &TransferRepository{
		DB:              db,
		ProtocolService: protocolService,
		Logger:          logger,
	}
}

func (r *TransferRepository) CreateTransfer(transferDB *TransferDB, filesDB []*utils.FilesDB) (uint, string, error) {
	// Gera o número do protocolo
	protocolNumber, err := r.ProtocolService.GenerateProtocolNumber()
	if err != nil {
		return 0, "", err
	}

	// Associe o número do protocolo à transferência
	transferDB.ProtocolNumber = protocolNumber

	// Define o status da transferência como "Pending"
	transferDB.Status = "Pending"

	tx := r.DB.Begin()

	// Cria a transferência
	if err := tx.Create(transferDB).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	// Define o ServiceID e ProtocolNumber dos arquivos como o ID e o protocolo da transferência
	for i, file := range filesDB {
		file.ServiceID = transferDB.ID
		file.ProtocolNumber = protocolNumber
		filesDB[i] = file
	}

	// Cria os arquivos da transferência
	if err := tx.Create(&filesDB).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	return transferDB.ID, protocolNumber, nil
}

func (r *TransferRepository) GetTransferByID(id uint) (*TransferDB, []*utils.FilesDB, error) {
	var transferDB TransferDB
	if err := r.DB.First(&transferDB, id).Error; err != nil {
		return nil, nil, err

	}
	var files []*utils.FilesDB
	if err := r.DB.Where("id_serviço = ?", id).Find(&files).Error; err != nil {
		return nil, nil, err
	}
	return &transferDB, files, nil
}

func (r *TransferRepository) UpdateTransfer(id uint, transferDB *TransferDB, filesDB []*utils.FilesDB) error {
	tx := r.DB.Begin()

	// Atualiza a transfer
	if err := tx.Model(&TransferDB{}).Where("id = ?", id).Updates(transferDB).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Atualiza os files
	for _, file := range filesDB {
		if err := tx.Model(&utils.FilesDB{}).Where("numero_protocolo = ? AND id_serviço = ?", transferDB.ProtocolNumber, id).Updates(file).Error; err != nil {
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

func (r *TransferRepository) DeleteTransfer(id uint) error {
	
	// Exclui os files associados à ninhada
	if err := r.DB.Where("id_serviço = ?", id).Delete(&utils.FilesDB{}).Error; err != nil {
		return err
	}
	
	
	if err := r.DB.Where("id = ?", id).Delete(&TransferDB{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransferRepository) GetAlltransfers() ([]TransferDB, error) {
	r.Logger.Info("Repo GetAlltransfers")
	var transfers []TransferDB
	if err := r.DB.Find(&transfers).Error; err != nil {
		return nil, err
	}
	r.Logger.Info("Repo GetAlltransfers OK")
	return transfers, nil
}

