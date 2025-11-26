package cat

// Repository defines the interface for cat data operations
type Repository interface {
	// Breed operations
	AllBreeds() ([]*Breed, error)
	GetBreedByID(id int) (*Breed, error)

	// Cat operations
	AllCats() ([]*Cat, error)
	GetCatByID(id int) (*Cat, error)
	InsertCat(cat *Cat) (int, error)
	UpdateCat(cat *Cat) error
	DeleteCat(id int) error
}
