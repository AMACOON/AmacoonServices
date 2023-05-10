package catservice

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CatServiceRepository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

func NewCatServiceRepository(db *gorm.DB, logger *logrus.Logger) *CatServiceRepository {
	return &CatServiceRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatServiceRepository) GetCatServiceByID(id string) (*CatServiceData, error) {
	var cat CatService
	var owner OwnerService
	var catServiceData CatServiceData

	err := r.DB.Table("cats").
		Select("cats.id, cats.name, cats.registration, cats.microchip, breeds.breed_name, colors.name as color_name, colors.ems_code as ems_code, cats.gender, fathers.name as father_name, mothers.name as mother_name").
		Joins("left join breeds on cats.breed_id = breeds.id").
		Joins("left join colors on cats.color_id = colors.id").
		Joins("left join cats fathers on cats.father_id = fathers.id").
		Joins("left join cats mothers on cats.mother_id = mothers.id").
		Where("cats.id = ?", id).
		First(&cat).Error
	if err != nil {
		return &catServiceData, err
	}

	err = r.DB.Table("owners").
		Select("owners.id, owners.name, owners.cpf, owners.address, owners.city, owners.state, owners.zip_code, countries.name as country_name, owners.phone").
		Joins("left join countries on owners.country_id = countries.id").
		Joins("left join cats on owners.id = cats.owner_id").
		Where("cats.id = ?", id).
		First(&owner).Error
	if err != nil {
		return &catServiceData, err
	}

	catServiceData.CatData = cat
	catServiceData.OwnerData = owner

	return &catServiceData, nil
}

func (r *CatServiceRepository) GetAllCatsServiceByOwnerAndGender(ownerID string, gender string) ([]CatServiceData, error) {
	var catServicesData []CatServiceData

	rows, err := r.DB.Table("cats").
		Select("cats.id, cats.name, cats.registration, cats.microchip, breeds.breed_name, colors.name as color_name, colors.ems_code as ems_code, cats.gender, fathers.name as father_name, mothers.name as mother_name, owners.id as owner_id, owners.name as owner_name, owners.cpf, owners.address, owners.city, owners.state, owners.zip_code, countries.name as country_name, owners.phone").
		Joins("left join breeds on cats.breed_id = breeds.id").
		Joins("left join colors on cats.color_id = colors.id").
		Joins("left join cats fathers on cats.father_id = fathers.id").
		Joins("left join cats mothers on cats.mother_id = mothers.id").
		Joins("left join owners on owners.id = cats.owner_id").
		Joins("left join countries on owners.country_id = countries.id").
		Where("cats.owner_id = ? AND cats.gender = ?", ownerID, gender).
		Rows()
	if err != nil {
		return catServicesData, err
	}
	defer rows.Close()

	for rows.Next() {
		var catServiceData CatServiceData
		err = rows.Scan(&catServiceData.CatData.ID, &catServiceData.CatData.Name, &catServiceData.CatData.Registration, &catServiceData.CatData.Microchip, &catServiceData.CatData.BreedName, &catServiceData.CatData.ColorName, &catServiceData.CatData.EmsCode, &catServiceData.CatData.Gender, &catServiceData.CatData.FatherName, &catServiceData.CatData.MotherName, &catServiceData.OwnerData.ID, &catServiceData.OwnerData.Name, &catServiceData.OwnerData.CPF, &catServiceData.OwnerData.Address, &catServiceData.OwnerData.City, &catServiceData.OwnerData.State, &catServiceData.OwnerData.ZipCode, &catServiceData.OwnerData.CountryName, &catServiceData.OwnerData.Phone)
		if err != nil {
			return catServicesData, err
		}
		catServicesData = append(catServicesData, catServiceData)
	}

	return catServicesData, nil
}

