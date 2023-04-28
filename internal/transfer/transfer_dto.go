package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/service"
)

type TransferRequest struct {
	CatData        service.CatServiceRequest   `json:"catData"`
	SellerData     service.OwnerServiceRequest `json:"sellerData"`
	BuyerData      service.OwnerServiceRequest `json:"buyerData"`
	Status         string                      `json:"status"`
	ProtocolNumber string                      `json:"protocolNumber"`
	RequesterID    string                      `json:"requesterID"`
}
