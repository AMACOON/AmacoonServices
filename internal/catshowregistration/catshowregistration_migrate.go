package catshowregistration

import (
	"log"
	"time"

	//"errors"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/internal/catshow"
	"github.com/scuba13/AmacoonServices/internal/catshowcat"
	"github.com/scuba13/AmacoonServices/internal/catshowclass"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/judge"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Inscricao struct {
	IDInscricoes   uint      `gorm:"primaryKey;column:id_inscricoes"`
	IDExposicao    uint      `gorm:"column:id_exposicao"`
	IDExposicaoSub uint      `gorm:"column:id_exposicao_sub"`
	IDExpositor    uint      `gorm:"column:id_expositor"`
	IDGato         uint      `gorm:"column:id_gato"`
	IDClasse       string    `gorm:"column:id_classe"`
	DataCadastro   time.Time `gorm:"column:datacadastro"`
	IDJuiz         uint      `gorm:"column:id_juiz"`
	Numero         int       `gorm:"column:numero"`
	Observacao     []byte    `gorm:"column:observacao"`
}

// TableName configura o nome da tabela no banco de dados para GORM
func (Inscricao) TableName() string {
	return "inscricoes"
}

type InscricaoUpdate struct {
	IDExposicao    uint      `gorm:"column:id_exposicao"`     // Identificador da exposição
	IDExposicaoSub uint      `gorm:"column:id_exposicao_sub"` // Identificador da sub-exposição
	IDGato         uint      `gorm:"column:id_gato"`          // Identificador do gato
	Numero         int       `gorm:"column:numero"`           // Número da inscrição
	IDClasse       uint      `gorm:"column:r_id_classe"`      // Referência para a classe
	IDCor          uint      `gorm:"column:r_id_cor"`         // Referência para a cor
	IDJuiz         uint      `gorm:"column:r_id_juiz"`        // Referência para o juiz
	Nascimento     time.Time `gorm:"column:r_nascimento"`     // Data de nascimento
	Sexo           string    `gorm:"column:r_sexo"`           // Sexo
}

// TableName configura o nome da tabela no banco de dados para GORM
func (InscricaoUpdate) TableName() string {
	return "inscricoes_alteracoes"
}

