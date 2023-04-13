package cat

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatComplete struct {
	ID               primitive.ObjectID
	Name             string
	Registration     string
	RegistrationType string
	Microchip        string
	Sex              string
	Birthdate        time.Time
	Neutered         bool
	Validated        bool
	Observation      string
	Fifecat          bool
	FatherName       string
	MotherName       string
	Titles           []string
	Country          country.CountryMongo
	Federation       federation.FederationMongo
	Breed            breed.BreedMongo
	Color            color.ColorMongo
	Cattery          cattery.CatteryMongo
	Owner            owner.OwnerMongo
}
