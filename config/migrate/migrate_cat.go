package migrate

import (
	"context"
	"log"

	models "github.com/scuba13/AmacoonServices/internal/cat"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"github.com/scuba13/AmacoonServices/config/migrate/models/sql"
	"go.mongodb.org/mongo-driver/bson"
)
const batchSize = 100


func MigrateCats(db *gorm.DB, mongoClient *mongo.Client) error {
	var catMigrations []sql.CatTable
	var err error
	var offset int

	catCollection := mongoClient.Database("amacoon").Collection("catteries")
	for {
		catMigrations, err = GetCatsMigrate(db, offset, batchSize)
		if err != nil {
			return err
		}

		if len(catMigrations) == 0 {
			break
		}
		

		log.Printf("Migrating cats %d - %d\n", offset, offset+len(catMigrations)-1)

		catMongos := make([]interface{}, len(catMigrations))
		for i, cat := range catMigrations {
			log.Printf("Migrating cat %d: %s\n", i, cat.Name)
		
			filter := bson.M{"name": cat.Name}
			count, err := catCollection.CountDocuments(context.Background(), filter)
			if err != nil {
				return err
			}
		
			if count > 0 {
				log.Printf("Cat already exists: %s", cat.Name)
				continue
			}
		
			neutered := cat.Neutered == "s"
			validated := cat.Validated == "s"
			fifecat := cat.FifeCat == "s"
		
			titles := []string{}
			if cat.AdultTitle != "0" {
				titles = append(titles, cat.AdultTitle)
			}
		
			if cat.NeuterTitle != "0" {
				titles = append(titles, cat.NeuterTitle)
			}
		
			if cat.WW == "1" {
				titles = append(titles, "WW")
			}
			if cat.SW == "1" {
				titles = append(titles, "SW")
			}
			if cat.NW == "1" {
				titles = append(titles, "NW")
			}
			if cat.JW == "1" {
				titles = append(titles, "JW")
			}
			if cat.DVM == "1" {
				titles = append(titles, "DVM")
			}
			if cat.DSM == "1" {
				titles = append(titles, "DSM")
			}
			if cat.DM == "1" {
				titles = append(titles, "DM")
			}
		
			federationID, err := getFederationID(mongoClient, cat.FedName)
			if err != nil {
				log.Printf("Error getting federation ID for cat %d: %v\n", i, err)
				return err
			}
		
			breedID, err := getBreedID(mongoClient, cat.BreedName)
			if err != nil {
				log.Printf("Error getting getBreedID for cat %d: %v\n", i, err)
				return err
			}
		
			colorID, err := getColorID(mongoClient, cat.EmsCode, cat.BreedID)
			if err != nil {
				log.Printf("Error getting getColorID for cat %d: %v\n", i, err)
				return err
			}
		
			countryId, err := findCountryIdByCode(mongoClient, cat.Country)
			if err != nil {
				log.Printf("Error getting findCountryIdByCode for cat %d: %v\n", i, err)
				return err
			}
		
			catteryId, err := getCatteryID(mongoClient, cat.BreederName)
			if err != nil {
				log.Printf("Error getting getCatteryID for cat %d: %v\n", i, err)
				return err
			}
			
			ownerId, err := getOwnerID(mongoClient, cat.OwnerName)
			if err != nil {
				log.Printf("Error getting getOwnerID for cat %d: %v\n", i, err)
				return err
			}
	
			sexString := ""
			if cat.Sex == "1" {
				sexString = "male"
			} else if cat.Sex == "2" {
				sexString = "female"
			}
	
		
	
			catMongos[i] = models.CatMongo{
				ID:                       primitive.NewObjectID(),
				Name:                     cat.Name,
				Registration:             cat.Registration,
				RegistrationType:         cat.RegType,
				Microchip:                cat.Microchip,
				Sex:                      sexString,
				Birthdate:                cat.BirthDate,
				Neutered:                 neutered,
				Validated:                validated,
				Observation:              "",
				Fifecat:                  fifecat,
				Titles:                   titles,
				FederationID: federationID,
				BreedID:                  breedID,
				ColorID:                  colorID,
				FatherID:                 primitive.NilObjectID,
				MotherID:                 primitive.NilObjectID,
				CatteryID:                catteryId,
				OwnerID:                  ownerId,
				CountryID:                countryId,
			}
	
			// Insert cat into database
			_, err = catCollection.InsertOne(context.Background(), catMongos[i])
			if err != nil {
				log.Printf("Error inserting cat %s: %v\n", cat.Name, err)
				return err
			}
		}
	
		offset += len(catMigrations)
	}
	
	log.Printf("Migrated all cats\n")
	return nil
	
}	




func GetCatsMigrate(db *gorm.DB, offset, batchSize int) ([]sql.CatTable, error) {
	var cats []sql.CatTable

	query := db.Unscoped().Joins("LEFT JOIN racas ON gatos.id_raca = racas.id_racas").
        Joins("LEFT JOIN cores ON gatos.id_cor = cores.id_cores").
        Joins("LEFT JOIN gatis ON gatos.id_gatil = gatis.id_gatis").
        Joins("LEFT JOIN expositores ON gatos.id_expositor= expositores.id_expositores").
        Joins("LEFT JOIN federacoes ON gatos.registro_federacao= federacoes.id_federacoes").
        Select("gatos.*, racas.nome AS nome_raca, cores.id_emscode AS id_emscode, cores.descricao AS nome_cor, gatis.nome_gatil AS nome_gatil , expositores.nome AS nome_expositor, federacoes.descricao AS nome_federacao").
        Limit(batchSize).Offset(offset).
        Find(&cats)

	if err := query.Error; err != nil {
		return nil, err
	}

	return cats, nil
}