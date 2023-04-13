package migrate

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	mongomodels "github.com/scuba13/AmacoonServices/config/migrate/models/mongo"
	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"unicode"
)

func cleanParentName(name string) string {
	// Remover a vírgula e o asterisco
	name = strings.ReplaceAll(name, ",", "")
	name = strings.ReplaceAll(name, "*", "")

	// Converter a string para minúsculas
	name = strings.ToLower(name)

	// Lista das palavras a serem removidas
	wordsToRemove := []string{"ch", "pr", "ic", "ip", "gic", "gip", "sc", "sp", "nw", "aw", "bw", "cew", "mw", "nsw", "sw", "ww", "jw", "dsw", "dm", "dsm", "dvm"}

	// Cria uma expressão regular a partir das palavras a serem removidas
	regexPattern := `(\b(?:` + strings.Join(wordsToRemove, "|") + `)\b\s*)`
	regex := regexp.MustCompile(regexPattern)

	// Remove as palavras do nome
	cleanedName := regex.ReplaceAllString(name, "")

	// Se os dois primeiros caracteres são letras e o terceiro é um espaço, remova-os
	if len(cleanedName) >= 3 && unicode.IsLetter(rune(cleanedName[0])) && unicode.IsLetter(rune(cleanedName[1])) && unicode.IsSpace(rune(cleanedName[2])) {
		cleanedName = cleanedName[3:]
	}

	// Se as duas primeiras letras forem "br" ou "gb", remova-as
	if len(cleanedName) >= 2 && (cleanedName[:2] == "br" || cleanedName[:2] == "gb") {
		cleanedName = cleanedName[2:]
	}

	// Remove um ponto no início da string, se houver
	if len(cleanedName) > 0 && cleanedName[0] == '.' {
		cleanedName = cleanedName[1:]
	}

	// Se o primeiro caractere for um apóstrofo, remova-o e todos os números seguintes até encontrar uma letra
	if len(cleanedName) > 0 && cleanedName[0] == '\'' {
		cleanedName = cleanedName[1:]
		if len(cleanedName) > 0 {
			cleanedName = strings.TrimLeftFunc(cleanedName, unicode.IsDigit)
		}
	}
	// Se os dois primeiros caracteres são números, remova-os
	if len(cleanedName) >= 2 && unicode.IsDigit(rune(cleanedName[0])) && unicode.IsDigit(rune(cleanedName[1])) {
		cleanedName = cleanedName[2:]
	}

	// Remover espaços no início e no fim
	cleanedName = strings.TrimSpace(cleanedName)

	// Remover espaços extras entre as palavras
	space := regexp.MustCompile(`\s+`)
	cleanedName = space.ReplaceAllString(cleanedName, " ")

	return cleanedName
}

func MigrateCatsPattentNameToCats1(db *gorm.DB, mongoClient *mongo.Client) error {
	// Fetch CatMigration records from SQL
	var catMigrationRecords []sql.CatMigration
	if err := db.Unscoped().Find(&catMigrationRecords).Error; err != nil {
		return fmt.Errorf("failed to fetch CatMigration records: %w", err)
	}
	// Prepare data for insertion into MongoDB
	var catMongoRecords []interface{}
	for _, cat := range catMigrationRecords {
		catMongoRecords = append(catMongoRecords, mongomodels.CatTemp{
			CatID:      cat.ID,
			Registry:   cat.Registry,
			CatName:    cat.CatName,
			FatherName: cleanParentName(cat.FatherName),
			MotherName: cleanParentName(cat.MotherName),
		})
	}

	// Insert records into MongoDB
	mongoDB := mongoClient.Database("amacoon")
	catsCollection := mongoDB.Collection("cats_temp")

	// // Create a unique index on id_gatos to avoid duplicate entries
	// indexModel := mongo.IndexModel{
	// 	Keys: bson.M{
	// 		"id_gatos": 1,
	// 	},
	// 	Options: options.Index().SetUnique(true),
	// }
	// if _, err := catsCollection.Indexes().CreateOne(context.Background(), indexModel); err != nil {
	// 	return fmt.Errorf("failed to create index: %w", err)
	// }

	// Bulk insert data into MongoDB
	_, err := catsCollection.InsertMany(context.Background(), catMongoRecords)
	if err != nil {
		return fmt.Errorf("failed to insert records into MongoDB: %w", err)
	}

	return nil
}