func (r *CatServiceRepository) GetAllCatsServiceByOwner(ownerID string) ([]CatServiceData, error) {
	var catServicesData []CatServiceData

	rows, err := r.DB.Table("cats").
		Select("cats.id, cats.name, cats.registration, cats.microchip, breeds.breed_name, colors.name as color_name, colors.ems_code as ems_code, cats.gender, fathers.name as father_name, mothers.name as mother_name, owners.id as owner_id, owners.name as owner_name, owners.cpf, owners.address, owners.city, owners.state, owners.zip_code, countries.name as country_name, owners.phone").
		Joins("left join breeds on cats.breed_id = breeds.id").
		Joins("left join colors on cats.color_id = colors.id").
		Joins("left join cats fathers on cats.father_id = fathers.id").
		Joins("left join cats mothers on cats.mother_id = mothers.id").
		Joins("left join owners on owners.id = cats.owner_id").
		Joins("left join countries on owners.country_id = countries.id").
		Where("cats.owner_id = ?", ownerID).
		Rows()
	if err != nil {
		return catServicesData, err
	}
	defer rows.Close()

	for rows.Next() {
		var catServiceData CatServiceData
		err = rows.Scan(&catServiceData.CatData.ID, &catServiceData.CatData.Name, &catServiceData.CatData.Registration, &catServiceData.CatData.Microchip, &catServiceData.CatData.BreedName, &catServiceData.CatData.ColorName, &catServiceData.CatData.EmsCode, &catServiceData.CatData.Gender, &catServiceData.CatData.FatherName, &catServiceData.CatData.MotherName, &catServiceData.OwnerData.ID, &catServiceData.OwnerData.Name, &catServiceData.OwnerData.CPF, &catServiceData.OwnerData.Address, &catServiceData.OwnerData.City, &catServiceData.OwnerData.State, &catServiceData.OwnerData.ZipCode, &catServiceData.OwnerData.CountryName, &catServiceData.OwnerData.Phone)
		if err != nil {
			return catServicesData, err
		}
		catServicesData = append(catServicesData, catServiceData)
	}

	return catServicesData, nil
}

func (r *CatServiceRepository) GetCatServiceByRegistration(registration string) ([]CatServiceData, error) {
	var catServicesData []CatServiceData

	rows, err := r.DB.Table("cats").
		Select("cats.id, cats.name, cats.registration, cats.microchip, breeds.breed_name, colors.name as color_name, colors.ems_code as ems_code, cats.gender, fathers.name as father_name, mothers.name as mother_name, owners.id as owner_id, owners.name as owner_name, owners.cpf, owners.address, owners.city, owners.state, owners.zip_code, countries.name as country_name, owners.phone").
		Joins("left join breeds on cats.breed_id = breeds.id").
		Joins("left join colors on cats.color_id = colors.id").
		Joins("left join cats fathers on cats.father_id = fathers.id").
		Joins("left join cats mothers on cats.mother_id = mothers.id").
		Joins("left join owners on owners.id = cats.owner_id").
		Joins("left join countries on owners.country_id = countries.id").
		Where("cats.registration = ?", registration).
		Rows()
	if err != nil {
		return catServicesData, err
	}
	defer rows.Close()

	for rows.Next() {
		var catServiceData CatServiceData
		err = rows.Scan(&catServiceData.CatData.ID, &catServiceData.CatData.Name, &catServiceData.CatData.Registration, &catServiceData.CatData.Microchip, &catServiceData.CatData.BreedName, &catServiceData.CatData.ColorName, &catServiceData.CatData.EmsCode, &catServiceData.CatData.Gender, &catServiceData.CatData.FatherName, &catServiceData.CatData.MotherName, &catServiceData.OwnerData.ID, &catServiceData.OwnerData.Name, &catServiceData.OwnerData.CPF, &catServiceData.OwnerData.Address, &catServiceData.OwnerData.City, &catServiceData.OwnerData.State, &catServiceData.OwnerData.ZipCode, &catServiceData.OwnerData.CountryName, &catServiceData.OwnerData.Phone)
		if err != nil {
			return catServicesData, err
		}
		catServicesData = append(catServicesData, catServiceData)
	}

	return catServicesData, nil
}
