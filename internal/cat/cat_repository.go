package cat

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewCatRepository(db *mongo.Client, logger *logrus.Logger) *CatRepository {
	return &CatRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *CatRepository) GetCatCompleteByID(id string) (*CatComplete, error) {
	r.Logger.Infof("Repository GetCatCompleteByID")

	catCollection := r.DB.Database(database).Collection(catsCollection)

	cat := &CatMongo{}
	catID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.WithError(err).Errorf("invalid cat ID: %v", id)
		return nil, err
	}

	if err := catCollection.FindOne(context.Background(), bson.M{"_id": catID}).Decode(cat); err != nil {
		if err == mongo.ErrNoDocuments {
			r.Logger.WithField("id", id).Warn("Cat not found")
			return nil, err
		}
		r.Logger.WithError(err).Errorf("error getting cat: %v", err)
		return nil, err
	}

	catComplete := &CatComplete{
		ID:               cat.ID,
		Name:             cat.Name,
		Registration:     cat.Registration,
		RegistrationType: cat.RegistrationType,
		Microchip:        cat.Microchip,
		Sex:              cat.Sex,
		Birthdate:        cat.Birthdate,
		Neutered:         cat.Neutered,
		Validated:        cat.Validated,
		Observation:      cat.Observation,
		Fifecat:          cat.Fifecat,
		Titles:           cat.Titles,
	}

	if fatherName, err := r.getFather(cat); err != nil {
		return nil, err
	} else {
		catComplete.FatherName = fatherName
	}

	if country, err := r.getCountry(cat); err != nil {
		return nil, err
	} else {
		catComplete.Country = *country
	}

	if motherName, err := r.getMother(cat); err != nil {
		return nil, err
	} else {
		catComplete.MotherName = motherName
	}

	if breed, err := r.getBreed(cat); err != nil {
		return nil, err
	} else {
		catComplete.Breed = *breed
	}

	if color, err := r.getColor(cat); err != nil {
		return nil, err
	} else {
		catComplete.Color = *color
	}

	if cattery, err := r.getCattery(cat); err != nil {
		return nil, err
	} else {
		catComplete.Cattery = *cattery
	}

	if owner, err := r.getOwner(cat); err != nil {
		return nil, err
	} else {
		catComplete.Owner = *owner
	}

	if federation, err := r.getFederation(cat); err != nil {
		return nil, err
	} else {
		catComplete.Federation = *federation
	}

	return catComplete, nil
}
