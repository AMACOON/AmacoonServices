package transfer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (tr *TransferRequest) ToTransferMongo() (*TransferMongo, error) {
	catService, err := tr.CatData.ToCatService()
	if err != nil {
		return nil, err
	}

	sellerService, err := tr.SellerData.ToOwnerService()
	if err != nil {
		return nil, err
	}

	buyerService, err := tr.BuyerData.ToOwnerService()
	if err != nil {
		return nil, err
	}

	requesterID, err := primitive.ObjectIDFromHex(tr.RequesterID)
	if err != nil {
		return nil, err
	}

	return &TransferMongo{
		CatData:        catService,
		SellerData:     sellerService,
		BuyerData:      buyerService,
		Status:         tr.Status,
		ProtocolNumber: tr.ProtocolNumber,
		RequesterID:    requesterID,
	}, nil
}
