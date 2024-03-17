package cat

import (
	"time"

	"gorm.io/gorm"
)

type CatTable struct {
	*gorm.Model
	CatID           int       `gorm:"column:id_gatos;primaryKey"`
	BreedID         string    `gorm:"column:id_raca"`
	BreedName       string    `gorm:"column:nome_raca"`
	OwnerID         int       `gorm:"column:id_expositor;index:idx_1"`
	OwnerName       string    `gorm:"column:nome_expositor"`
	Registration    string    `gorm:"column:registro"`
	RegType         string    `gorm:"column:registro_tipo"`
	RegFed          int       `gorm:"column:registro_federacao"`
	FedName         string    `gorm:"column:nome_federacao"`
	FifeCat         string    `gorm:"column:fifecat"`
	Neutered        string    `gorm:"column:neutro"`
	Name            string    `gorm:"column:nome_do_gato;index:idx_1"`
	Country         string    `gorm:"column:pais_do_gato"`
	WW              string    `gorm:"column:ww"`
	SW              string    `gorm:"column:sw"`
	NW              string    `gorm:"column:nw"`
	AdultTitle      string    `gorm:"column:titulo_adulto"`
	NeuterTitle     string    `gorm:"column:titulo_castrado"`
	JW              string    `gorm:"column:jw"`
	DVM             string    `gorm:"column:dvm"`
	DSM             string    `gorm:"column:dsm"`
	DM              string    `gorm:"column:dm"`
	ColorID         int       `gorm:"column:id_cor"`
	ColorName       string    `gorm:"column:nome_cor"`
	EmsCode         string    `gorm:"column:id_emscode"`
	FatherName      string    `gorm:"column:nome_do_pai"`
	FatherBreed     string    `gorm:"column:raca_do_pai"`
	FatherColor     int       `gorm:"column:cor_do_pai"`
	FatherIdEmscode string    `gorm:"column:fatherId_emscode"`
	MotherName      string    `gorm:"column:nome_da_mae"`
	MotherBreed     string    `gorm:"column:raca_da_mae"`
	MotherColor     int       `gorm:"column:cor_da_mae"`
	MotherIdEmscode string    `gorm:"column:motherId_emscode"`
	BreederName     string    `gorm:"column:nome_gatil"`
	BreederOwner    string    `gorm:"column:criador"`
	BreederID       int       `gorm:"column:id_gatil"`
	Sex             string    `gorm:"column:sexo"`
	BirthDate       time.Time `gorm:"column:nascimento"`
	Microchip       string    `gorm:"column:microchip"`
	Notes           string    `gorm:"column:observacao"`
	Validated       string    `gorm:"column:validado"`
	CreatedAt       string    `gorm:"column:datacadastro"`
	BW              string    `gorm:"column:bw"`
}

func (c *CatTable) TableName() string {
	return "gatos"
}

// func GetCatsMigrate(db *gorm.DB, offset int) ([]CatTable, error) {
//     var cats []CatTable

//     query := db.Unscoped().
//         Joins("LEFT JOIN racas ON gatos.id_raca = racas.id_racas").
//         Joins("LEFT JOIN cores ON gatos.id_cor = cores.id_cores").
//         Joins("LEFT JOIN cores AS cores_pai ON gatos.cor_do_pai = cores_pai.id_cores").
//         Joins("LEFT JOIN cores AS cores_mae ON gatos.cor_da_mae = cores_mae.id_cores").
//         Joins("LEFT JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
//         Joins("LEFT JOIN expositores ON gatos.id_expositor = expositores.id_expositores").
//         Joins("LEFT JOIN federacoes ON gatos.registro_federacao = federacoes.id_federacoes").
//         Select("gatos.*, racas.nome AS nome_raca, cores.id_emscode AS id_emscode, cores.descricao AS nome_cor, cores_pai.id_emscode AS fatherId_emscode, cores_mae.id_emscode AS motherId_emscode, gatis.nome_gatil AS nome_gatil, expositores.nome AS nome_expositor, federacoes.descricao AS nome_federacao").
//         Order("id_gatos ASC").
//         Limit(2). // Isso deve limitar os resultados a 2 registros
//         Offset(offset).
//         Find(&cats)

//     if query.Error != nil {
//         log.Printf("Error fetching cats for migration: %v\n", query.Error)
//         return nil, query.Error
//     }

//     // Log the fetched cats for inspection
//     for i, cat := range cats {
//         log.Printf("Cat %d: ID=%d, Name=%s, BreedName=%s, ColorName=%s, FatherIdEmscode=%s, MotherIdEmscode=%s, Breed_code=%s\n",
//             i+1, cat.CatID, cat.Name, cat.BreedName, cat.ColorName, cat.FatherIdEmscode, cat.MotherIdEmscode, cat.BreedID )
//     }

//     return cats, nil
// }

func GetCatsMigrate(db *gorm.DB, offset, batchSize int) ([]CatTable, error) {
	var cats []CatTable

	query := db.Unscoped().
		Joins("LEFT JOIN racas ON gatos.id_raca = racas.id_racas").
		Joins("LEFT JOIN cores ON gatos.id_cor = cores.id_cores").
		Joins("LEFT JOIN cores AS cores_pai ON gatos.cor_do_pai = cores_pai.id_cores"). // Join adicional para cor do pai
		Joins("LEFT JOIN cores AS cores_mae ON gatos.cor_da_mae = cores_mae.id_cores"). // Join adicional para cor da mãe
		Joins("LEFT JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
		Joins("LEFT JOIN expositores ON gatos.id_expositor = expositores.id_expositores").
		Joins("LEFT JOIN federacoes ON gatos.registro_federacao = federacoes.id_federacoes").
		Select("gatos.*, racas.nome AS nome_raca, cores.id_emscode AS id_emscode, cores.descricao AS nome_cor, cores_pai.id_emscode AS fatherId_emscode, cores_mae.id_emscode AS motherId_emscode, gatis.nome_gatil AS nome_gatil, expositores.nome AS nome_expositor, federacoes.descricao AS nome_federacao").
		Limit(batchSize).Offset(offset).
		Find(&cats)

	if err := query.Error; err != nil {
		return nil, err
	}

	return cats, nil
}
