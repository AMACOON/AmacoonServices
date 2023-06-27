package titlerecognition

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TitleRecognitionRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewTitleRecognitionRepository(db *gorm.DB, logger *logrus.Logger) *TitleRecognitionRepository {
	return &TitleRecognitionRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *TitleRecognitionRepository) CreateTitleRecognition(titleRecognition TitleRecognition) (TitleRecognition, error) {
	r.Logger.Infof("Repository CreateTitleRecognition")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the TitleRecognition record
	if err := tx.Create(&titleRecognition).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Error("failed to create TitleRecognition")
		return TitleRecognition{}, err
	}

	// If everything goes well, commit the transaction
	tx.Commit()

	r.Logger.Infof("Repository CreateTitleRecognition OK")
	return titleRecognition, nil
}

func (r *TitleRecognitionRepository) GetTitleRecognitionByID(id uint) (TitleRecognition, error) {
	r.Logger.Infof("Repository GetTitleRecognitionByID")
	var titleRecognition TitleRecognition
	
	err := r.DB.Preload("Titles").Preload("Files").First(&titleRecognition, id).Error
	if err != nil {
		return TitleRecognition{}, err
	}

	r.Logger.Infof("Repository GetTitleRecognitionByID OK")
	return titleRecognition, nil
}

func (r *TitleRecognitionRepository) UpdateTitleRecognition(id uint, titleRecognition TitleRecognition) error {
	r.Logger.Infof("Repository UpdateTitleRecognition")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Atualizar os campos do registro na tabela "title_recognition"
	if err := tx.Model(&TitleRecognition{}).Where("id = ?", id).Updates(titleRecognition).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error updating title recognition with id %v", id)
		return err
	}

	// Atualizar os campos dos títulos relacionados na tabela "service_title_recognition_titles"
	for _, title := range titleRecognition.Titles {
		if err := tx.Model(&Title{}).Where("id = ?", title.ID).Updates(title).Error; err != nil {
			tx.Rollback()
			r.Logger.WithError(err).Errorf("error updating title recognition title record with id %v", title.ID)
			return err
		}
	}

	// If everything goes well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		r.Logger.WithError(err).Errorf("error committing transaction")
		return err
	}

	r.Logger.Infof("Repository UpdateTitleRecognition OK")
	return nil
}



func (r *TitleRecognitionRepository) UpdateTitleRecognitionStatus(id uint, status string) error {
	r.Logger.Infof("Repository UpdateTitleRecognitionStatus")

	var titleRecognition TitleRecognition
    result := r.DB.Model(&titleRecognition).Where("id = ?", id).Update("status", status)
    if result.Error != nil {
        return result.Error
    }

	r.Logger.Infof("Repository UpdateTitleRecognitionStatus OK")
	return nil
}



func (r *TitleRecognitionRepository) GetAllTitleRecognitionByRequesterID(requesterID string) ([]TitleRecognition, error) {
	r.Logger.Infof("Repository GetAllTitlesByRequesterID")
	var titlesRecognition []TitleRecognition
	if err := r.DB.Where("requester_id = ?", requesterID).Find(&titlesRecognition).Error; err != nil {
		return nil, err
	}

	r.Logger.Infof("Repository GetAllTitlesByRequesterID OK")
	return titlesRecognition, nil
}

func (r *TitleRecognitionRepository) DeleteTitleRecognition(id uint) error {
	r.Logger.Infof("Repository DeleteTitleRecognition")

	// Start a new transaction
	tx := r.DB.Begin()

	// Rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete all related titles from the "service_title_recognition_titles" table
	if err := tx.Where("title_recognition_id = ?", id).Delete(&Title{}).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error deleting title recognition title record with title recognition id %v", id)
		return err
	}

	// Delete the record from the "title_recognition" table
	if err := tx.Delete(&TitleRecognition{}, id).Error; err != nil {
		tx.Rollback()
		r.Logger.WithError(err).Errorf("error deleting title recognition by id: %v", id)
		return err
	}

	// If everything goes well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		r.Logger.WithError(err).Errorf("error committing transaction")
		return err
	}

	r.Logger.Infof("Repository DeleteTitleRecognition OK")
	return nil
}





