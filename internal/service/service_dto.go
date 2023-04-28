package service

type CatServiceRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Registration string `json:"registration"`
	Microchip    string `json:"microchip"`
	BreedName    string `json:"breedName"`
	EmsCode      string `json:"emsCode"`
	ColorName    string `json:"colorName"`
	Gender       string `json:"gender"`
	FatherName   string `json:"fatherName"`
	MotherName   string `json:"motherName"`
}

type OwnerServiceRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CPF         string `json:"cpf"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zipCode"`
	CountryName string `json:"countryName"`
	Phone       string `json:"phone"`
}
