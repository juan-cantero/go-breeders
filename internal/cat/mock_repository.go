package cat

import "time"

// MockRepository is a mock implementation for testing
type MockRepository struct{}

// NewMockRepository creates a new mock repository for cats
func NewMockRepository() Repository {
	return &MockRepository{}
}

// AllBreeds returns mock cat breed data
func (m *MockRepository) AllBreeds() ([]*Breed, error) {
	return []*Breed{
		{
			ID:               1,
			Breed:            "Persian",
			WeightLowLbs:     7,
			WeightHighLbs:    12,
			AverageWeight:    10,
			Lifespan:         15,
			Details:          "Long-haired, gentle cat",
			AlternateNames:   "",
			GeographicOrigin: "Iran",
		},
		{
			ID:               2,
			Breed:            "Siamese",
			WeightLowLbs:     8,
			WeightHighLbs:    12,
			AverageWeight:    10,
			Lifespan:         15,
			Details:          "Vocal, social cat",
			AlternateNames:   "",
			GeographicOrigin: "Thailand",
		},
	}, nil
}

// GetBreedByID returns a single mock cat breed
func (m *MockRepository) GetBreedByID(id int) (*Breed, error) {
	breeds, _ := m.AllBreeds()
	for _, breed := range breeds {
		if breed.ID == id {
			return breed, nil
		}
	}
	return nil, nil
}

// AllCats returns mock cat data
func (m *MockRepository) AllCats() ([]*Cat, error) {
	return []*Cat{
		{
			ID:               1,
			CatName:          "Whiskers",
			BreedID:          1,
			BreederID:        1,
			Color:            "Orange Tabby",
			DateOfBirth:      time.Date(2021, 5, 10, 0, 0, 0, 0, time.UTC),
			SpayedOrNeutered: 1,
			Description:      "Playful tabby cat",
			Weight:           12,
		},
		{
			ID:               2,
			CatName:          "Luna",
			BreedID:          2,
			BreederID:        1,
			Color:            "Seal Point",
			DateOfBirth:      time.Date(2020, 8, 15, 0, 0, 0, 0, time.UTC),
			SpayedOrNeutered: 1,
			Description:      "Talkative Siamese",
			Weight:           10,
		},
	}, nil
}

// GetCatByID returns a single mock cat
func (m *MockRepository) GetCatByID(id int) (*Cat, error) {
	cats, _ := m.AllCats()
	for _, cat := range cats {
		if cat.ID == id {
			return cat, nil
		}
	}
	return nil, nil
}

// InsertCat simulates inserting a cat
func (m *MockRepository) InsertCat(cat *Cat) (int, error) {
	return 999, nil
}

// UpdateCat simulates updating a cat
func (m *MockRepository) UpdateCat(cat *Cat) error {
	return nil
}

// DeleteCat simulates deleting a cat
func (m *MockRepository) DeleteCat(id int) error {
	return nil
}
