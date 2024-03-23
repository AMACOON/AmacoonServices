package catshowresult

import (
	"log"

	//"errors"

	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/scuba13/AmacoonServices/internal/catshowregistration"
	"gorm.io/gorm"
)

type ExposicoesRanking struct {
	IDExposicoesRanking uint `gorm:"primaryKey;column:id_exposicoes_ranking"`
	IDExposicao         uint `gorm:"column:id_exposicao"`
	IDExposicaoSub      uint `gorm:"column:id_exposicao_sub"`
	Numero              int  `gorm:"column:numero"`
	IDGato              uint `gorm:"column:id_gato"`
}

// TableName configura o nome da tabela no banco de dados para GORM
func (ExposicoesRanking) TableName() string {
	return "exposicoes_ranking"
}

type RankingMatrix struct {
	IDRankingMatrix uint   `gorm:"primaryKey;column:id_ranking_matrix"`
	IDExposicao     uint   `gorm:"column:id_exposicao"`
	Descricao       string `gorm:"column:descricao"`
	Pontuacao       int    `gorm:"column:pontuacao"`
}

// TableName configura o nome da tabela no banco de dados para GORM
func (RankingMatrix) TableName() string {
	return "ranking_matrix"
}

type RankingMatrixScore struct {
	IDRankingMatrix uint   `gorm:"column:id_ranking_matrix"`
	IDExposicaoRanking uint `gorm:"primaryKey;column:id_exposicao_ranking"`
}

// TableName configura o nome da tabela no banco de dados para GORM
func (RankingMatrixScore) TableName() string {
	return "exposicoes_ranking_score"
}

func MigrateExposicoesRankingMatrix(dbOld, dbNew *gorm.DB) error {
	var rankingsMatrixs []RankingMatrix
	log.Println("Inicio Migração Resultados Matrix")

	// Buscar todas as inscrições
	if err := dbOld.Find(&rankingsMatrixs).Error; err != nil {
		log.Printf("Erro ao buscar rankingsMatrix: %v", err)
		return err
	}

	// Logando a quantidade de inscrições recuperadas
	log.Printf("Total de rankingsMatrixs encontradas: %d", len(rankingsMatrixs))

	for _, rankingsMatrix := range rankingsMatrixs {
		log.Printf("Iniciando processamento da rankingsMatrixs ID: %d", rankingsMatrix.IDRankingMatrix)

		//Busca  Exposicao pelo ID DB Old
		var exposicao catshow.Exposicao
		if err := dbOld.Where("id_exposicoes = ?", rankingsMatrix.IDExposicao).First(&exposicao).Error; err != nil {
			log.Printf("Erro ao buscar exposicao com ID %d: %v", rankingsMatrix.IDExposicao, err)
			continue // Ou trate o erro conforme necessário
		}
		//Busca ID do Show pela descricao no DB Novo
		var catShow catshow.CatShow
		if err := dbNew.Where("description = ?", exposicao.Descricao).First(&catShow).Error; err != nil {
			log.Printf("Erro ao buscar exposicao com ID %d: %v", exposicao.Descricao, err)
			continue // Ou trate o erro conforme necessário
		}

		catShowResultMatrix := CatShowResultMatrix{
			CatShowID:   &catShow.ID,
			Description: rankingsMatrix.Descricao,
			Score:       rankingsMatrix.Pontuacao,
		}

		if err := dbNew.Create(&catShowResultMatrix).Error; err != nil {
			log.Printf("Erro ao criar registro de catShowResultMatrix ID: %d: %v", &catShow.ID, err)
			// Considerar a inclusão de uma instrução 'continue' ou tratamento de erro aqui
		} else {
			log.Println("catShowResultMatrix criado com sucesso.")
		}

		log.Printf("Processamento da catShowResultMatrix ID: %d concluído com sucesso", catShowResultMatrix.ID)

	}

	log.Printf("Fim Migração Resultados Matrix")
	return nil
}

