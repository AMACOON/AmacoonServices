package converter

import (
	"amacoonservices/models/services"
	
)

// CatTransferOwnershipDBToCatTransOwner converte um objeto CatTransferOwnershipDB em um objeto CatTransOwner
func CatTransferOwnershipDBToCatTransferOwnerShip(db models.TransferDB) models.Transfer {
	return models.Transfer{
		ID: db.ID,
		Cat: models.CatTransfer{
			ID:           db.ID,
			Name:         db.CatName,
			Breeding:     db.CatBreeding,
			Registration: db.CatRegistration,
			Pedigree:     db.CatPedigree,
			Microchip:    db.CatMicrochip,
			ColorID:      db.CatColorID,
			ColorName:    db.CatColorName,
			EmsCode:      db.CatEmsCode,
			Sex:          db.CatSex,
			Birthdate:    db.CatBirthdate,
			CountryCode:  db.CatCountryCode,
			FatherName:   db.CatFatherName,
			MotherName:   db.CatMotherName,
		},
		Seller: models.SellerTransfer{
			Name:          db.SellerName,
			Email:         db.SellerEmail,
			Phone:         db.SellerPhone,
			MobilePhone:   db.SellerMobilePhone,
			Cep:           db.SellerZipCode,
			Address:       db.SellerAddress,
			Neighborhood:  db.SellerDistrict,
			City:          db.SellerCity,
			Number:        db.SellerNumber,
			Complement:    db.SellerComplement,
			Country:       db.SellerCountry,
		},
		Buyer: models.BuyerTransfer{
			Name:          db.BuyerName,
			CpfOrDoc:      db.BuyerDocument,
			Email:         db.BuyerEmail,
			Phone:         db.BuyerPhone,
			MobilePhone:   db.BuyerMobilePhone,
			Cep:           db.BuyerZipCode,
			Address:       db.BuyerAddress,
			Neighborhood:  db.BuyerDistrict,
			City:          db.BuyerCity,
			Number:        db.BuyerNumber,
			Complement:    db.BuyerComplement,
			Country:       db.BuyerCountry,
		},
		Status:         db.Status,
		ProtocolNumber: db.ProtocolNumber,
	}
}

// CatTransOwnerToCatTransferOwnershipDB converte um objeto CatTransOwner em um objeto CatTransferOwnershipDB
func CatTransferOwnershipToCatTransferOwnerShipDB(m models.Transfer) models.TransferDB {
	return models.TransferDB{
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
}

