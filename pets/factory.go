package pets

// Pet is a simple struct for factory pattern examples
type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"life_span"`
}

func NewPet(species string) *Pet {
	pet := Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description ",
		LifeSpan:    0,
	}

	return &pet
}
