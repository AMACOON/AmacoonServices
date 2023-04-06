package transfer

import (
	
	"github.com/scuba13/AmacoonServices/internal/utils"
	"github.com/sirupsen/logrus"
)

type TransferConverter struct {
	Logger *logrus.Logger
}

func NewTransferConverter(logger *logrus.Logger) *TransferConverter {
	return &TransferConverter{
		Logger:          logger,
	}
}


func (c *TransferConverter) TransferDBToTransfer(transferDB *TransferDB, filesDB []*utils.FilesDB) *Transfer {
	transferData := Transfer{
		ID: transferDB.ID,
		Cat: CatTransfer{
			ID:           transferDB.ID,
			Name:         transferDB.CatName,
			Breeding:     transferDB.CatBreeding,
			Registration: transferDB.CatRegistration,
			Pedigree:     transferDB.CatPedigree,
			Microchip:    transferDB.CatMicrochip,
			ColorID:      transferDB.CatColorID,
			ColorName:    transferDB.CatColorName,
			EmsCode:      transferDB.CatEmsCode,
			Sex:          transferDB.CatSex,
			Birthdate:    transferDB.CatBirthdate,
			CountryCode:  transferDB.CatCountryCode,
			FatherName:   transferDB.CatFatherName,
			MotherName:   transferDB.CatMotherName,
		},
		Seller: SellerTransfer{
			Name:          transferDB.SellerName,
			Email:         transferDB.SellerEmail,
			Phone:         transferDB.SellerPhone,
			MobilePhone:   transferDB.SellerMobilePhone,
			Cep:           transferDB.SellerZipCode,
			Address:       transferDB.SellerAddress,
			Neighborhood:  transferDB.SellerDistrict,
			City:          transferDB.SellerCity,
			Number:        transferDB.SellerNumber,
			Complement:    transferDB.SellerComplement,
			Country:       transferDB.SellerCountry,
		},
		Buyer: BuyerTransfer{
			Name:          transferDB.BuyerName,
			CpfOrDoc:      transferDB.BuyerDocument,
			Email:         transferDB.BuyerEmail,
			Phone:         transferDB.BuyerPhone,
			MobilePhone:   transferDB.BuyerMobilePhone,
			Cep:           transferDB.BuyerZipCode,
			Address:       transferDB.BuyerAddress,
			Neighborhood:  transferDB.BuyerDistrict,
			City:          transferDB.BuyerCity,
			Number:        transferDB.BuyerNumber,
			Complement:    transferDB.BuyerComplement,
			Country:       transferDB.BuyerCountry,
		},
		Status:         transferDB.Status,
		ProtocolNumber: transferDB.ProtocolNumber,
	}
	fileSlice := make([]utils.Files, len(filesDB))
	for i, file := range filesDB {
		fileData := utils.Files{
			ID : file.ID,
			Name   : file.Name,   
			Type    : file.Type,  
			Base64   : file.Base64, 
			ProtocolNumber  : file.ProtocolNumber,
			ServiceID : file.ServiceID,
			
		}
		fileSlice[i] = fileData
	}
	transferData.Files = fileSlice
	
	return &transferData
}


func (c *TransferConverter) TransferToTransferDB(m Transfer) (TransferDB, []*utils.FilesDB) {
	transfer := TransferDB{
		CatID:             int(m.ID),
		CatName:           m.Cat.Name,
		CatBreeding:       m.Cat.Breeding,
		CatRegistration:   m.Cat.Registration,
		CatPedigree:       m.Cat.Pedigree,
		CatMicrochip:      m.Cat.Microchip,
		CatColorID:        m.Cat.ColorID,
		CatColorName:      m.Cat.ColorName,
		CatEmsCode:        m.Cat.EmsCode,
		CatSex:            m.Cat.Sex,
		CatBirthdate:      m.Cat.Birthdate,
		CatCountryCode:    m.Cat.CountryCode,
		CatFatherName:     m.Cat.FatherName,
		CatMotherName:     m.Cat.MotherName,
		SellerName:        m.Seller.Name,
		SellerEmail:       m.Seller.Email,
		SellerPhone:       m.Seller.Phone,
		SellerMobilePhone: m.Seller.MobilePhone,
		SellerZipCode:     m.Seller.Cep,
		SellerAddress:     m.Seller.Address,
		SellerDistrict:    m.Seller.Neighborhood,
		SellerCity:        m.Seller.City,
		SellerNumber:      m.Seller.Number,
		SellerComplement:  m.Seller.Complement,
		SellerCountry:     m.Seller.Country,
		BuyerName:         m.Buyer.Name,
		BuyerDocument:     m.Buyer.CpfOrDoc,
		BuyerEmail:        m.Buyer.Email,
		BuyerPhone:        m.Buyer.Phone,
		BuyerMobilePhone:  m.Buyer.MobilePhone,
		BuyerZipCode:      m.Buyer.Cep,
		BuyerAddress:      m.Buyer.Address,
		BuyerDistrict:     m.Buyer.Neighborhood,
		BuyerCity:         m.Buyer.City,
		BuyerNumber:       m.Buyer.Number,
		BuyerComplement:   m.Buyer.Complement,
		BuyerCountry:      m.Buyer.Country,
		Status:            m.Status,
		ProtocolNumber:    m.ProtocolNumber,
	}
	// Transform each file into a utils.FilesDB struct
	filePointers := make([]*utils.FilesDB, 0)
	for _, fileData := range m.Files {
		files := utils.FilesDB{
			ID : fileData.ID,
			Name   : fileData.Name,   
			Type    : fileData.Type,  
			Base64   : fileData.Base64, 
			ProtocolNumber  : fileData.ProtocolNumber,
			ServiceID : fileData.ServiceID,
		}
		filePointers = append(filePointers, &files)
	}

	return transfer, filePointers
}

