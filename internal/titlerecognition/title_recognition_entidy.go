package titlerecognition

import (
	"time"

	"github.com/scuba13/AmacoonServices/internal/service"
	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TitleRecognitionMongo struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	CatData        service.CatService   `bson:"catData"`
	OwnerData      service.OwnerService `bson:"ownerData"`
	Titles         []TitleMongo         `bson:"titles"`
	Status         string               `bson:"status"`
	ProtocolNumber string               `bson:"protocolNumber"`
	RequesterID    primitive.ObjectID   `bson:"requesterID"`
	Files          []utils.Files        `bson:"files"`
}

type TitleMongo struct {
	TitleID     primitive.ObjectID `bson:"titleID"`
	TitleCode   string             `bson:"titleCode"`
	TitleName   string             `bson:"titleName"`
	Certificate string             `bson:"certificate"`
	Date        time.Time          `bson:"date"`
	Judge       string             `bson:"judge"`
}
