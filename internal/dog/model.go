package dog

import "time"

// Breed represents a dog breed
type Breed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

// Dog represents an individual dog
type Dog struct {
	ID               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            Breed     `json:"breed,omitempty"`
}
