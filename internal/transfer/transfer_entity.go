package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/catservice"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	CatData        catservice.CatService   `gorm:"embedded;embeddedPrefix:cat_"`
	SellerData     catservice.OwnerService `gorm:"embedded;embeddedPrefix:seller_"`
	BuyerData      catservice.OwnerService `gorm:"embedded;embeddedPrefix:buyer_"`
	Status         string                  
	ProtocolNumber string                 
	RequesterID    string                 
}

func (Transfer) TableName() string {
	return "service_transfers"
}
