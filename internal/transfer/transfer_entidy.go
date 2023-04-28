package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/service"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransferMongo struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	CatData        service.CatService   `bson:"catData"`
	SellerData     service.OwnerService `bson:"sellerData"`
	BuyerData      service.OwnerService `bson:"buyerData"`
	Status         string               `bson:"status"`
	ProtocolNumber string               `bson:"protocolNumber"`
	RequesterID    primitive.ObjectID   `bson:"requesterID"`
	Files          []utils.Files        `bson:"files"`
}
