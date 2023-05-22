package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/catservice"
	"github.com/scuba13/AmacoonServices/internal/utils"
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
	Files          *[]FilesTransfer `gorm:"foreignKey:TransferID"`
}

func (Transfer) TableName() string {
	return "service_transfers"
}

type FilesTransfer struct {
	gorm.Model
	TransferID uint
	FileData   utils.Files `gorm:"embedded"`
}

func (FilesTransfer) TableName() string {
	return "service_transfers_files"
}