func MigrateExposicoesRanking(dbOld, dbNew *gorm.DB) error {
	var exposicoesRankings []ExposicoesRanking
	log.Println("Inicio Migração Resultados")

	// Buscar todas as inscrições
	if err := dbOld.Find(&exposicoesRankings).Error; err != nil {
		log.Printf("Erro ao buscar exposicoesRanking: %v", err)
		return err
	}

	// Logando a quantidade de inscrições recuperadas
	log.Printf("Total de exposicoesRanking encontradas: %d", len(exposicoesRankings))

	for _, exposicoesRanking := range exposicoesRankings {
		log.Printf("Iniciando processamento da exposicoesRankings ID: %d", exposicoesRanking.IDExposicoesRanking)

		//Busca  Exposicao pelo ID DB Old
		var exposicao catshow.Exposicao
		if err := dbOld.Where("id_exposicoes = ?", exposicoesRanking.IDExposicao).First(&exposicao).Error; err != nil {
			log.Printf("Erro ao buscar exposicao com ID %d: %v", exposicoesRanking.IDExposicao, err)
			continue // Ou trate o erro conforme necessário
		}
		//Busca ID do Show pela descricao no DB Novo
		var catShow catshow.CatShow
		if err := dbNew.Where("description = ?", exposicao.Descricao).First(&catShow).Error; err != nil {
			log.Printf("Erro ao buscar exposicao com ID %d: %v", exposicao.Descricao, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca  Exposicao Sub pelo ID DB Old
		var exposicaoSub catshow.ExposicaoSub
		if err := dbOld.Where("id_exposicoes_sub = ?", exposicoesRanking.IDExposicaoSub).First(&exposicaoSub).Error; err != nil {
			log.Printf("Erro ao buscar exposicao sub com ID %d: %v", exposicoesRanking.IDExposicaoSub, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca ID do Show Sub pela descricao no DB Novo
		var catShowSub catshow.CatShowSub
		if err := dbNew.Where("description = ?", exposicaoSub.DescricaoExpo).First(&catShowSub).Error; err != nil {
			log.Printf("Erro ao buscar exposicao Sub  com ID %d: %v", exposicaoSub.IDExposicao, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca  exposicoesRankingMatrix pelo ID DB Old
		var rankingMatrixScore RankingMatrixScore
		if err := dbOld.Where("id_exposicao_ranking = ?", exposicoesRanking.IDExposicoesRanking).First(&rankingMatrixScore).Error; err != nil {
			log.Printf("Erro ao buscar RankingMatrixScore com ID %d: %v", exposicoesRanking.IDExposicoesRanking, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca  RankingMatrix pelo ID DB Old
		var rankingMatrix RankingMatrix
		if err := dbOld.Where("id_ranking_matrix = ?", rankingMatrixScore.IDRankingMatrix).First(&rankingMatrix).Error; err != nil {
			log.Printf("Erro ao buscar RankingMatrixOld com ID %d: %v", rankingMatrixScore.IDRankingMatrix, err)
			continue // Ou trate o erro conforme necessário
		}

		// Buscar o  CatShowResultMatrixID no DB Novo
		var catShowResultMatrix CatShowResultMatrix
		if err := dbNew.
			Where("description = ?", rankingMatrix.Descricao).
			Where("cat_show_id = ?", catShow.ID).
			Where("score = ?", rankingMatrix.Pontuacao).
			First(&catShowResultMatrix).Error; err != nil {
			log.Printf("Erro ao buscar catShowResultMatrix  %d: %v", exposicoesRanking.IDGato, err)
			continue // Ou trate o erro conforme necessário
		}



		// Buscar o RegistrationID no DB Novo
		var registration catshowregistration.Registration
		if err := dbNew.
			Where("cat_id_old = ?", exposicoesRanking.IDGato).
			Where("cat_show_id = ?", catShow.ID).
			Where("cat_show_sub_id = ?", catShowSub.ID).
			First(&registration).Error; err != nil {
			log.Printf("Erro ao buscar resgistation com cat_id_old %d: %v", exposicoesRanking.IDGato, err)
			continue // Ou trate o erro conforme necessário
		}

		catShowResult := CatShowResult{
			RegistrationID:        &registration.ID,
			CatShowID:             &catShow.ID,
			CatShowSubID:          &catShowSub.ID,
			Number:                exposicoesRanking.Numero,
			CatShowResultMatrixID: &catShowResultMatrix.ID,
		}

		if err := dbNew.Create(&catShowResult).Error; err != nil {
			log.Printf("Erro ao criar registro de resultado para a inscrição ID: %d: %v", registration.ID, err)
			// Considerar a inclusão de uma instrução 'continue' ou tratamento de erro aqui
		} else {
			log.Println("Resultado criado com sucesso.")
		}

		log.Printf("Processamento da Resultado ID: %d concluído com sucesso", registration.ID)

	}

	log.Printf("Fim Migração Resultado")
	return nil
}
