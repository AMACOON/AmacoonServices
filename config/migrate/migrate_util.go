package migrate

import (
	"context"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"gorm.io/gorm"
)

func populateCollection(db *gorm.DB, client *mongo.Client, collectionName string, items interface{}, convertItem func(item interface{}) interface{}, lenItems func(items interface{}) int, filterItem func(item interface{}) bson.M) error {
	collection := client.Database("amacoon").Collection(collectionName)

	batchSize := 500
	offset := 0

	for {
		if err := db.Unscoped().Limit(batchSize).Offset(offset).Find(items).Error; err != nil {
			return err
		}

		if lenItems(items) == 0 {
			break
		}

		valueItems := reflect.ValueOf(items).Elem()
		for i := 0; i < valueItems.Len(); i++ {
			item := valueItems.Index(i).Interface()
			filter := filterItem(item)
			count, err := collection.CountDocuments(context.Background(), filter)
			if err != nil {
				return err
			}

			if count == 0 {
				mongoItem := convertItem(item)
				_, err := collection.InsertOne(context.Background(), mongoItem)
				if err != nil {
					return err
				}
			}
		}

		if lenItems(items) < batchSize {
			break
		}

		offset += batchSize
	}

	return nil
}






func findCountryIdByCode(client *mongo.Client, countryCode string) (primitive.ObjectID, error) {
	if countryCode == "" || countryCode == "0" {
		return primitive.NilObjectID, nil
	}

	var country struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := client.Database("amacoon").Collection("countries").FindOne(context.Background(), bson.M{"code": countryCode}).Decode(&country)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return country.ID, nil
}

func getFederationID(client *mongo.Client, federationName string) (primitive.ObjectID, error) {

	if federationName == "" || federationName == "0" {
		return primitive.NilObjectID, nil
	}

	var federation struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := client.Database("amacoon").Collection("federations").FindOne(context.Background(), bson.M{"name": federationName}).Decode(&federation)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return federation.ID, nil
}

func getBreedID(client *mongo.Client, breedName string) (primitive.ObjectID, error) {
	if breedName == "" || breedName == "0" {
		return primitive.NilObjectID, nil
	}

	var breed struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := client.Database("amacoon").Collection("breeds").FindOne(context.Background(), bson.M{"breedName": breedName}).Decode(&breed)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return breed.ID, nil
}

func getColorID(client *mongo.Client, emsCode, breedCode string) (primitive.ObjectID, error) {
	if emsCode == "" || emsCode == "0" || breedCode == "" || breedCode == "0" {
		return primitive.NilObjectID, nil
	}

	var color struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	mongoDB := client.Database("amacoon")

	err := mongoDB.Collection("colors").FindOne(context.Background(), bson.M{"emsCode": emsCode, "breedCode": breedCode}).Decode(&color)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, errors.New("color not found")
		}
		return primitive.NilObjectID, err
	}

	return color.ID, nil
}

func getCatteryID(client *mongo.Client, breederName string) (primitive.ObjectID, error) {
	if breederName == "" || breederName == "0" {
		return primitive.NilObjectID, nil
	}
	var cattery struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	mongoDB := client.Database("amacoon")
	err := mongoDB.Collection("catteries").FindOne(context.Background(), bson.M{"name": breederName}).Decode(&cattery)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, errors.New("cattery not found")
		}
		return primitive.NilObjectID, err
	}

	return cattery.ID, nil
}

func getOwnerID(client *mongo.Client, ownerName string) (primitive.ObjectID, error) {

	if ownerName == "" || ownerName == "0" {
		return primitive.NilObjectID, nil
	}
	var owner struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	mongoDB := client.Database("amacoon")
	err := mongoDB.Collection("owners").FindOne(context.Background(), bson.M{"name": ownerName}).Decode(&owner)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, errors.New("owner not found")
		}
		return primitive.NilObjectID, err
	}

	return owner.ID, nil
}


