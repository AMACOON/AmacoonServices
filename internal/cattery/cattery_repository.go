package cattery

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatteryRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatteryRepository(db *gorm.DB, logger *logrus.Logger) *CatteryRepository {
	return &CatteryRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatteryRepository) GetCatteryByID(id string) (*Cattery, error) {
	r.Logger.Infof("Repository GetCatteryByID")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid id: %s", id)
		return nil, err
	}
	var cattery Cattery
	if err := r.DB.
		Preload("Owner").
		Preload("Country").
		Preload("Files").
		First(&cattery, uintID).
		Error; err != nil {
		r.Logger.WithError(err).Errorf("failed to get cattery with ID %s", id)
		return nil, err
	}
	r.Logger.Infof("Repository GetCatteryByID OK")
	return &cattery, nil
}

func (r *CatteryRepository) GetAllCatteries() ([]CatteryInfo, error) {
	r.Logger.Infof("Repository GetAllCatteries")
	var catteriesInfo []CatteryInfo

	err := r.DB.Table("catteries").
		Select("catteries.name, catteries.breeder_name, countries.name as country_name").
		Joins("JOIN countries ON catteries.country_id = countries.id").
		Scan(&catteriesInfo).Error

	if err != nil {
		r.Logger.WithError(err).Errorf("failed to get all catteries")
		return nil, err
	}
	r.Logger.Infof("Repository GetAllCatteries OK")
	return catteriesInfo, nil
}

func (r *CatteryRepository) CreateCattery(cattery *Cattery) (*Cattery, error) {
	r.Logger.Infof("Repository CreateCattery")

	if err := r.DB.Create(cattery).Error; err != nil {
		r.Logger.Errorf("Error creating cattery: %v", err)
		return nil, err
	}

	r.Logger.Infof("Repository CreateCattery OK")
	return cattery, nil
}

func (r *CatteryRepository) UpdateCattery(id string, cattery *Cattery) (*Cattery, error) {
	r.Logger.Infof("Repository UpdateCattery")

	// Primeiro, verifica se a Cattery com o ID fornecido existe
	existingCattery := &Cattery{}
	if err := r.DB.First(existingCattery, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.Logger.Errorf("No cattery found with id: %d", id)
			return nil, err
		}
		r.Logger.Errorf("Error retrieving cattery: %v", err)
		return nil, err
	}

	// Se existir, atualiza com os novos dados
	idInt, err := strconv.Atoi(id)
	if err != nil {
		r.Logger.Errorf("Error converting string to int: %v", err)
		return nil, err
	}
	cattery.ID = uint(idInt)

	if err := r.DB.Save(cattery).Error; err != nil {
		r.Logger.Errorf("Error updating cattery: %v", err)
		return nil, err
	}

	r.Logger.Infof("Repository UpdateCattery OK")
	return cattery, nil
}

func (r *CatteryRepository) ValidateCattery(name string) (*Cattery, error) {
	r.Logger.Infof("Repository ValidateCattery")

	var cattery Cattery
	if err := r.DB.Where("name = ?", name).First(&cattery).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &cattery, nil
}

