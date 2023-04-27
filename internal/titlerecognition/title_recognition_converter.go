package titlerecognition


import (

	"github.com/scuba13/AmacoonServices/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func ConvertTitleRecognitionRequestToTitleRecognitionMongo(request TitleRecognitionRequest) (TitleRecognitionMongo, error) {
	catID, err := primitive.ObjectIDFromHex(request.CatData.ID)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}

	ownerID, err := primitive.ObjectIDFromHex(request.OwnerData.ID)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}

	requesterID, err := primitive.ObjectIDFromHex(request.RequesterID)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}

	titleID, err := primitive.ObjectIDFromHex(request.TitleID)
	if err != nil {
		return TitleRecognitionMongo{}, err
	}

	catTitle := CatTitle{
		ID:           catID,
		Name:         request.CatData.Name,
		Registration: request.CatData.Registration,
		Microchip:    request.CatData.Microchip,
		BreedName:    request.CatData.BreedName,
		EmsCode:      request.CatData.EmsCode,
		ColorName:    request.CatData.ColorName,
		Gender:       request.CatData.Gender,
		FatherName:   request.CatData.FatherName,
		MotherName:   request.CatData.MotherName,
	}

	ownerTitle := OwnerTitle{
		ID:          ownerID,
		Name:        request.OwnerData.Name,
		CPF:         request.OwnerData.CPF,
		Address:     request.OwnerData.Address,
		City:        request.OwnerData.City,
		State:       request.OwnerData.State,
		ZipCode:     request.OwnerData.ZipCode,
		CountryName: request.OwnerData.CountryName,
		Phone:       request.OwnerData.Phone,
	}



	titleRecognition := TitleRecognitionMongo{
		ID:             request.ID,
		CatData:        catTitle,
		OwnerData:      ownerTitle,
		TitleID:        titleID,
		TitleCode:      request.TitleCode,
		TitleName:      request.TitleName,
		Certificate:    request.Certificate,
		Date:           request.Date,
		Judge:          request.Judge,
		Status:         request.Status,
		ProtocolNumber: request.ProtocolNumber,
		RequesterID:    requesterID,
		Files:          utils.ConvertFilesReqToFiles(request.Files),
	}

	return titleRecognition, nil
}
