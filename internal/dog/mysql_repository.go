package dog

import (
	"context"
	"database/sql"
	"time"
)

// MySQLRepository is the MySQL implementation of Repository
type MySQLRepository struct {
	DB *sql.DB
}

// NewMySQLRepository creates a new MySQL repository for dogs
func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{DB: db}
}

// AllBreeds returns all dog breeds from MySQL
func (r *MySQLRepository) AllBreeds() ([]*Breed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs,
			CAST(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) AS average_weight,
			lifespan, COALESCE(details, ''),
			COALESCE(alternate_names, ''), COALESCE(geographic_origin, '')
			FROM dog_breeds ORDER BY breed`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var breeds []*Breed
	for rows.Next() {
		var b Breed
		err := rows.Scan(
			&b.ID, &b.Breed, &b.WeightLowLbs, &b.WeightHighLbs,
			&b.AverageWeight, &b.Lifespan, &b.Details,
			&b.AlternateNames, &b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}

	return breeds, rows.Err()
}

// GetBreedByID returns a single dog breed by ID
func (r *MySQLRepository) GetBreedByID(id int) (*Breed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs,
			CAST(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) AS average_weight,
			lifespan, COALESCE(details, ''),
			COALESCE(alternate_names, ''), COALESCE(geographic_origin, '')
			FROM dog_breeds WHERE id = ?`

	var breed Breed
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&breed.ID, &breed.Breed, &breed.WeightLowLbs, &breed.WeightHighLbs,
		&breed.AverageWeight, &breed.Lifespan, &breed.Details,
		&breed.AlternateNames, &breed.GeographicOrigin,
	)
	if err != nil {
		return nil, err
	}

	return &breed, nil
}

// AllDogs returns all dogs from MySQL
func (r *MySQLRepository) AllDogs() ([]*Dog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, dog_name, breed_id, breeder_id, color,
			date_of_birth, spayed_neutered, description, weight
			FROM dogs ORDER BY dog_name`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []*Dog
	for rows.Next() {
		var d Dog
		err := rows.Scan(
			&d.ID, &d.DogName, &d.BreedID, &d.BreederID,
			&d.Color, &d.DateOfBirth, &d.SpayedOrNeutered,
			&d.Description, &d.Weight,
		)
		if err != nil {
			return nil, err
		}
		dogs = append(dogs, &d)
	}

	return dogs, rows.Err()
}

// GetDogByID returns a single dog by ID
func (r *MySQLRepository) GetDogByID(id int) (*Dog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, dog_name, breed_id, breeder_id, color,
			date_of_birth, spayed_neutered, description, weight
			FROM dogs WHERE id = ?`

	var dog Dog
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&dog.ID, &dog.DogName, &dog.BreedID, &dog.BreederID,
		&dog.Color, &dog.DateOfBirth, &dog.SpayedOrNeutered,
		&dog.Description, &dog.Weight,
	)
	if err != nil {
		return nil, err
	}

	return &dog, nil
}

// InsertDog inserts a new dog and returns the ID
func (r *MySQLRepository) InsertDog(dog *Dog) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO dogs (dog_name, breed_id, breeder_id, color,
			date_of_birth, spayed_neutered, description, weight)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.DB.ExecContext(ctx, query,
		dog.DogName, dog.BreedID, dog.BreederID, dog.Color,
		dog.DateOfBirth, dog.SpayedOrNeutered, dog.Description, dog.Weight,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// UpdateDog updates an existing dog
func (r *MySQLRepository) UpdateDog(dog *Dog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE dogs SET dog_name = ?, breed_id = ?, breeder_id = ?,
			color = ?, date_of_birth = ?, spayed_neutered = ?,
			description = ?, weight = ? WHERE id = ?`

	_, err := r.DB.ExecContext(ctx, query,
		dog.DogName, dog.BreedID, dog.BreederID, dog.Color,
		dog.DateOfBirth, dog.SpayedOrNeutered, dog.Description,
		dog.Weight, dog.ID,
	)

	return err
}

// DeleteDog deletes a dog by ID
func (r *MySQLRepository) DeleteDog(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM dogs WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
