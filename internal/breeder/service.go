package breeder

// Service provides business logic for breeder operations
type Service struct {
	repo Repository
}

// NewService creates a new breeder service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetAllBreeders returns all breeders
func (s *Service) GetAllBreeders() ([]*Breeder, error) {
	return s.repo.AllBreeders()
}

// GetBreederByID returns a specific breeder
func (s *Service) GetBreederByID(id int) (*Breeder, error) {
	return s.repo.GetBreederByID(id)
}

// CreateBreeder creates a new breeder
func (s *Service) CreateBreeder(breeder *Breeder) (int, error) {
	return s.repo.InsertBreeder(breeder)
}

// UpdateBreeder updates an existing breeder
func (s *Service) UpdateBreeder(breeder *Breeder) error {
	return s.repo.UpdateBreeder(breeder)
}

// DeleteBreeder deletes a breeder
func (s *Service) DeleteBreeder(id int) error {
	return s.repo.DeleteBreeder(id)
}
