package litter

import (
	
	"github.com/scuba13/AmacoonServices/internal/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LitterRepository struct {
	DB              *gorm.DB
	ProtocolService *utils.ProtocolService
	Logger          *logrus.Logger
}

func NewLitterRepository(db *gorm.DB, logger *logrus.Logger) *LitterRepository {
	protocolService := utils.NewProtocolService(db)
	return &LitterRepository{
		DB:              db,
		ProtocolService: protocolService,
		Logger:          logger,
	}
}

func (r *LitterRepository) GetAllLitters() ([]LitterDB, error) {
	r.Logger.Info("Repo GetAllLitters")
	var litters []LitterDB
	if err := r.DB.Find(&litters).Error; err != nil {
		return nil, err
	}
	r.Logger.Info("Repo GetAllLitters OK")
	return litters, nil
}

func (r *LitterRepository) GetLitterByID(id uint) (*LitterDB, []*KittenDB, []*utils.FilesDB, error) {
	var litter LitterDB
	if err := r.DB.First(&litter, id).Error; err != nil {
		return nil, nil, nil, err
	}

	var kittens []*KittenDB
	if err := r.DB.Where("id_ninhadas = ?", id).Find(&kittens).Error; err != nil {
		return nil, nil, nil, err
	}

	var files []*utils.FilesDB
	if err := r.DB.Where("id_serviço = ?", litter.ID).Find(&files).Error; err != nil {
		return nil, nil, nil, err
	}

	return &litter, kittens, files, nil
}

func (r *LitterRepository) CreateLitter(litterDB *LitterDB, kittensDB []*KittenDB, filesDB []*utils.FilesDB) (uint, string, error) {

	// Gere o número do protocolo
	protocolNumber, err := r.ProtocolService.GenerateProtocolNumber()
	if err != nil {
		return 0, "", err
	}

	// Associe o número do protocolo à ninhada
	litterDB.ProtocolNumber = protocolNumber

	// Define o status da ninhada como "Pending"
	litterDB.Status = "Pending"

	// Define o status dos filhotes como "Pending"
	for _, kitten := range kittensDB {
		kitten.Status = "Pending"
	}

	tx := r.DB.Begin()

	// Cria a ninhada
	if err := tx.Create(litterDB).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	// Define o LitterID
	for i, kitten := range kittensDB {
		kitten.LitterID = litterDB.ID
		kittensDB[i] = kitten
	}

	// Cria os gatos da ninhada
	if err := tx.Create(&kittensDB).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	// Define o ServiceID dos FilesDB como o ID da ninhada
	for i, file := range filesDB {
		file.ServiceID = litterDB.ID
		file.ProtocolNumber = protocolNumber
		filesDB[i] = file

	}

	// Cria os arquivos da ninhada
	if err := tx.Create(&filesDB).Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, "", err
	}

	return litterDB.ID, protocolNumber, nil
}

func (r *LitterRepository) UpdateLitter(litterID uint, litterDB *LitterDB, kittensDB []*KittenDB, filesDB []*utils.FilesDB) error {
	tx := r.DB.Begin()

	// Atualiza a ninhada
	if err := tx.Model(&LitterDB{}).Where("id = ?", litterID).Updates(litterDB).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Atualiza os filhotes
	for _, kitten := range kittensDB {
		if err := tx.Model(&KittenDB{}).Where("id = ? AND id_ninhadas = ?", kitten.ID, litterID).Updates(kitten).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Atualiza os files
	for _, file := range filesDB {
		if err := tx.Model(&utils.FilesDB{}).Where("numero_protocolo = ? AND id_serviço = ?", litterDB.ProtocolNumber, litterID).Updates(file).Error; err != nil {
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
	if err := r.DB.Where("id_ninhadas = ?", litterID).Delete(&KittenDB{}).Error; err != nil {
		return err
	}

	// Exclui os files associados à ninhada
	if err := r.DB.Where("id_serviço = ?", litterID).Delete(&utils.FilesDB{}).Error; err != nil {
		return err
	}

	// Exclui a ninhada
	if err := r.DB.Where("id = ?", litterID).Delete(&LitterDB{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *LitterRepository) GetKittensByLitterID(litterID uint) ([]*KittenDB, error) {
	var kittens []*KittenDB
	if err := r.DB.Where("id_ninhadas = ?", litterID).Find(&kittens).Error; err != nil {
		return nil, err
	}
	return kittens, nil
}


