package catshow

import (
	"log"
	"time"

	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
)

// Defina as estruturas para capturar dados de cada tabela
type Exposicao struct {
	IDExposicoes         int       `gorm:"column:id_exposicoes;primaryKey"`
	IDFederacao          int       `gorm:"column:id_federacao"`
	IDClube              int       `gorm:"column:id_clube"`
	Descricao            string    `gorm:"column:descricao"`
	Localizacao          string    `gorm:"column:localizacao"`
	Cidade               string    `gorm:"column:cidade"`
	DataExpoInicio       time.Time `gorm:"column:data_expo_inicio"`
	DataExpoFinal        time.Time `gorm:"column:data_expo_final"`
	Finalizado           string    `gorm:"column:finalizado"`
	DataInscricaoInicio  time.Time `gorm:"column:data_inscricao_inicio"`
	DataInscricaoFinal   time.Time `gorm:"column:data_inscricao_final"`
	Separado             string    `gorm:"column:separado"`
	QtdMaxGatos          int       `gorm:"column:qtd_max_gatos"`
	QtdMaxGatosExpositor int       `gorm:"column:qtd_max_gatos_expositor"`
	Validado             string    `gorm:"column:validado"`
	GatosDesignados      string    `gorm:"column:gatos_designados"`
	Certificado          string    `gorm:"column:certificado"`
	DescricaoDatas       string    `gorm:"column:descricao_datas"`
}

func (Exposicao) TableName() string {
	return "exposicoes"
}

type ExposicaoSub struct {
	IDExposicoesSub int       `gorm:"column:id_exposicoes_sub;primaryKey"`
	IDExposicao     int       `gorm:"column:id_exposicao"`
	NumeroExpo      string    `gorm:"column:numero_expo"`
	DescricaoExpo   string    `gorm:"column:descricao_expo"`
	DataExpo        time.Time `gorm:"column:data_expo"`
	TipoExpo        string    `gorm:"column:tipo_expo"`
}

func (ExposicaoSub) TableName() string {
	return "exposicoes_sub"
}

type ExposicaoJuiz struct {
	IDExposicao int `gorm:"column:id_exposicao"`
	IDJuiz      int `gorm:"column:id_juiz"`
}

func (ExposicaoJuiz) TableName() string {
	return "exposicoes_juizes"
}

func getFederationIDBySigla(db *gorm.DB, sigla string) (uint, error) {
	var federation struct {
		ID uint `gorm:"primaryKey"`
	}
	if err := db.Table("federations").Where("federation_code = ?", sigla).First(&federation).Error; err != nil {
		return 0, err
	}
	return federation.ID, nil
}

// Função para buscar o ID do clube usando o nome
func getClubIDByName(db *gorm.DB, name string) (uint, error) {
	var club struct {
		ID uint `gorm:"primaryKey"`
	}
	if err := db.Table("clubs").Where("name = ?", name).First(&club).Error; err != nil {
		return 0, err
	}
	return club.ID, nil
}

func getJudgeIDByName(db *gorm.DB, name string) (uint, error) {
	var judge struct {
		ID uint `gorm:"primaryKey"`
	}
	// Substitua "judges" pelo nome real da sua tabela de juízes e "nome_juiz" pela coluna que armazena o nome do juiz
	if err := db.Table("judges").Where("name = ?", name).First(&judge).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Tratar caso em que o juiz não é encontrado, se necessário
			return 0, fmt.Errorf("judge not found with name: %s", name)
		}
		// Tratar outros erros de banco de dados
		return 0, err
	}
	return judge.ID, nil
}

// Supondo que a função uintPtr já esteja definida:
func uintPtr(u uint) *uint {
    return &u
}


