package litter

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LitterRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewLitterRepository(db *mongo.Client, logger *logrus.Logger) *LitterRepository {
	return &LitterRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "litters"

func (r *LitterRepository) GetLitterByID(id string) (Litter, error) {
	r.Logger.Infof("Repository GetLitterByID")
	var litter Litter
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Litter{}, err
	}
	filter := bson.M{"_id": objID}
	err = r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&litter)
	if err != nil {
		return Litter{}, err
	}

	r.Logger.Infof("Repository GetLitterByID OK")
	return litter, nil
}

func (r *LitterRepository) CreateLitter(litter Litter) (Litter, error) {
    r.Logger.Infof("Repository CreateLitter")
    litter.RequesterID = primitive.NewObjectID()
    litter.Status = "submitted"
    _, err := r.DB.Database(database).Collection(collection).InsertOne(context.Background(), litter)
    if err != nil {
        return Litter{}, err
    }
    r.Logger.Infof("Repository CreateLitter OK")
    return litter, nil
}

func (r *LitterRepository) UpdateLitterStatus(id string, status string) error {
	r.Logger.Infof("Repository UpdateLitterStatus")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"status": status}}
	_, err = r.DB.Database(database).Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	r.Logger.Infof("Repository UpdateLitterStatus OK")
	return nil
}

func (r *LitterRepository) AddLitterFiles(id string, files []utils.Files) error {
	r.Logger.Infof("Repository AddLitterFiles id %s", id)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"files": files}}
	_, err = r.DB.Database(database).Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	r.Logger.Infof("Repository AddLitterFiles OK")
	return nil
}

func (r *LitterRepository) GetAllLittersByOwner(ownerID string) ([]Litter, error) {
	r.Logger.Infof("Repository GetAllLittersByOwner")
	objID, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"requesterID": objID}
	cur, err := r.DB.Database(database).Collection(collection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var litters []Litter
	for cur.Next(context.Background()) {
		var litter Litter
		err := cur.Decode(&litter)
		if err != nil {
			return nil, err
		}
		litters = append(litters, litter)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetAllLittersByOwner OK")
	return litters, nil
}

func (r *LitterRepository) GetLitterFilesByID(id string) ([]utils.Files, error) {
    r.Logger.Infof("Repository GetLitterFilesByID")
    
	objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        r.Logger.Errorf("error parsing id to ObjectID: %v", err)
        return nil, err
    }
    filter := bson.M{"_id": objID}
    result := r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter)
    if err := result.Err(); err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    var litter Litter
    if err := result.Decode(&litter); err != nil {
        return nil, err
    }
    r.Logger.Infof("Repository GetLitterFilesByID OK")
    return litter.Files, nil
}

func (r *LitterRepository) UpdateLitter(litter Litter) error {
    r.Logger.Infof("Repository UpdateLitter")
    filter := bson.M{"_id": litter.MotherData.ID}
    update := bson.M{"$set": bson.M{
        "motherData":     litter.MotherData,
        "fatherData":     litter.FatherData,
        "birthData":      litter.BirthData,
        "kittenData":     litter.KittenData,
        "status":         litter.Status,
        "protocolNumber": litter.ProtocolNumber,
        "requesterID":    litter.RequesterID,
        "files":          litter.Files,
    }}
    _, err := r.DB.Database(database).Collection(collection).UpdateOne(context.Background(), filter, update)
    if err != nil {
        return err
    }
    r.Logger.Infof("Repository UpdateLitter OK")
    return nil
}


