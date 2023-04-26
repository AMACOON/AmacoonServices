package cat

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatRepoInterface interface {
	
	GetCatCompleteByID(id string) (*CatComplete, error)
	GetAllByOwnerAndGender(ownerID, gender string) ([]*CatComplete, error)
	GetCatCompleteByRegistration(id string) (*CatComplete, error)
	GetAllByOwner(ownerID string) ([]*CatComplete, error)
	
}

type CatRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewCatRepository(db *mongo.Client, logger *logrus.Logger) CatRepoInterface {
	return &CatRepository{
		DB:     db,
		Logger: logger,
	}
}
