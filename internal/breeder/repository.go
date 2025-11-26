package breeder

// Repository defines the interface for breeder data operations
type Repository interface {
	AllBreeders() ([]*Breeder, error)
	GetBreederByID(id int) (*Breeder, error)
	InsertBreeder(breeder *Breeder) (int, error)
	UpdateBreeder(breeder *Breeder) error
	DeleteBreeder(id int) error
}
