package cat

import (
	"context"
	"database/sql"
	"time"
)

// MySQLRepository is the MySQL implementation of Repository
type MySQLRepository struct {
	DB *sql.DB
}

// NewMySQLRepository creates a new MySQL repository for cats
func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{DB: db}
}

// AllBreeds returns all cat breeds from MySQL
func (r *MySQLRepository) AllBreeds() ([]*Breed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs,
			CAST(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) AS average_weight,
			lifespan, COALESCE(details, ''),
			COALESCE(alternate_names, ''), COALESCE(geographic_origin, '')
			FROM cat_breeds ORDER BY breed`

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

// GetBreedByID returns a single cat breed by ID
func (r *MySQLRepository) GetBreedByID(id int) (*Breed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breed, weight_low_lbs, weight_high_lbs,
			CAST(((weight_low_lbs + weight_high_lbs) / 2) AS unsigned) AS average_weight,
			lifespan, COALESCE(details, ''),
			COALESCE(alternate_names, ''), COALESCE(geographic_origin, '')
			FROM cat_breeds WHERE id = ?`

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

// AllCats returns all cats from MySQL
func (r *MySQLRepository) AllCats() ([]*Cat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, cat_name, breed_id, breeder_id, color,
			date_of_birth, spayed_neutered, description, weight
			FROM cats ORDER BY cat_name`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []*Cat
	for rows.Next() {
		var c Cat
		err := rows.Scan(
			&c.ID, &c.CatName, &c.BreedID, &c.BreederID,
			&c.Color, &c.DateOfBirth, &c.SpayedOrNeutered,
			&c.Description, &c.Weight,
		)
		if err != nil {
			return nil, err
		}
		cats = append(cats, &c)
	}

	return cats, rows.Err()
}

// GetCatByID returns a single cat by ID
func (r *MySQLRepository) GetCatByID(id int) (*Cat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, cat_name, breed_id, breeder_id, color,
			date_of_birth, spayed_neutered, description, weight
			FROM cats WHERE id = ?`

	var cat Cat
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&cat.ID, &cat.CatName, &cat.BreedID, &cat.BreederID,
		&cat.Color, &cat.DateOfBirth, &cat.SpayedOrNeutered,
		&cat.Description, &cat.Weight,
	)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

// InsertCat inserts a new cat and returns the ID
func (r *MySQLRepository) InsertCat(cat *Cat) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO cats (cat_name, breed_id, breeder_id, color,
			date_of_birth, spayed_neutered, description, weight)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.DB.ExecContext(ctx, query,
		cat.CatName, cat.BreedID, cat.BreederID, cat.Color,
		cat.DateOfBirth, cat.SpayedOrNeutered, cat.Description, cat.Weight,
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

// UpdateCat updates an existing cat
func (r *MySQLRepository) UpdateCat(cat *Cat) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE cats SET cat_name = ?, breed_id = ?, breeder_id = ?,
			color = ?, date_of_birth = ?, spayed_neutered = ?,
			description = ?, weight = ? WHERE id = ?`

	_, err := r.DB.ExecContext(ctx, query,
		cat.CatName, cat.BreedID, cat.BreederID, cat.Color,
		cat.DateOfBirth, cat.SpayedOrNeutered, cat.Description,
		cat.Weight, cat.ID,
	)

	return err
}

// DeleteCat deletes a cat by ID
func (r *MySQLRepository) DeleteCat(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM cats WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
