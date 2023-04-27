package titlerecognition

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TitleRecognitionRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewTitleRecognitionRepository(db *mongo.Client, logger *logrus.Logger) *TitleRecognitionRepository {
	return &TitleRecognitionRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "title_recognition"

func (r *TitleRecognitionRepository) GetTitleRecognitionByID(id string) (TitleRecognitionMongo, error) {
	r.Logger.Infof("Repository GetTitleRecognitionByID")
	var titleRecognition TitleRecognitionMongo
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}
	filter := bson.M{"_id": objID}
	err = r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&titleRecognition)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}

	r.Logger.Infof("Repository GetTitleRecognitionByID OK")
	return titleRecognition, nil
}

func (r *TitleRecognitionRepository) CreateTitleRecognition(titleRecognition TitleRecognitionMongo) (TitleRecognitionMongo, error) {
	r.Logger.Infof("Repository CreateTitleRecognition")

	res, err := r.DB.Database(database).Collection(collection).InsertOne(context.Background(), titleRecognition)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}

	titleRecognition.ID = res.InsertedID.(primitive.ObjectID)
	r.Logger.Infof("Repository CreateTitleRecognition OK")
	return titleRecognition, nil
}

func (r *TitleRecognitionRepository) UpdateTitleRecognitionStatus(id string, status string) error {
	r.Logger.Infof("Repository UpdateTitleRecognitionStatus")

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

	r.Logger.Infof("Repository UpdateTitleRecognitionStatus OK")
	return nil
}

func (r *TitleRecognitionRepository) AddTitleRecognitionFiles(id string, files []utils.Files) error {
	r.Logger.Infof("Repository AddTitleRecognitionFiles id %s", id)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Assign a new ObjectID to each file
	for i := range files {
		files[i].ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{"files": bson.M{"$each": files}}}
	_, err = r.DB.Database(database).Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	r.Logger.Infof("Repository AddTitleRecognitionFiles OK")
	return nil
}

func (r *TitleRecognitionRepository) GetAllTitleRecognitionByRequesterID(requesterID string) ([]TitleRecognitionMongo, error) {
	r.Logger.Infof("Repository GetAllTitlesByRequesterID")
	objID, err := primitive.ObjectIDFromHex(requesterID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"requesterID": objID}
	cur, err := r.DB.Database(database).Collection(collection).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var titles []TitleRecognitionMongo
	for cur.Next(context.Background()) {
		var title TitleRecognitionMongo
		err := cur.Decode(&title)
		if err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetAllTitlesByRequesterID OK")
	return titles, nil
}

func (r *TitleRecognitionRepository) GetTitleRecognitionFilesByID(id string) ([]utils.Files, error) {
	r.Logger.Infof("Repository GetTitleRecognitionFilesByID")

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
	var titleRecognition TitleRecognitionMongo
	if err := result.Decode(&titleRecognition); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetTitleRecognitionFilesByID OK")
	return titleRecognition.Files, nil
}

func (r *TitleRecognitionRepository) UpdateTitleRecognition(id string, titleRecognition TitleRecognitionMongo) error {
	r.Logger.Infof("Repository UpdateTitleRecognition")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.Errorf("error parsing id to ObjectID: %v", err)
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": titleRecognition}
	_, err = r.DB.Database(database).Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	r.Logger.Infof("Repository UpdateTitleRecognition OK")
	return nil
}



