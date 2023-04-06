package breed



type BreedService struct {
	BreedRepo BreedRepository
}

func NewBreedService(breedRepo BreedRepository) *BreedService {
	return &BreedService{
		BreedRepo: breedRepo,
	}
}


func (s *BreedService) GetAllBreeds() ([]Breed, error) {
	return s.BreedRepo.GetAllBreeds()
}

func (s *BreedService) GetCompatibleBreeds(breedID string) ([]string, error) {
	return s.BreedRepo.GetCompatibleBreeds(breedID)
}
