package repositories

import (
	"amacoonservices/models"

	"gorm.io/gorm"
)

type CatRepository struct {
	DB *gorm.DB
}

func (r *CatRepository) GetCatsByExhibitorAndSex(idExhibitor int, sex int) ([]models.Cat, error) {

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

func (r *CatRepository) GetCatsByExhibitorAndSexService(idExhibitor int, sex int) ([]models.CatService, error) {

	var cats []models.CatService

	query := r.DB.Unscoped().Joins("JOIN racas ON gatos.id_raca = racas.id_racas").
		Joins("JOIN cores ON gatos.id_cor = cores.id_cores").
		Joins("JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
		Joins("JOIN expositores ON gatos.id_expositor= expositores.id_expositores").
		Select(`gatos.id_gatos, gatos.nome_do_gato, gatos.registro, gatos.pais_do_gato,
				gatos.microchip, racas.nome AS nome_raca, gatos.id_raca,
				gatos.id_cor, cores.descricao AS nome_cor, cores.id_emscode AS id_emscode,
				gatos.sexo, gatos.nascimento,
				gatos.nome_do_pai, gatos.nome_da_mae, gatis.nome_gatil,
				gatis.criador_gatil, gatos.id_gatil, gatos.id_expositor,
				gatos.criador, expositores.nome AS nome_expositor, expositores.endereco, expositores.cep,
				expositores.cidade, expositores.estado, expositores.telefone`).
		Where("gatos.id_expositor = ? AND gatos.sexo = ?", idExhibitor, sex).
		Find(&cats)

	if err := query.Error; err != nil {
		return nil, err
	}

	return cats, nil
}

func (r *CatRepository) GetCatByRegistrationService(registration string) ([]models.CatService, error) {

	var cats []models.CatService

	query := r.DB.Unscoped().Joins("JOIN racas ON gatos.id_raca = racas.id_racas").
		Joins("JOIN cores ON gatos.id_cor = cores.id_cores").
		Joins("JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
		Joins("JOIN expositores ON gatos.id_expositor= expositores.id_expositores").
		Select(`gatos.id_gatos, gatos.nome_do_gato, gatos.registro,
				gatos.microchip, racas.nome AS nome_raca, gatos.id_raca,
				gatos.id_cor, cores.descricao AS nome_cor, cores.id_emscode AS id_emscode,
				gatos.sexo, gatos.nascimento,
				gatos.nome_do_pai, gatos.nome_da_mae, gatis.nome_gatil,
				gatis.criador_gatil, gatos.id_gatil, gatos.id_expositor,
				gatos.criador, expositores.nome AS nome_expositor, expositores.endereco, expositores.cep,
				expositores.cidade, expositores.estado, expositores.telefone`).
		Where("registro = ?", registration).
		Find(&cats)

	if err := query.Error; err != nil {
		return nil, err
	}

	return cats, nil
}
