package cat

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/breed"
	"github.com/scuba13/AmacoonServices/internal/cattery"
	"github.com/scuba13/AmacoonServices/internal/color"
	"github.com/scuba13/AmacoonServices/internal/country"
	"github.com/scuba13/AmacoonServices/internal/federation"
	"github.com/scuba13/AmacoonServices/internal/owner"
	"github.com/scuba13/AmacoonServices/internal/title"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatComplete struct {
	ID               primitive.ObjectID         `bson:"_id"`
	Name             string                     `bson:"name"`
	FullName         string                     `bson:"FullName"`
	Registration     string                     `bson:"registration"`
	RegistrationType string                     `bson:"registrationType"`
	Microchip        string                     `bson:"microchip"`
	Gender           string                     `bson:"gender"`
	Birthdate        time.Time                  `bson:"birthdate"`
	Neutered         bool                       `bson:"neutered"`
	Validated        bool                       `bson:"validated"`
	Observation      string                     `bson:"observation"`
	Fifecat          bool                       `bson:"fifecat"`
	FatherName       string                     `bson:"fatherName"`
	MotherName       string                     `bson:"motherName"`
	Titles           []TitlesCats               `bson:"titles"`
	Country          country.CountryMongo       `bson:"country"`
	Federation       federation.FederationMongo `bson:"federation"`
	Breed            breed.BreedMongo           `bson:"breed"`
	Color            color.ColorMongo           `bson:"color"`
	Cattery          cattery.CatteryMongo       `bson:"cattery"`
	Owner            owner.OwnerMongo           `bson:"owner"`
	Files            []utils.Files              `bson:"files"`
}

type TitlesCats struct {
    Title          title.TitlesMongo `bson:"title"`
    Date           time.Time `bson:"date"`
    FederationName string    `bson:"federationName"`
}