func MigrateInscricoes(dbOld, dbNew *gorm.DB) error {
	var inscricoes []Inscricao
	log.Println("Inicio Migração Inscricoes")

	// Buscar todas as inscrições
	if err := dbOld.Find(&inscricoes).Error; err != nil {
		log.Printf("Erro ao buscar inscrições: %v", err)
		return err
	}

	// Logando a quantidade de inscrições recuperadas
	log.Printf("Total de inscrições encontradas: %d", len(inscricoes))

	for _, inscricao := range inscricoes {
		log.Printf("Iniciando processamento da inscrição ID: %d", inscricao.IDInscricoes)
		// Buscar o CatID correspondente na tabela 'cats' usando IDGato
		var catData cat.Cat
		if err := dbNew.Where("cat_id_old = ?", inscricao.IDGato).First(&catData).Error; err != nil {
			log.Printf("Erro ao buscar gato com ID %d: %v", inscricao.IDGato, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca  Exposicao pelo ID DB Old
		var exposicao catshow.Exposicao
		if err := dbOld.Where("id_exposicoes = ?", inscricao.IDExposicao).First(&exposicao).Error; err != nil {
			log.Printf("Erro ao buscar exposicao com ID %d: %v", inscricao.IDExposicao, err)
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
		if err := dbOld.Where("id_exposicoes_sub = ?", inscricao.IDExposicaoSub).First(&exposicaoSub).Error; err != nil {
			log.Printf("Erro ao buscar exposicao sub com ID %d: %v", inscricao.IDExposicaoSub, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca ID do Show Sub pela descricao no DB Novo
		var catShowSub catshow.CatShowSub
		if err := dbNew.Where("description = ?", exposicaoSub.DescricaoExpo).First(&catShowSub).Error; err != nil {
			log.Printf("Erro ao buscar exposicao Sub  com ID %d: %v", exposicaoSub.IDExposicao, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca Expositor pelo ID DB Old
		var expositor owner.OwnerS
		// Ajustando a consulta para incluir um campo 'id' através de um alias para 'id_expositores'
		query := "SELECT id_expositores AS id, expositores.* FROM expositores WHERE id_expositores = ? ORDER BY id_expositores LIMIT 1"
		if err := dbOld.Raw(query, inscricao.IDExpositor).Scan(&expositor).Error; err != nil {
			log.Printf("Erro ao buscar expositor com ID %d: %v", inscricao.IDExpositor, err)
			continue
		}

		//Busca ID do Owner pelo email no DB Novo
		var owner owner.Owner
		if err := dbNew.Where("email = ?", expositor.Email).First(&owner).Error; err != nil {
			log.Printf("Erro ao buscar expositor com ID %d: %v", expositor.Email, err)
			continue
		}

		//Busca Juiz pelo ID no DB Antigo
		var juiz judge.JudgesS
		// Ajustando a consulta para incluir um campo 'id' através de um alias para 'id_juizes'
		queryJ := "SELECT id_juizes AS id, juizes.* FROM juizes WHERE id_juizes = ? ORDER BY id_juizes LIMIT 1"
		if err := dbOld.Raw(queryJ, inscricao.IDJuiz).Scan(&juiz).Error; err != nil {
			log.Printf("Erro ao buscar juiz OLD com ID %d: %v", inscricao.IDJuiz, err)
			continue
		}

		//Buscar Judge pelo email no DB Novo
		var judge judge.Judge
		if err := dbNew.Where("email = ?", juiz.Email).First(&judge).Error; err != nil {
			log.Printf("Erro ao buscar juiz com ID %d: %v", juiz.Email, err)
			continue // Ou trate o erro conforme necessário
		}

		//Buscar Classe pelo ID DB Antigo
		var classe catshowclass.OldClass
		if err := dbOld.Where("id_classes = ?", inscricao.IDClasse).First(&classe).Error; err != nil {
			log.Printf("Erro ao buscar classe com ID %d: %v", inscricao.IDClasse, err)
			continue // Ou trate o erro conforme necessário
		}
		//Buscar Classe pelo Descricao DB Novo
		var class catshowclass.Class
		if err := dbNew.Where("description = ?", classe.DescricaoClasse).First(&class).Error; err != nil {
			log.Printf("Erro ao buscar classe com ID %d: %v", classe.Classe, err)
			continue // Ou trate o erro conforme necessário
		}

		// Montar o objeto Registration com as informações necessárias
		registration := Registration{
			CatShowID:        &catShow.ID,
			CatShowSubID:     &catShowSub.ID,
			OwnerID:          &owner.ID,
			CatID:            &catData.ID,
			ClassID:          &class.ID,
			JudgeID:          &judge.ID,
			RegistrationDate: inscricao.DataCadastro,
			Number:           inscricao.Numero,
			Observations:     string(inscricao.Observacao),
			Updated:          false,
			Active:           true,
			CatIDOld:         catData.CatIDOld,
		}

		logger := logrus.New() // Cria uma nova instância de logger
		catShowRegRepo := NewCatShowRegistrationRepository(dbNew, logger)
		//catShowRepo := catshow.NewCatShowRepository(dbNew, logger)
		//catShowService := catshow.NewCatShowService(catShowRepo, logger)
		catRepo := cat.NewCatRepository(dbNew, logger)
		FileCatRepo := cat.NewFilesCatRepository(dbNew, logger)
		fileService := utils.NewFilesService(&s3.S3{}, logger)
		catFileService := cat.NewCatFileService(fileService, FileCatRepo, logger)
		catService := cat.NewCatService(catRepo, catFileService, logger)
		catShowcatRepo := catshowcat.NewCatShowCatRepository(dbNew, logger)
		filesCatShowCatRepo := catshowcat.NewFilesCatShowCatRepository(dbNew, logger)
		catFileServiceShow := catshowcat.NewCatShowCatFileService(fileService, filesCatShowCatRepo, logger)
		catShowCatService := catshowcat.NewCCatShowtService(catShowcatRepo, catFileServiceShow, logger)

		catShowRegService := NewCatShowRegistrationService(logger, catShowCatService, catService, catShowRegRepo)

		var fileDescription []utils.FileWithDescription
		log.Printf("Criando registro para o gato: CatShowID=%v, CatShowSubID=%v, OwnerID=%v, CatID=%v, ClassID=%v, JudgeID=%v, RegistrationDate=%v, Number=%d, Observations=%s, Updated=%t, Active=%t",
			*registration.CatShowID, *registration.CatShowSubID, *registration.OwnerID, *registration.CatID, *registration.ClassID, *registration.JudgeID,
			registration.RegistrationDate, registration.Number, registration.Observations, registration.Updated, registration.Active)

		catShowRegService.CreateCatShowRegistration(&registration, fileDescription)

		log.Printf("Processamento da inscrição ID: %d concluído com sucesso", inscricao.IDInscricoes)

	}

	log.Printf("Fim Migração Inscricoes")
	return nil
}

func MigrateInscricoesUpdate(dbOld, dbNew *gorm.DB) error {
	var inscricoesUpdate []InscricaoUpdate
	log.Println("Inicio Migração Inscricoes Update")

	// Buscar todas as inscrições
	if err := dbOld.Find(&inscricoesUpdate).Error; err != nil {
		log.Printf("Erro ao buscar inscrições: %v", err)
		return err
	}

	// Logando a quantidade de inscrições recuperadas
	log.Printf("Total de inscrições encontradas: %d", len(inscricoesUpdate))

	for _, inscricaoUpdate := range inscricoesUpdate {
		log.Printf("Iniciando processamento da inscriçãoUpdate")

		//Busca  Exposicao pelo ID DB Old
		var exposicao catshow.Exposicao
		if err := dbOld.Where("id_exposicoes = ?", inscricaoUpdate.IDExposicao).First(&exposicao).Error; err != nil {
			log.Printf("Erro ao buscar exposicao com ID %d: %v", inscricaoUpdate.IDExposicao, err)
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
		if err := dbOld.Where("id_exposicoes_sub = ?", inscricaoUpdate.IDExposicaoSub).First(&exposicaoSub).Error; err != nil {
			log.Printf("Erro ao buscar exposicao sub com ID %d: %v", inscricaoUpdate.IDExposicaoSub, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca ID do Show Sub pela descricao no DB Novo
		var catShowSub catshow.CatShowSub
		if err := dbNew.Where("description = ?", exposicaoSub.DescricaoExpo).First(&catShowSub).Error; err != nil {
			log.Printf("Erro ao buscar exposicao Sub  com ID %d: %v", exposicaoSub.IDExposicao, err)
			continue // Ou trate o erro conforme necessário
		}

		//Busca Juiz pelo ID no DB Antigo
		var juiz judge.JudgesS
		// Ajustando a consulta para incluir um campo 'id' através de um alias para 'id_juizes'
		queryJ := "SELECT id_juizes AS id, juizes.* FROM juizes WHERE id_juizes = ? ORDER BY id_juizes LIMIT 1"
		if err := dbOld.Raw(queryJ, inscricaoUpdate.IDJuiz).Scan(&juiz).Error; err != nil {
			log.Printf("Erro ao buscar juiz OLD com ID %d: %v", inscricaoUpdate.IDJuiz, err)
			continue
		}

		//Buscar Judge pelo email no DB Novo
		var judge judge.Judge
		if err := dbNew.Where("email = ?", juiz.Email).First(&judge).Error; err != nil {
			log.Printf("Erro ao buscar juiz com ID %d: %v", juiz.Email, err)
			continue // Ou trate o erro conforme necessário
		}

		//Buscar Classe pelo ID DB Antigo
		var classe catshowclass.OldClass
		if err := dbOld.Where("id_classes = ?", inscricaoUpdate.IDClasse).First(&classe).Error; err != nil {
			log.Printf("Erro ao buscar classe com ID %d: %v", inscricaoUpdate.IDClasse, err)
			continue // Ou trate o erro conforme necessário
		}
		//Buscar Classe pelo Descricao DB Novo
		var class catshowclass.Class
		if err := dbNew.Where("description = ?", classe.DescricaoClasse).First(&class).Error; err != nil {
			log.Printf("Erro ao buscar classe com ID %d: %v", classe.Classe, err)
			continue // Ou trate o erro conforme necessário
		}

		// Buscar Cor pelo ID no DB Antigo
		var cor color.ColorS
		// Ajustando a consulta para incluir um campo 'id' através de um alias para 'id_cores'
		queryC := "SELECT id_cores AS id, cores.* FROM cores WHERE id_cores = ? ORDER BY id_cores LIMIT 1"
		if err := dbOld.Raw(queryC, inscricaoUpdate.IDCor).Scan(&cor).Error; err != nil {
			log.Printf("Erro ao buscar Cor com ID %d: %v", inscricaoUpdate.IDCor, err)
			continue // Ou trate o erro conforme necessário
		}

		//Buscar Color pelo Descricao DB Novo
		var color catshowclass.Class
		if err := dbNew.Table("colors").Where("ems_code = ? AND breed_code = ?", cor.EmsCode, cor.BreedID).First(&color).Error; err != nil {
			log.Printf("Erro ao buscar Color com ID %d: %v", cor.EmsCode, err)
			continue // Ou trate o erro conforme necessário
		}
		// Buscar o registro de inscrição correspondente
		var registration Registration
		if err := dbNew.Where("cat_id_old = ? AND cat_show_id = ? AND cat_show_sub_id = ?", inscricaoUpdate.IDGato, catShow.ID, catShowSub.ID).First(&registration).Error; err != nil {
			log.Printf("Erro ao buscar gato com ID OLD %d, CatShowID %d, CatShowSubID %d: %v", inscricaoUpdate.IDGato, catShow.ID, catShowSub.ID, err)
			continue // Ou trate o erro conforme necessário
		}

		sexString := ""
		if inscricaoUpdate.Sexo == "1" {
			sexString = "male"
		} else if inscricaoUpdate.Sexo == "2" {
			sexString = "female"
		}

		// Montar o objeto Registration com as informações necessárias
		registrationUpdate := RegistrationUpdated{
			RegistrationID: &registration.ID,
			ClassID:        &class.ID,
			ColorID:        &color.ID,
			JudgeID:        &judge.ID,
			Birthdate:      inscricaoUpdate.Nascimento,
			Gender:         sexString,
		}

		if err := dbNew.Create(&registrationUpdate).Error; err != nil {
			log.Printf("Erro ao criar registro de atualização para a inscrição ID: %d: %v", registration.ID, err)
			// Considerar a inclusão de uma instrução 'continue' ou tratamento de erro aqui
		} else {
			log.Println("Registro criado com sucesso.")
		}
		// var existingRegistration RegistrationUpdated
		// err := dbNew.Where(
		// 	"class_id = ? AND color_id = ? AND judge_id = ? AND birthdate = ? AND gender = ?",
		// 	class.ID, color.ID, judge.ID, inscricaoUpdate.Nascimento, sexString,
		// ).First(&existingRegistration).Error

		// if err != nil {
		// 	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 		// Registro não encontrado, criar novo
		// 		if err := dbNew.Create(&registrationUpdate).Error; err != nil {
		// 			log.Printf("Erro ao criar registro de atualização para a inscrição ID: %d: %v", registration.ID, err)
		// 			// Considerar a inclusão de uma instrução 'continue' ou tratamento de erro aqui
		// 		} else {
		// 			log.Println("Registro criado com sucesso.")
		// 		}
		// 	} else {
		// 		// Tratar outros erros que não o de "não encontrado"
		// 		log.Printf("Erro ao buscar registro existente: %v", err)
		// 	}
		// } else {
		// 	log.Println("Um registro com as características especificadas já existe.")
		// 	// Aqui você pode decidir atualizar o registro existente ou simplesmente pular a criação
		// }

		log.Printf("Iniciando processamento da inscriçãoUpdate OK")

	}

	log.Printf("Fim Migração Inscricoes")
	return nil
}
