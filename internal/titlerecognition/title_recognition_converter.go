package titlerecognition

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (trr *TitleRecognitionRequest) ToTitleRecognitionMongo() (*TitleRecognitionMongo, error) {
	titles := make([]TitleMongo, len(trr.Titles))
	for i, title := range trr.Titles {
		titleMongo, err := title.ToTitleMongo()
		if err != nil {
			return nil, err
		}
		titles[i] = titleMongo
	}

	catService, err := trr.CatData.ToCatService()
	if err != nil {
		return nil, err
	}

	ownerService, err := trr.OwnerData.ToOwnerService()
	if err != nil {
		return nil, err
	}

	requesterID, err := primitive.ObjectIDFromHex(trr.RequesterID)
	if err != nil {
		return nil, err
	}

	return &TitleRecognitionMongo{
		CatData:        catService,
		OwnerData:      ownerService,
		Titles:         titles,
		Status:         trr.Status,
		ProtocolNumber: trr.ProtocolNumber,
		RequesterID:    requesterID,
	}, nil
}

func (tr *TitleRequest) ToTitleMongo() (TitleMongo, error) {
	titleID, err := primitive.ObjectIDFromHex(tr.TitleID)
	if err != nil {
		return TitleMongo{}, err
	}

	return TitleMongo{
		TitleID:     titleID,
		TitleCode:   tr.TitleCode,
		TitleName:   tr.TitleName,
		Certificate: tr.Certificate,
		Date:        tr.Date,
		Judge:       tr.Judge,
	}, nil
}
