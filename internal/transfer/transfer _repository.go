package transfer

import (
	"context"

	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransferRepository struct {
	DB     *mongo.Client
	Logger *logrus.Logger
}

func NewTransferRepository(db *mongo.Client, logger *logrus.Logger) *TransferRepository {
	return &TransferRepository{
		DB:     db,
		Logger: logger,
	}
}

var database = "amacoon"
var collection = "transfers"

func (r *TransferRepository) GetTransferByID(id string) (TransferMongo, error) {
	r.Logger.Infof("Repository GetTransferByID")
	var transfer TransferMongo
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return TransferMongo{}, err
	}
	filter := bson.M{"_id": objID}
	err = r.DB.Database(database).Collection(collection).FindOne(context.Background(), filter).Decode(&transfer)
	if err != nil {
		return TransferMongo{}, err
	}

	r.Logger.Infof("Repository GetTransferByID OK")
	return transfer, nil
}

func (r *TransferRepository) CreateTransfer(transfer TransferMongo) (TransferMongo, error) {
	r.Logger.Infof("Repository CreateTransfer")

	res, err := r.DB.Database(database).Collection(collection).InsertOne(context.Background(), transfer)
	if err != nil {
		return TransferMongo{}, err
	}

	transfer.ID = res.InsertedID.(primitive.ObjectID)
	r.Logger.Infof("Repository CreateTransfer OK")
	return transfer, nil
}

func (r *TransferRepository) UpdateTransferStatus(id string, status string) error {
	r.Logger.Infof("Repository UpdateTransferStatus")

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

	r.Logger.Infof("Repository UpdateTransferStatus OK")
	return nil
}

func (r *TransferRepository) AddTransferFiles(id string, files []utils.Files) error {
	r.Logger.Infof("Repository AddTransferFiles id %s", id)
	
	r.Logger.Infof("Repository AddTransferFiles OK")
	return nil
}

func (r *TransferRepository) GetAllTransfersByRequesterID(requesterID string) ([]TransferMongo, error) {
	r.Logger.Infof("Repository GetAllTransfersByRequesterID")
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
	var transfers []TransferMongo
	for cur.Next(context.Background()) {
		var transfer TransferMongo
		err := cur.Decode(&transfer)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetAllTransfersByRequesterID OK")
	return transfers, nil
}

func (r *TransferRepository) GetTransferFilesByID(id string) ([]utils.Files, error) {
	r.Logger.Infof("Repository GetTransferFilesByID")

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
	var transfer TransferMongo
	if err := result.Decode(&transfer); err != nil {
		return nil, err
	}
	r.Logger.Infof("Repository GetTransferFilesByID OK")
	return transfer.Files, nil
}

func (r *TransferRepository) UpdateTransfer(id string, transfer TransferMongo) error {
	r.Logger.Infof("Repository UpdateTransfer")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.Logger.Errorf("error parsing id to ObjectID: %v", err)
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": transfer}
	_, err = r.DB.Database(database).Collection(collection).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	r.Logger.Infof("Repository UpdateTransfer OK")
	return nil
}
