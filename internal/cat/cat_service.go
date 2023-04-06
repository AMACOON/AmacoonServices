package cat



type CatService struct {
	CatRepo CatRepository
}

func NewCatService(catRepo CatRepository) *CatService {
	return &CatService{
		CatRepo: catRepo,
	}
}

func (s *CatService) GetCatsByExhibitorAndSexTable(idExhibitor int, sex int) ([]CatTable, error) {
	return s.CatRepo.GetCatsByExhibitorAndSexTable(idExhibitor, sex)
}

func (s *CatService) GetCatByRegistrationTable(registration string) (*CatTable, error) {
	return s.CatRepo.GetCatByRegistrationTable(registration)
}

func (s *CatService) GetCatsByExhibitorAndSex(idExhibitor int, sex int) ([]Cat, error) {
	return s.CatRepo.GetCatsByExhibitorAndSex(idExhibitor, sex)
}

func (s *CatService) GetCatByRegistration(registration string) (*Cat, error) {
	return s.CatRepo.GetCatByRegistration(registration)
}
