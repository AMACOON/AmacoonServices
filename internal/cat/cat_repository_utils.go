package cat

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var database = "amacoon"
var catsCollection = "cats"
var breedsCollection = "breeds"
var colorsCollection = "colors"
var federationsCollection = "federations"
var catteriesCollection = "catteries"
var ownersCollection = "owners"
var countriesCollection = "countries"


func (r *CatRepository) getFatherName(cat *CatMongo) (string, error) {
	catCollection := r.DB.Database(database).Collection(catsCollection)

	father := &CatMongo{}
	if err := catCollection.FindOne(context.Background(), bson.M{"_id": cat.FatherID}).Decode(father); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		r.Logger.WithError(err).Errorf("error getting father: %v", err)
		return "", err
	}
	return father.Name, nil
}

func (r *CatRepository) getCountry(cat *CatMongo) (*country.CountryMongo, error) {
	country := &country.CountryMongo{}
	if err := r.DB.Database(database).Collection(countriesCollection).FindOne(context.Background(), bson.M{"_id": cat.CountryID}).Decode(country); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.Logger.WithError(err).Errorf("error getting country: %v", err)
		return nil, err
	}
	return country, nil
}

func (r *CatRepository) getMotherName(cat *CatMongo) (string, error) {
	catCollection := r.DB.Database(database).Collection(catsCollection)

	mother := &CatMongo{}
	if err := catCollection.FindOne(context.Background(), bson.M{"_id": cat.MotherID}).Decode(mother); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		r.Logger.WithError(err).Errorf("error getting mother: %v", err)
		return "", err
	}
	return mother.Name, nil
}

func (r *CatRepository) getBreed(cat *CatMongo) (*breed.BreedMongo, error) {
	breed := &breed.BreedMongo{}
	if err := r.DB.Database(database).Collection(breedsCollection).FindOne(context.Background(), bson.M{"_id": cat.BreedID}).Decode(breed); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.Logger.WithError(err).Errorf("error getting breed: %v", err)
		return nil, err
	}
	return breed, nil
}

func (r *CatRepository) getColor(cat *CatMongo) (*color.ColorMongo, error) {
	color := &color.ColorMongo{}
	if err := r.DB.Database(database).Collection(colorsCollection).FindOne(context.Background(), bson.M{"_id": cat.ColorID}).Decode(color); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.Logger.WithError(err).Errorf("error getting color: %v", err)
		return nil, err
	}
	return color, nil
}

func (r *CatRepository) getCattery(cat *CatMongo) (*cattery.CatteryMongo, error) {
	cattery := &cattery.CatteryMongo{}
	if err := r.DB.Database(database).Collection(catteriesCollection).FindOne(context.Background(), bson.M{"_id": cat.CatteryID}).Decode(cattery); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.Logger.WithError(err).Errorf("error getting cattery: %v", err)
		return nil, err
	}
	return cattery, nil
}

// Busca o proprietário do gato
func (r *CatRepository) getOwner(cat *CatMongo) (*owner.OwnerMongo, error) {
	owner := &owner.OwnerMongo{}
	if err := r.DB.Database(database).Collection(ownersCollection).FindOne(context.Background(), bson.M{"_id": cat.OwnerID}).Decode(owner); err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithField("id", cat.OwnerID).Warn("Owner not found")
			owner = nil
		} else {
			r.Logger.WithError(err).Errorf("error getting owner: %v", err)
			return nil, err
		}
	}
	return owner, nil
}



// Busca o pai do gato
func (r *CatRepository) getFather(cat *CatMongo) (string, error) {
	father := &CatMongo{}
	if err := r.DB.Database(database).Collection(catsCollection).FindOne(context.Background(), bson.M{"_id": cat.FatherID}).Decode(father); err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithField("id", cat.FatherID).Warn("Father not found")
			return "", nil
		} else {
			r.Logger.WithError(err).Errorf("error getting father: %v", err)
			return "", err
		}
	}
	return father.Name, nil
}

// Busca a mãe do gato
func (r *CatRepository) getMother(cat *CatMongo) (string, error) {
	mother := &CatMongo{}
	if err := r.DB.Database(database).Collection(catsCollection).FindOne(context.Background(), bson.M{"_id": cat.MotherID}).Decode(mother); err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithField("id", cat.MotherID).Warn("Mother not found")
			return "", nil
		} else {
			r.Logger.WithError(err).Errorf("error getting mother: %v", err)
			return "", err
		}
	}
	return mother.Name, nil
}


func (r *CatRepository) getFederation(cat *CatMongo) (*federation.FederationMongo, error) {
	federation := &federation.FederationMongo{}
	r.Logger.Info("ID Federation: ", cat.RegistrationFederationID)
	if err := r.DB.Database(database).Collection(federationsCollection).FindOne(context.Background(), bson.M{"_id": cat.RegistrationFederationID}).Decode(federation); err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.Infof("Federation not found")
			return nil, nil
		}
		r.Logger.WithError(err).Errorf("error getting federation: %v", err)
		return nil, err
	}
	return federation, nil
}
