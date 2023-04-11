package cat

import (
	"time"

	"gorm.io/gorm"
)

type CatTable struct {
	*gorm.Model
	CatID        int       `gorm:"column:id_gatos;primaryKey"`
	BreedID      string    `gorm:"column:id_raca"`
	BreedName    string    `gorm:"column:nome_raca"`
	OwnerID      int       `gorm:"column:id_expositor;index:idx_1"`
	OwnerName    string    `gorm:"column:nome_expositor"`
	Registration string    `gorm:"column:registro"`
	RegType      string    `gorm:"column:registro_tipo"`
	RegFed       int       `gorm:"column:registro_federacao"`
	FedName      string    `gorm:"column:nome_federacao"`
	FifeCat      string    `gorm:"column:fifecat"`
	Neutered     string    `gorm:"column:neutro"`
	Name         string    `gorm:"column:nome_do_gato;index:idx_1"`
	Country      string    `gorm:"column:pais_do_gato"`
	WW           string    `gorm:"column:ww"`
	SW           string    `gorm:"column:sw"`
	NW           string    `gorm:"column:nw"`
	AdultTitle   string    `gorm:"column:titulo_adulto"`
	NeuterTitle  string    `gorm:"column:titulo_castrado"`
	JW           string    `gorm:"column:jw"`
	DVM          string    `gorm:"column:dvm"`
	DSM          string    `gorm:"column:dsm"`
	DM           string    `gorm:"column:dm"`
	ColorID      int       `gorm:"column:id_cor"`
	ColorName    string    `gorm:"column:nome_cor"`
	EmsCode      string    `gorm:"column:id_emscode"`
	FatherName   string    `gorm:"column:nome_do_pai"`
	FatherBreed  string    `gorm:"column:raca_do_pai"`
	FatherColor  int       `gorm:"column:cor_do_pai"`
	MotherName   string    `gorm:"column:nome_da_mae"`
	MotherBreed  string    `gorm:"column:raca_da_mae"`
	MotherColor  int       `gorm:"column:cor_da_mae"`
	BreederName  string    `gorm:"column:nome_gatil"`
	BreederOwner string    `gorm:"column:criador"`
	BreederID    int       `gorm:"column:id_gatil"`
	Sex          string    `gorm:"column:sexo"`
	BirthDate    time.Time `gorm:"column:nascimento"`
	Microchip    string    `gorm:"column:microchip"`
	Notes        string    `gorm:"column:observacao"`
	Validated    string    `gorm:"column:validado"`
	CreatedAt    string    `gorm:"column:datacadastro"`
	BW           string    `gorm:"column:bw"`
}

func (c *CatTable) TableName() string {
	return "gatos"
}

type Cat struct {
	*gorm.Model
	CatID        int       `gorm:"column:id_gatos;primaryKey"`
	Name         string    `gorm:"column:nome_do_gato;index:idx_1"`
	Registration string    `gorm:"column:registro"`
	Microchip    string    `gorm:"column:microchip"`
	BreedName    string    `gorm:"column:nome_raca"`
	BreedID      string    `gorm:"column:id_raca"`
	ColorID      int       `gorm:"column:id_cor"`
	ColorName    string    `gorm:"column:nome_cor"`
	EmsCode      string    `gorm:"column:id_emscode"`
	Sex          string    `gorm:"column:sexo"`
	BirthDate    time.Time `gorm:"column:nascimento"`
	CountryCode  string    `gorm:"column:pais_do_gato"`
	FatherName   string    `gorm:"column:nome_do_pai"`
	MotherName   string    `gorm:"column:nome_da_mae"`
	BreederName  string    `gorm:"column:nome_gatil"`
	BreederOwner string    `gorm:"column:criador"`
	BreederID    int       `gorm:"column:id_gatil"`
	OwnerID      int       `gorm:"column:id_expositor;index:idx_1"`
	OwnerName    string    `gorm:"column:nome_expositor"`
	Address      string    `gorm:"column:endereco"`
	ZipCode      string    `gorm:"column:cep"`
	City         string    `gorm:"column:cidade"`
	State        string    `gorm:"column:estado"`
	Phone        string    `gorm:"column:telefone"`
}

func (c *Cat) TableName() string {
	return "gatos"
}

type CatMigration struct {
	ID            int       `gorm:"primaryKey;column:id_gatos"`
	BreedID       string    `gorm:"column:id_raca"`
	ExhibitorID   int       `gorm:"column:id_expositor"`
	Registry      string    `gorm:"column:registro"`
	RegistryType  string    `gorm:"column:registro_tipo"`
	FederationID  int       `gorm:"column:registro_federacao"`
	FIFECat       string    `gorm:"column:fifecat"`
	Neutered      string    `gorm:"column:neutro"`
	CatName       string    `gorm:"column:nome_do_gato"`
	CatCountry    string    `gorm:"column:pais_do_gato"`
	WW            string    `gorm:"column:ww"`
	SW            string    `gorm:"column:sw"`
	NW            string    `gorm:"column:nw"`
	AdultTitle    string    `gorm:"column:titulo_adulto"`
	NeuteredTitle string    `gorm:"column:titulo_castrado"`
	JW            string    `gorm:"column:jw"`
	DVM           string    `gorm:"column:dvm"`
	DSM           string    `gorm:"column:dsm"`
	DM            string    `gorm:"column:dm"`
	ColorID       int       `gorm:"column:id_cor"`
	FatherName    string    `gorm:"column:nome_do_pai"`
	FatherBreed   string    `gorm:"column:raca_do_pai"`
	FatherColorID int       `gorm:"column:cor_do_pai"`
	MotherName    string    `gorm:"column:nome_da_mae"`
	MotherBreed   string    `gorm:"column:raca_da_mae"`
	MotherColorID int       `gorm:"column:cor_da_mae"`
	Breeder       string    `gorm:"column:criador"`
	Gender        string    `gorm:"column:sexo"`
	BirthDate     time.Time `gorm:"column:nascimento"`
	CatteryID     int       `gorm:"column:id_gatil"`
	Microchip     string    `gorm:"column:microchip"`
	Observation   []byte    `gorm:"column:observacao"`
	Validated     string    `gorm:"column:validado"`
	CreatedAt     time.Time `gorm:"column:datacadastro"`
	BW            *string   `gorm:"column:bw"`
}

func (c *CatMigration) TableName() string {
	return "gatos"
}