func MigrateCatShows(dbOld, dbNew *gorm.DB, catShowService *CatShowService) error {
	
	// Passo 1: Buscar exposições
	type ExposicaoDetalhada struct {
		Exposicao    // Herda os campos da estrutura Exposicao
		SiglaFederacao string `gorm:"column:sigla_federacoes"`
		NomeClube     string `gorm:"column:nome"`
	}
	
	var expos []ExposicaoDetalhada
	
	// Realiza uma única consulta com JOINs para buscar exposições e informações relacionadas
	if err := dbOld.Table("exposicoes").
		Select("exposicoes.*, federacoes.sigla_federacoes, clubes.nome").
		Joins("join federacoes on federacoes.id_federacoes = exposicoes.id_federacao").
		Joins("join clubes on clubes.id_clubes = exposicoes.id_clube").
		Find(&expos).Error; err != nil {
		log.Printf("Erro ao buscar exposições: %v", err)
		return err
	}
	

	// if err := dbOld.Table("exposicoes").
	// 	Select("exposicoes.*, federacoes.sigla_federacoes, clubes.nome").
	// 	Joins("join federacoes on federacoes.id_federacoes = exposicoes.id_federacao").
	// 	Joins("join clubes on clubes.id_clubes = exposicoes.id_clube").
	// 	Order("exposicoes.id_exposicoes DESC"). // Ordena corretamente por id_exposicoes
	// 	Limit(2). // Busca os dois últimos registros
	// 	Find(&expos).Error; err != nil {
	// 	log.Printf("Erro ao buscar as últimas exposições: %v", err)
	// 	return err
	// }
	log.Printf("Exposições encontradas: %d", len(expos))
	



	for _, expo := range expos {
		var catShow CatShow
		var countryID uint = 32

		// Atribuição direta de campos que têm correspondência direta em tipos e conceitos
		catShow.Description = expo.Descricao
		catShow.Location = expo.Localizacao
		catShow.City = expo.Cidade
		catShow.State = "SP"           // Suponha que todos os estados sejam SP
		catShow.CountryID = &countryID // Suponha que o país seja Brasil
		catShow.StartDate = expo.DataExpoInicio
		catShow.EndDate = expo.DataExpoFinal
		catShow.RegistrationStart = expo.DataInscricaoInicio
		catShow.RegistrationEnd = expo.DataInscricaoFinal
		catShow.MaxCats = expo.QtdMaxGatos
		catShow.MaxCatsPerExhibitor = expo.QtdMaxGatosExpositor
		catShow.Certificate = expo.Certificado
		catShow.DatesDescription = expo.DescricaoDatas

		// Conversão de valores do tipo string para bool, como por exemplo o campo 'Finalizado'
		catShow.Finished = expo.Finalizado == "s"

		// Conversão de valores do tipo string para bool para os campos 'Separado', 'Validado', 'GatosDesignados'
		catShow.Separated = expo.Separado == "s"
		catShow.Validated = expo.Validado == "s"
		catShow.CatsDesignated = expo.GatosDesignados == "s"

	
		
			// Para FederationID e ClubID, você busca os IDs correspondentes no novo banco de dados
			federationID, err := getFederationIDBySigla(dbNew, expo.SiglaFederacao)
			if err != nil {
				log.Printf("Erro ao buscar ID da federação pela sigla '%s': %v", expo.SiglaFederacao, err)
				// Decida como lidar com o erro (p.ex., continue para a próxima iteração)
				continue
			}
			catShow.FederationID = &federationID

			clubID, err := getClubIDByName(dbNew, expo.NomeClube)
			if err != nil {
				log.Printf("Erro ao buscar ID do clube pelo nome '%s': %v", expo.NomeClube, err)
				// Decida como lidar com o erro (p.ex., continue para a próxima iteração)
				continue
			}
			catShow.ClubID = &clubID
				

				// Convertendo federationID e clubID para *uint e atribuindo aos campos do catShow
				catShow.FederationID = uintPtr(federationID)
				catShow.ClubID = uintPtr(clubID)
			



		// Passo 2: Buscar subexposições
		var subs []ExposicaoSub
		if err := dbOld.Where("id_exposicao = ?", expo.IDExposicoes).Find(&subs).Error; err != nil {
			log.Printf("Failed to fetch subs for exposicao %d: %v", expo.IDExposicoes, err)
			return err // Ou continue, dependendo da política de erro
		}

		// Processar cada subexposição para criar CatShowSubs e adicionar à catShow
		for _, sub := range subs {
			var catShowSub CatShowSub
			// Atribuir campos de ExposicaoSub para CatShowSub

			catShowSub.CatShowNumber, _ = strconv.Atoi(sub.NumeroExpo) // Assumindo que CatShowNumber é int
			catShowSub.Description = sub.DescricaoExpo
			catShowSub.CatShowDate = sub.DataExpo
			catShowSub.CatShowType = sub.TipoExpo // Assumindo que CatShowType é string
			// A conversão de TipoExpo (string) para o tipo equivalente em CatShowSub pode precisar de lógica adicional
			// Por exemplo, se CatShowType for um enum ou um conjunto específico de valores, você pode precisar mapear os valores de `sub.TipoExpo`
			catShowSub.CatShowType = sub.TipoExpo // Assumindo que é uma atribuição direta

			// Adiciona o catShowSub montado à lista de CatShowSubs do catShow
			catShow.CatShowSubs = append(catShow.CatShowSubs, catShowSub)
		}

		// Supondo que esteja dentro do loop das exposições

		// Passo 3: Buscar juízes
		var juizesNomes []struct {
			IDJuiz int
			Nome   string
		}

		// Realiza o join para obter os nomes dos juízes relacionados à exposição
		if err := dbOld.Table("exposicoes_juizes").
			Select("juizes.id_juizes, juizes.nome").
			Joins("join juizes on juizes.id_juizes = exposicoes_juizes.id_juiz").
			Where("exposicoes_juizes.id_exposicao = ?", expo.IDExposicoes).
			Scan(&juizesNomes).Error; err != nil {
			log.Printf("Failed to fetch judges for exposicao %d: %v", expo.IDExposicoes, err)
			return err // Ou continue, dependendo da política de erro
		}

		// Para cada nome de juiz obtido, buscar o ID correspondente no novo banco de dados
		for _, juizNome := range juizesNomes {
			juizID, err := getJudgeIDByName(dbNew, juizNome.Nome)
			if err != nil {
				log.Printf("Failed to fetch judge ID for nome %s: %v", juizNome.Nome, err)
				continue // Ou trate o erro conforme a política de erro
			}

			var catShowJudge CatShowJudge
			// Aqui você atribui o ID do juiz encontrado no novo banco de dados
			catShowJudge.JudgeID = juizID

			// O CatShowID será atribuído depois que catShow for salvo e tiver um ID
			// catShowJudge.CatShowID = catShow.ID

			// Adiciona o catShowJudge montado à lista de CatShowJudges do catShow
			catShow.CatShowJudges = append(catShow.CatShowJudges, catShowJudge)
		}

		logger := logrus.New() // Cria uma nova instância de logger
		catShowRepo := NewCatShowRepository(dbNew, logger) // Supondo que db é sua conexão com o banco de dados
		catShowService := NewCatShowService(catShowRepo, logger)
		catShowService.CreateCatShow(&catShow)

		
		log.Printf("CatShow %s migrated successfully", catShow.Description)

	}

	return nil
}
