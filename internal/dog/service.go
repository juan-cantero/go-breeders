package dog

// Service provides business logic for dog operations
// This is where you put validation, transformations, complex logic
type Service struct {
	repo Repository
}

// NewService creates a new dog service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetAllBreeds returns all dog breeds
// Business logic can be added here (filtering, sorting, etc.)
func (s *Service) GetAllBreeds() ([]*Breed, error) {
	return s.repo.AllBreeds()
}

// GetBreedByID returns a specific dog breed
func (s *Service) GetBreedByID(id int) (*Breed, error) {
	return s.repo.GetBreedByID(id)
}

// GetAllDogs returns all dogs
func (s *Service) GetAllDogs() ([]*Dog, error) {
	return s.repo.AllDogs()
}

// GetDogByID returns a specific dog
func (s *Service) GetDogByID(id int) (*Dog, error) {
	return s.repo.GetDogByID(id)
}

// CreateDog creates a new dog
// Here you can add validation before inserting
func (s *Service) CreateDog(dog *Dog) (int, error) {
	// Example: Add validation here
	// if dog.DogName == "" {
	//     return 0, errors.New("dog name is required")
	// }

	return s.repo.InsertDog(dog)
}

// UpdateDog updates an existing dog
func (s *Service) UpdateDog(dog *Dog) error {
	return s.repo.UpdateDog(dog)
}

// DeleteDog deletes a dog
func (s *Service) DeleteDog(id int) error {
	return s.repo.DeleteDog(id)
}
