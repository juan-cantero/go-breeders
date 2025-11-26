package breeder

import (
	"context"
	"database/sql"
	"time"
)

// MySQLRepository is the MySQL implementation of Repository
type MySQLRepository struct {
	DB *sql.DB
}

// NewMySQLRepository creates a new MySQL repository for breeders
func NewMySQLRepository(db *sql.DB) Repository {
	return &MySQLRepository{DB: db}
}

// AllBreeders returns all breeders from MySQL
func (r *MySQLRepository) AllBreeders() ([]*Breeder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breeder_name, address, city, prov_state,
			country, zip, phone, email, active
			FROM breeders ORDER BY breeder_name`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var breeders []*Breeder
	for rows.Next() {
		var b Breeder
		err := rows.Scan(
			&b.ID, &b.BreederName, &b.Address, &b.City,
			&b.ProvState, &b.Country, &b.Zip, &b.Phone,
			&b.Email, &b.Active,
		)
		if err != nil {
			return nil, err
		}
		breeders = append(breeders, &b)
	}

	return breeders, rows.Err()
}

// GetBreederByID returns a single breeder by ID
func (r *MySQLRepository) GetBreederByID(id int) (*Breeder, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, breeder_name, address, city, prov_state,
			country, zip, phone, email, active
			FROM breeders WHERE id = ?`

	var breeder Breeder
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&breeder.ID, &breeder.BreederName, &breeder.Address, &breeder.City,
		&breeder.ProvState, &breeder.Country, &breeder.Zip, &breeder.Phone,
		&breeder.Email, &breeder.Active,
	)
	if err != nil {
		return nil, err
	}

	return &breeder, nil
}

// InsertBreeder inserts a new breeder and returns the ID
func (r *MySQLRepository) InsertBreeder(breeder *Breeder) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO breeders (breeder_name, address, city, prov_state,
			country, zip, phone, email, active)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.DB.ExecContext(ctx, query,
		breeder.BreederName, breeder.Address, breeder.City, breeder.ProvState,
		breeder.Country, breeder.Zip, breeder.Phone, breeder.Email, breeder.Active,
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

// UpdateBreeder updates an existing breeder
func (r *MySQLRepository) UpdateBreeder(breeder *Breeder) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE breeders SET breeder_name = ?, address = ?, city = ?,
			prov_state = ?, country = ?, zip = ?, phone = ?, email = ?,
			active = ? WHERE id = ?`

	_, err := r.DB.ExecContext(ctx, query,
		breeder.BreederName, breeder.Address, breeder.City, breeder.ProvState,
		breeder.Country, breeder.Zip, breeder.Phone, breeder.Email,
		breeder.Active, breeder.ID,
	)

	return err
}

// DeleteBreeder deletes a breeder by ID
func (r *MySQLRepository) DeleteBreeder(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM breeders WHERE id = ?`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
