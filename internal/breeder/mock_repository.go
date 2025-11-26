package breeder

// MockRepository is a mock implementation for testing
type MockRepository struct{}

// NewMockRepository creates a new mock repository for breeders
func NewMockRepository() Repository {
	return &MockRepository{}
}

// AllBreeders returns mock breeder data
func (m *MockRepository) AllBreeders() ([]*Breeder, error) {
	return []*Breeder{
		{
			ID:          1,
			BreederName: "Happy Paws Breeders",
			Address:     "123 Main Street",
			City:        "Portland",
			ProvState:   "OR",
			Country:     "USA",
			Zip:         "97201",
			Phone:       "555-1234",
			Email:       "info@happypaws.com",
			Active:      1,
		},
		{
			ID:          2,
			BreederName: "Furry Friends Inc",
			Address:     "456 Oak Avenue",
			City:        "Seattle",
			ProvState:   "WA",
			Country:     "USA",
			Zip:         "98101",
			Phone:       "555-5678",
			Email:       "contact@furryfriends.com",
			Active:      1,
		},
	}, nil
}

// GetBreederByID returns a single mock breeder
func (m *MockRepository) GetBreederByID(id int) (*Breeder, error) {
	breeders, _ := m.AllBreeders()
	for _, breeder := range breeders {
		if breeder.ID == id {
			return breeder, nil
		}
	}
	return nil, nil
}

// InsertBreeder simulates inserting a breeder
func (m *MockRepository) InsertBreeder(breeder *Breeder) (int, error) {
	return 999, nil
}

// UpdateBreeder simulates updating a breeder
func (m *MockRepository) UpdateBreeder(breeder *Breeder) error {
	return nil
}

// DeleteBreeder simulates deleting a breeder
func (m *MockRepository) DeleteBreeder(id int) error {
	return nil
}
