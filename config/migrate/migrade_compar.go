package migrate

import (
	"context"
	"encoding/csv"
	"fmt"

	models "github.com/scuba13/AmacoonServices/internal/cat"
	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"os"

	"gorm.io/gorm"
	//"strings"
)

func GenerateCSV(db *gorm.DB, client *mongo.Client) error {
	// Abrir o arquivo para escrita
	file, err := os.Create("cats_parents.csv")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Criar um writer CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escrever o cabeçalho
	writer.Write([]string{"Registro MySQL", "Nome do Pai My SQL", "Nome da Mãe MySQL", "Registration Mongo", "FatherName Mongo", "MotherName Mongo", "Comparação"})

	// Query all rows from CatMigration table
	var catMigrations []sql.CatMigration
	if err := db.Unscoped().Find(&catMigrations).Error; err != nil {
		return fmt.Errorf("error querying CatMigration table: %v", err)
	}

	// Para cada linha, buscar os dados dos pais na collection cats e escrever no arquivo
	for _, cat := range catMigrations {
		// Buscar os dados do gato na collection cats pelo registro
		var catMongo models.CatMongo
		filter := bson.M{"registration": cat.Registry}
		if err := client.Database("amacoon").Collection("cats").FindOne(context.Background(), filter).Decode(&catMongo); err != nil {
			if err == mongo.ErrNoDocuments {
				continue // Se o gato não existe, pular para o próximo
			}
			return fmt.Errorf("error querying cats collection: %v", err)
		}

		// Buscar os dados do pai pelo seu ID na collection cats
		var father models.CatMongo
		if !catMongo.FatherID.IsZero() {
			if err := client.Database("amacoon").Collection("cats").FindOne(context.Background(), bson.M{"_id": catMongo.FatherID}).Decode(&father); err != nil {
				if err == mongo.ErrNoDocuments {
					father.Name = ""
				} else {
					return fmt.Errorf("error querying cats collection: %v", err)
				}
			}
		} else {
			father.Name = ""
		}

		// Buscar os dados da mãe pelo seu ID na collection cats
		var mother models.CatMongo
		if !catMongo.MotherID.IsZero() {
			if err := client.Database("amacoon").Collection("cats").FindOne(context.Background(), bson.M{"_id": catMongo.MotherID}).Decode(&mother); err != nil {
				if err == mongo.ErrNoDocuments {
					mother.Name = ""
				} else {
					return fmt.Errorf("error querying cats collection: %v", err)
				}
			}
		} else {
			mother.Name = ""
		}
				// // Escrever os dados no arquivo CSV e comparar os nomes dos pais com a collection cats
				// var comparacao string
				// if strings.EqualFold(father.Name, models.CatMongo.F) && strings.EqualFold(mother.Name, catMongo.MotherName) {
				// 	comparacao = "ok"
				// } else {
				// 	comparacao = "ko"
				// }
				writer.Write([]string{
					cat.Registry,
					cat.FatherName,
					cat.MotherName,
					catMongo.Registration,
					father.Name,
					mother.Name,
					//comparacao, // Adicionar coluna de comparação
				})
				
			}
			return nil
		}
		