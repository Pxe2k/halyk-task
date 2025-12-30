package repository

import (
	"context"
	"database/sql"

	"github.com/Pxe2k/halyk-task/internal/domain"
)

type seatRepo struct {
	db *sql.DB
}

func NewSeatRepo(db *sql.DB) *seatRepo {
	return &seatRepo{db: db}
}

func (r *seatRepo) CheckAndReserveTx(
	ctx context.Context,
	tx *sql.Tx,
	showID int,
	seatNumbers []int,
) error {
	if len(seatNumbers) == 0 {
		return nil
	}

	query := `
	UPDATE seats
	SET is_reserved = true,
	    reserved_at = NOW()
	WHERE show_id = ?
	  AND seat_number = ?
	  AND is_reserved = false`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var reservedCount int64
	for _, seatNum := range seatNumbers {
		res, err := stmt.ExecContext(ctx, showID, seatNum)
		if err != nil {
			return err
		}

		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		reservedCount += affected
	}

	if reservedCount != int64(len(seatNumbers)) {
		tx.Rollback()
		return domain.ErrSeatUnavailable
	}

	return nil
}
