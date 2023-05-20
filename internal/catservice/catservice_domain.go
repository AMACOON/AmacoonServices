package catservice



type CatServiceData struct {
	CatData CatService `json:"catData"`
	OwnerData OwnerService `json:"ownerData"`
}

type CatService struct {
	ID           uint `json:"id"`
	Name         string `json:"name"`
	Registration *string `json:"registration"`
	Microchip    *string `json:"microchip"`
	BreedName    string `json:"breedName"`
	EmsCode      string `json:"emsCode"`
	ColorName    string `json:"colorName"`
	Gender       string `json:"gender"`
	FatherName   *string `json:"fatherName"`
	MotherName   *string `json:"motherName"`
	OwnerID      uint `json:"-"`
	
}

type OwnerService struct {
	ID          uint `json:"id"`
	Name        string `json:"name"`
	CPF         *string `json:"cpf"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	ZipCode     *string `json:"zipCode"`
	CountryName *string `json:"countryName"`
	Phone       *string `json:"phone"`
}
