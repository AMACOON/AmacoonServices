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
	
	err := r.DB.Preload("Titles").First(&titleRecognition, id).Error
	if err != nil {
		return TitleRecognition{}, err
	}

	r.Logger.Infof("Repository GetTitleRecognitionByID OK")
	return titleRecognition, nil
}

func (r *TitleRecognitionRepository) UpdateTitleRecognition(id uint, titleRecognition TitleRecognition) error {
	r.Logger.Infof("Repository UpdateTitleRecognition")

	err := r.DB.Model(&TitleRecognition{}).Where("id = ?", id).Updates(titleRecognition).Error
	if err != nil {
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




