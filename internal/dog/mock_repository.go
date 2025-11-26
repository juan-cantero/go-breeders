package dog

import "time"

// MockRepository is a mock implementation for testing
type MockRepository struct{}

// NewMockRepository creates a new mock repository for dogs
func NewMockRepository() Repository {
	return &MockRepository{}
}

// AllBreeds returns mock dog breed data
func (m *MockRepository) AllBreeds() ([]*Breed, error) {
	return []*Breed{
		{
			ID:               1,
			Breed:            "Chihuahua",
			WeightLowLbs:     2,
			WeightHighLbs:    6,
			AverageWeight:    4,
			Lifespan:         15,
			Details:          "Small, alert dog with sassy personality",
			AlternateNames:   "",
			GeographicOrigin: "Mexico",
		},
		{
			ID:               2,
			Breed:            "German Shepherd",
			WeightLowLbs:     50,
			WeightHighLbs:    90,
			AverageWeight:    70,
			Lifespan:         12,
			Details:          "Intelligent, loyal working dog",
			AlternateNames:   "Alsatian",
			GeographicOrigin: "Germany",
		},
		{
			ID:               3,
			Breed:            "Labrador Retriever",
			WeightLowLbs:     55,
			WeightHighLbs:    80,
			AverageWeight:    68,
			Lifespan:         12,
			Details:          "Friendly, outgoing, and active",
			AlternateNames:   "Lab",
			GeographicOrigin: "Canada",
		},
	}, nil
}

// GetBreedByID returns a single mock dog breed
func (m *MockRepository) GetBreedByID(id int) (*Breed, error) {
	breeds, _ := m.AllBreeds()
	for _, breed := range breeds {
		if breed.ID == id {
			return breed, nil
		}
	}
	return nil, nil
}

// AllDogs returns mock dog data
func (m *MockRepository) AllDogs() ([]*Dog, error) {
	return []*Dog{
		{
			ID:               1,
			DogName:          "Max",
			BreedID:          2,
			BreederID:        1,
			Color:            "Black and Tan",
			DateOfBirth:      time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
			SpayedOrNeutered: 0,
			Description:      "Friendly German Shepherd",
			Weight:           75,
		},
		{
			ID:               2,
			DogName:          "Bella",
			BreedID:          1,
			BreederID:        1,
			Color:            "Tan",
			DateOfBirth:      time.Date(2021, 3, 20, 0, 0, 0, 0, time.UTC),
			SpayedOrNeutered: 1,
			Description:      "Small but mighty Chihuahua",
			Weight:           5,
		},
	}, nil
}

// GetDogByID returns a single mock dog
func (m *MockRepository) GetDogByID(id int) (*Dog, error) {
	dogs, _ := m.AllDogs()
	for _, dog := range dogs {
		if dog.ID == id {
			return dog, nil
		}
	}
	return nil, nil
}

// InsertDog simulates inserting a dog
func (m *MockRepository) InsertDog(dog *Dog) (int, error) {
	return 999, nil
}

// UpdateDog simulates updating a dog
func (m *MockRepository) UpdateDog(dog *Dog) error {
	return nil
}

// DeleteDog simulates deleting a dog
func (m *MockRepository) DeleteDog(id int) error {
	return nil
}
