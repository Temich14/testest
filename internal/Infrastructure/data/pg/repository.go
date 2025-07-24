package pg

import (
	"awesomeProject2/internal/core/entities"
	"context"
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) Add(ctx context.Context, profile entities.Profile) error {
	query := `INSERT INTO profiles (id, first_name, last_name, email, phone) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING;`

	_, err := r.db.ExecContext(ctx, query, profile.ID, profile.FirstName, profile.LastName, profile.Email, profile.Phone)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repo) GetByID(ctx context.Context, id string) (*entities.Profile, error) {
	query := `SELECT * FROM profiles WHERE id = $1 LIMIT 1;`

	row := r.db.QueryRowContext(ctx, query, id)

	var profile entities.Profile
	err := row.Scan(
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Email,
		&profile.Phone,
	)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
