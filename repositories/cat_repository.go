package repositories

import (
	"amacoonservices/models"
	"fmt"

	"gorm.io/gorm"
)

type CatRepository struct {
	DB *gorm.DB
}

func (r *CatRepository) GetCatsByExhibitorAndSex(idExhibitor int, sex int) ([]models.Cat, error) {
	fmt.Println("Entrou REPO Cats", sex)

	var cats []models.Cat

	query := r.DB.Unscoped().Joins("JOIN racas ON gatos.id_raca = racas.id_racas").
		Joins("JOIN cores ON gatos.id_cor = cores.id_cores").
		Joins("JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
		Joins("JOIN expositores ON gatos.id_expositor= expositores.id_expositores").
		Select("gatos.*, racas.nome AS nome_raca, cores.id_emscode AS id_emscode, cores.descricao AS nome_cor, gatis.nome_gatil AS nome_gatil , expositores.nome AS nome_expositor").
		Where("gatos.id_expositor = ? AND gatos.sexo = ?", idExhibitor, sex).
		Find(&cats)
	if err := query.Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *CatRepository) GetCatByRegistration(registration string) (*models.Cat, error) {
	var cat models.Cat

	query := r.DB.Unscoped().Joins("JOIN racas ON gatos.id_raca = racas.id_racas").
		Joins("JOIN cores ON gatos.id_cor = cores.id_cores").
		Joins("JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
		Joins("JOIN expositores ON gatos.id_expositor= expositores.id_expositores").
		Select("gatos.*, racas.nome AS nome_raca, cores.id_emscode AS id_emscode, cores.descricao AS nome_cor, gatis.nome_gatil AS nome_gatil , expositores.nome AS nome_expositor").
		Where("registro = ?", registration).
		Find(&cat)
	if err := query.Error; err != nil {
		return nil, err
	}

	return &cat, nil
}
