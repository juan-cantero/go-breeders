package cat

// Service provides business logic for cat operations
type Service struct {
	repo Repository
}

// NewService creates a new cat service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetAllBreeds returns all cat breeds
func (s *Service) GetAllBreeds() ([]*Breed, error) {
	return s.repo.AllBreeds()
}

// GetBreedByID returns a specific cat breed
func (s *Service) GetBreedByID(id int) (*Breed, error) {
	return s.repo.GetBreedByID(id)
}

// GetAllCats returns all cats
func (s *Service) GetAllCats() ([]*Cat, error) {
	return s.repo.AllCats()
}

// GetCatByID returns a specific cat
func (s *Service) GetCatByID(id int) (*Cat, error) {
	return s.repo.GetCatByID(id)
}

// CreateCat creates a new cat
func (s *Service) CreateCat(cat *Cat) (int, error) {
	return s.repo.InsertCat(cat)
}

// UpdateCat updates an existing cat
func (s *Service) UpdateCat(cat *Cat) error {
	return s.repo.UpdateCat(cat)
}

// DeleteCat deletes a cat
func (s *Service) DeleteCat(id int) error {
	return s.repo.DeleteCat(id)
}
