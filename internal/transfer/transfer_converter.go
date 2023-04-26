package transfer

import (
	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertTransferRequestToTransfer(req TransferRequest) (Transfer, error) {


	catDataID, err := primitive.ObjectIDFromHex(req.CatData.ID)
	if err != nil {
		return Transfer{}, err
	}

	sellerDataID, err := primitive.ObjectIDFromHex(req.SellerData.ID)
	if err != nil {
		return Transfer{}, err
	}

	buyerDataID, err := primitive.ObjectIDFromHex(req.BuyerData.ID)
	if err != nil {
		return Transfer{}, err
	}

	requesterID, err := primitive.ObjectIDFromHex(req.CatData.ID)
    if err != nil {
        return Transfer{}, err
    }

	catData := CatTransfer{
		ID:           catDataID,
		Name:         req.CatData.Name,
		Registration: req.CatData.Registration,
		Microchip:    req.CatData.Microchip,
		BreedName:    req.CatData.BreedName,
		EmsCode:      req.CatData.EmsCode,
		ColorName:    req.CatData.ColorName,
		Gender:       req.CatData.Gender,
		FatherName:   req.CatData.FatherName,
		MotherName:   req.CatData.MotherName,
	}

	sellerData := OwnerTransfer{
		ID:          sellerDataID,
		Name:        req.SellerData.Name,
		CPF:         req.SellerData.CPF,
		Address:     req.SellerData.Address,
		City:        req.SellerData.City,
		State:       req.SellerData.State,
		ZipCode:     req.SellerData.ZipCode,
		CountryName: req.SellerData.CountryName,
		Phone:       req.SellerData.Phone,
	}

	buyerData := OwnerTransfer{
		ID:          buyerDataID,
		Name:        req.BuyerData.Name,
		CPF:         req.BuyerData.CPF,
		Address:     req.BuyerData.Address,
		City:        req.BuyerData.City,
		State:       req.BuyerData.State,
		ZipCode:     req.BuyerData.ZipCode,
		CountryName: req.BuyerData.CountryName,
		Phone:       req.BuyerData.Phone,
	}

	transfer := Transfer{
		CatData:        catData,
		SellerData:     sellerData,
		BuyerData:      buyerData,
		Status:         req.Status,
		ProtocolNumber: req.ProtocolNumber,
		RequesterID:    requesterID,
		Files:          utils.ConvertFilesReqToFiles(req.Files),
	}

	return transfer, nil
}
