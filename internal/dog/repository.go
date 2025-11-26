package dog

// Repository defines the interface for dog data operations
// All implementations (MySQL, MongoDB, Mock) must implement this
type Repository interface {
	// Breed operations
	AllBreeds() ([]*Breed, error)
	GetBreedByID(id int) (*Breed, error)

	// Dog operations
	AllDogs() ([]*Dog, error)
	GetDogByID(id int) (*Dog, error)
	InsertDog(dog *Dog) (int, error)
	UpdateDog(dog *Dog) error
	DeleteDog(id int) error
}
