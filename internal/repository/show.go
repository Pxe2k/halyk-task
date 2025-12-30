package repository

import (
	"context"
	"database/sql"
)

type showRepo struct {
	db *sql.DB
}

func NewShowRepo(db *sql.DB) *showRepo {
	return &showRepo{db: db}
}

func (r *showRepo) ShowExistsTx(
	ctx context.Context,
	tx *sql.Tx,
	showID int,
) (bool, error) {

	query := `
	SELECT EXISTS (
		SELECT 1
		FROM shows
		WHERE id = ?
	)
	`

	var exists bool
	if err := tx.QueryRowContext(ctx, query, showID).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
